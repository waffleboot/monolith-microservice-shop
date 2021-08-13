package application

import (
	"log"
	"time"

	"github.com/waffleboot/monolith-microservice-shop/pkg/common/price"
)

type ordersService interface {
	MarkOrderAsPaid(orderID string) error
}

type PaymentsService struct {
	orders ordersService
}

func NewPaymentsService(orders ordersService) PaymentsService {
	return PaymentsService{orders}
}

func (s PaymentsService) InitializeOrderPayment(orderID string, price price.Price) error {
	// ...
	log.Printf("initializing payment for order %s", orderID)

	go func() {
		time.Sleep(time.Millisecond * 500)
		if err := s.postOrderPayment(orderID); err != nil {
			log.Printf("cannot post order payment: %s", err)
		}
	}()

	// simulating payments provider delay
	//time.Sleep(time.Second)

	return nil
}

func (s PaymentsService) postOrderPayment(orderID string) error {
	log.Printf("payment for order %s done, marking order as paid", orderID)
	return s.orders.MarkOrderAsPaid(orderID)
}
