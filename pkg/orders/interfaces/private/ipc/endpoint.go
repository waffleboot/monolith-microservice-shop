package ipc

import (
	"monolith-microservice-shop/pkg/orders/application"
	"monolith-microservice-shop/pkg/orders/domain/orders"
)

type IPCService struct {
	service application.OrdersService
}

func Endpoint(service application.OrdersService) IPCService {
	return IPCService{service}
}

func (o IPCService) MarkOrderAsPaid(orderID string) error {
	return o.service.MarkOrderAsPaid(
		application.MarkOrderAsPaidCommand{OrderID: orders.OrderID(orderID)})
}
