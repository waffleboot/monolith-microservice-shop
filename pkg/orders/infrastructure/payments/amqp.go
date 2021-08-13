package payments

import (
	"encoding/json"
	"log"

	"github.com/pkg/errors"
	"github.com/streadway/amqp"
	"github.com/waffleboot/monolith-microservice-shop/pkg/common/price"
	"github.com/waffleboot/monolith-microservice-shop/pkg/orders/domain/orders"
	payments_amqp "github.com/waffleboot/monolith-microservice-shop/pkg/payments/interfaces/amqp"
)

type AMQPService struct {
	queue   amqp.Queue
	channel *amqp.Channel
}

func NewAMQPService(url, queueName string) (AMQPService, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return AMQPService{}, err
	}
	ch, err := conn.Channel()
	if err != nil {
		return AMQPService{}, err
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
		return AMQPService{}, err
	}
	return AMQPService{q, ch}, nil
}

func (s AMQPService) InitializeOrderPayment(id orders.ID, price price.Price) error {
	order := payments_amqp.OrderToProcessView{
		ID: string(id),
		Price: payments_amqp.PriceView{
			Cents:    price.Cents(),
			Currency: price.Currency(),
		},
	}

	b, err := json.Marshal(order)
	if err != nil {
		return errors.Wrap(err, "cannot marshal order for amqp")
	}

	err = s.channel.Publish(
		"",
		s.queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        b,
		})
	if err != nil {
		return errors.Wrap(err, "cannot send order to amqp")
	}

	log.Printf("sent order %s to amqp", id)

	return nil
}

func (s AMQPService) Close() error {
	return s.channel.Close()
}
