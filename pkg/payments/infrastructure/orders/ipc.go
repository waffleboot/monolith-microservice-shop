package orders

import "monolith-microservice-shop/pkg/orders/interfaces/private/ipc"

type IPCService struct {
	orders ipc.OrdersService
}

func NewIPCService(orders ipc.OrdersService) IPCService {
	return IPCService{orders}
}

func (o IPCService) MarkOrderAsPaid(orderID string) error {
	return o.orders.MarkOrderAsPaid(orderID)
}
