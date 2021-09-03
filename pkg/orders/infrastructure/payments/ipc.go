package payments

import (
	"monolith-microservice-shop/pkg/common/price"
	domain "monolith-microservice-shop/pkg/orders/domain/orders"
	payments "monolith-microservice-shop/pkg/payments/interfaces/ipc"
)

type IPCService struct {
	orders chan<- payments.OrderToProcess
}

func NewIPCService(ch chan<- payments.OrderToProcess) IPCService {
	return IPCService{ch}
}

func (s IPCService) InitializeOrderPayment(id domain.OrderID, price price.Price) error {
	s.orders <- payments.OrderToProcess{ID: string(id), Price: price}
	return nil
}
