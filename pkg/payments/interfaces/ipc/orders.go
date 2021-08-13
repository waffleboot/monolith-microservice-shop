package ipc

import (
	"log"
	"sync"

	"monolith-microservice-shop/pkg/common/price"
	"monolith-microservice-shop/pkg/payments/application"
)

type OrderToProcess struct {
	ID    string
	Price price.Price
}

type PaymentsIPC struct {
	orders  <-chan OrderToProcess
	service application.PaymentsService
	wg      *sync.WaitGroup
	done    chan struct{}
}

func NewPaymentsIPC(orders <-chan OrderToProcess, service application.PaymentsService) PaymentsIPC {
	return PaymentsIPC{
		orders,
		service,
		&sync.WaitGroup{},
		make(chan struct{}, 1),
	}
}

func (o PaymentsIPC) Run() {
	defer func() {
		o.done <- struct{}{}
	}()
	for order := range o.orders {
		go func(orderToPay OrderToProcess) {
			o.wg.Add(1)
			defer o.wg.Done()

			if err := o.service.InitializeOrderPayment(orderToPay.ID, orderToPay.Price); err != nil {
				log.Print("Cannot initialize payment:", err)
			}
		}(order)
	}
}

func (o PaymentsIPC) Close() {
	o.wg.Wait()
	<-o.done
}
