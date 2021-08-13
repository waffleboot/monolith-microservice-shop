package orders

import "monolith-microservice-shop/pkg/orders/interfaces/private/ipc"

type IPCService struct {
	orders ipc.OrdersIPC
}

func NewIPCService(orders ipc.OrdersIPC) IPCService {
	return IPCService{orders}
}

func (o IPCService) MarkOrderAsPaid(orderID string) error {
	return o.orders.MarkOrderAsPaid(orderID)
}
