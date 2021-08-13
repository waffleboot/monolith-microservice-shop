package amqp

import (
	"context"
	"encoding/json"
	"log"

	"github.com/streadway/amqp"
	"github.com/waffleboot/monolith-microservice-shop/pkg/common/price"
	"github.com/waffleboot/monolith-microservice-shop/pkg/payments/application"
)

type OrderToProcessView struct {
	ID    string `json:"id"`
	Price PriceView
}

type PriceView struct {
	Cents    uint   `json:"cents"`
	Currency string `json:"currency"`
}

type PaymentsAMQP struct {
	conn    *amqp.Connection
	queue   amqp.Queue
	channel *amqp.Channel

	service application.PaymentsService
}

func NewPaymentsAMQP(url string, queueName string, service application.PaymentsService) (PaymentsAMQP, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return PaymentsAMQP{}, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return PaymentsAMQP{}, err
	}

	q, err := ch.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return PaymentsAMQP{}, err
	}

	return PaymentsAMQP{conn, q, ch, service}, nil
}

func (o PaymentsAMQP) Run(ctx context.Context) error {
	msgs, err := o.channel.Consume(
		o.queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	done := ctx.Done()
	defer func() {
		if err := o.conn.Close(); err != nil {
			log.Print("cannot close conn: ", err)
		}
		if err := o.channel.Close(); err != nil {
			log.Print("cannot close channel: ", err)
		}
	}()

	for {
		select {
		case msg := <-msgs:
			err := o.processMsg(msg)
			if err != nil {
				log.Printf("cannot process msg: %s, err: %s", msg.Body, err)
			}
		case <-done:
			return nil
		}
	}
}

func (o PaymentsAMQP) processMsg(msg amqp.Delivery) error {
	orderView := OrderToProcessView{}
	err := json.Unmarshal(msg.Body, &orderView)
	if err != nil {
		log.Printf("cannot decode msg: %s, error: %s", string(msg.Body), err)
	}

	orderPrice, err := price.NewPrice(orderView.Price.Cents, orderView.Price.Currency)
	if err != nil {
		log.Printf("cannot decode price for msg %s: %s", string(msg.Body), err)

	}

	return o.service.InitializeOrderPayment(orderView.ID, orderPrice)
}
