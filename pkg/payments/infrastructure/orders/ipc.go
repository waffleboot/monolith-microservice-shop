package orders

import "monolith-microservice-shop/pkg/orders/interfaces/private/ipc"

type wrapper struct {
	orders ipc.Wrapper
}

func WithOrders(orders ipc.Wrapper) wrapper {
	return wrapper{orders}
}

func (o wrapper) MarkOrderAsPaid(orderID string) error {
	return o.orders.MarkOrderAsPaid(orderID)
}
