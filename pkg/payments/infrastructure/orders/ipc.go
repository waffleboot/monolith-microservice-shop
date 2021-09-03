package orders

import "monolith-microservice-shop/pkg/orders/interfaces/private/ipc"

type IPCService struct {
	orders ipc.IPCService
}

func Wrap(orders ipc.IPCService) IPCService {
	return IPCService{orders}
}

func (o IPCService) MarkOrderAsPaid(orderID string) error {
	return o.orders.MarkOrderAsPaid(orderID)
}
