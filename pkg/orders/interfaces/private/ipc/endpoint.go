package ipc

import (
	"monolith-microservice-shop/pkg/orders/application"
	"monolith-microservice-shop/pkg/orders/domain/orders"
)

type Wrapper struct {
	service application.OrdersService
}

func WithOrders(service application.OrdersService) Wrapper {
	return Wrapper{service}
}

func (o Wrapper) MarkOrderAsPaid(orderID string) error {
	return o.service.MarkOrderAsPaid(
		application.MarkOrderAsPaidCommand{OrderID: orders.OrderID(orderID)})
}
