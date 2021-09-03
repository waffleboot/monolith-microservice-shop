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

type Payments struct {
	orders  <-chan OrderToProcess
	service application.PaymentsService
	wg      *sync.WaitGroup
	done    chan struct{}
}

func NewPayments(orders <-chan OrderToProcess, service application.PaymentsService) Payments {
	return Payments{
		orders,
		service,
		&sync.WaitGroup{},
		make(chan struct{}, 1),
	}
}

func (o Payments) Run() {
	defer func() {
		o.done <- struct{}{}
	}()
	for order := range o.orders {
		o.wg.Add(1)
		go func(orderToPay OrderToProcess) {
			defer o.wg.Done()

			if err := o.service.InitializeOrderPayment(orderToPay.ID, orderToPay.Price); err != nil {
				log.Print("Cannot initialize payment:", err)
			}
		}(order)
	}
}

func (o Payments) Close() {
	o.wg.Wait()
	<-o.done
}
