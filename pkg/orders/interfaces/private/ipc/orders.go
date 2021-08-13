package ipc

import (
	"github.com/ThreeDotsLabs/monolith-microservice-shop/pkg/orders/application"
)

func NewOrdersIPC(service application.OrdersService) OrdersIPC {
	return OrdersIPC{service}
}
