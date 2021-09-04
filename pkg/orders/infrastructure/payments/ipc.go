package payments

import (
	"monolith-microservice-shop/pkg/common/price"
	domain "monolith-microservice-shop/pkg/orders/domain/orders"
	payments "monolith-microservice-shop/pkg/payments/interfaces/ipc"
)

type wrapper struct {
	orders chan<- payments.OrderToProcess
}

func WithPaymentsOverChannel(ch chan<- payments.OrderToProcess) wrapper {
	return wrapper{ch}
}

func (s wrapper) InitializeOrderPayment(id domain.OrderID, price price.Price) error {
	s.orders <- payments.OrderToProcess{ID: string(id), Price: price}
	return nil
}
