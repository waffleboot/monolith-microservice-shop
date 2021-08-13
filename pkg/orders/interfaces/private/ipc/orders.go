package ipc

import (
	"github.com/waffleboot/monolith-microservice-shop/pkg/orders/application"
)

func NewOrdersIPC(service application.OrdersService) OrdersIPC {
	return OrdersIPC{service}
}
