package ipc

import (
	"monolith-microservice-shop/pkg/orders/application"
	"monolith-microservice-shop/pkg/orders/domain/orders"
)

type OrdersService struct {
	service application.OrdersService
}

func NewOrdersService(service application.OrdersService) OrdersService {
	return OrdersService{service}
}

func (o OrdersService) MarkOrderAsPaid(orderID string) error {
	return o.service.MarkOrderAsPaid(
		application.MarkOrderAsPaidCommand{OrderID: orders.ID(orderID)})
}
