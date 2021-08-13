package ipc

import (
	"github.com/waffleboot/monolith-microservice-shop/pkg/orders/application"
	"github.com/waffleboot/monolith-microservice-shop/pkg/orders/domain/orders"
)

type OrdersIPC struct {
	service application.OrdersService
}

func (o OrdersIPC) MarkOrderAsPaid(orderID string) error {
	return o.service.MarkOrderAsPaid(
		application.MarkOrderAsPaidCommand{OrderID: orders.ID(orderID)})
}
