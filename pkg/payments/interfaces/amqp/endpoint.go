package amqp

import (
	"context"
	"encoding/json"
	"log"

	"monolith-microservice-shop/pkg/common/price"
	"monolith-microservice-shop/pkg/payments/application"

	"github.com/streadway/amqp"
)

type Runner struct {
	conn    *amqp.Connection
	queue   amqp.Queue
	channel *amqp.Channel

	service application.PaymentsService
}

func NewRunner(url string, queue string,
	service application.PaymentsService) (Runner, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return Runner{}, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return Runner{}, err
	}

	q, err := ch.QueueDeclare(
		queue, // name
		true,  // durable
		false, // auto delete
		false, // exclusive
		false, // no wait
		nil,   // args
	)
	if err != nil {
		return Runner{}, err
	}

	return Runner{conn, q, ch, service}, nil
}

func (o Runner) Run(ctx context.Context) error {
	msgs, err := o.channel.Consume(
		o.queue.Name, // queue
		"",           // consumer
		true,         // auto ack
		false,        // exclusive
		false,        // no local
		false,        // no wait
		nil,          // args
	)
	if err != nil {
		return err
	}

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
		case <-ctx.Done():
			return nil
		}
	}
}

func (o Runner) processMsg(msg amqp.Delivery) error {
	var orderView OrderToProcessView
	if err := json.Unmarshal(msg.Body, &orderView); err != nil {
		log.Printf("cannot decode msg: %s, error: %s", string(msg.Body), err)
	}

	orderPrice, err := price.NewPrice(orderView.Price.Cents, orderView.Price.Currency)
	if err != nil {
		log.Printf("cannot decode price for msg %s: %s", string(msg.Body), err)

	}

	return o.service.InitializeOrderPayment(orderView.ID, orderPrice)
}
