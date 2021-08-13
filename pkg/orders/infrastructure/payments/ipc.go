package payments

import (
	"github.com/waffleboot/monolith-microservice-shop/pkg/common/price"
	"github.com/waffleboot/monolith-microservice-shop/pkg/orders/domain/orders"
	payments_ipc "github.com/waffleboot/monolith-microservice-shop/pkg/payments/interfaces/ipc"
)

type IPCService struct {
	orders chan<- payments_ipc.OrderToProcess
}

func NewIPCService(ch chan<- payments_ipc.OrderToProcess) IPCService {
	return IPCService{ch}
}

func (s IPCService) InitializeOrderPayment(id orders.ID, price price.Price) error {
	s.orders <- payments_ipc.OrderToProcess{ID: string(id), Price: price}
	return nil
}
