package orders

import "github.com/waffleboot/monolith-microservice-shop/pkg/orders/interfaces/private/ipc"

type IPCService struct {
	orders ipc.OrdersIPC
}

func NewIPCService(paymentsInterface ipc.OrdersIPC) IPCService {
	return IPCService{paymentsInterface}
}

func (o IPCService) MarkOrderAsPaid(orderID string) error {
	return o.orders.MarkOrderAsPaid(orderID)
}
