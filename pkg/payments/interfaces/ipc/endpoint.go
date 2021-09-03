package ipc

import (
	"log"
	"sync"

	"monolith-microservice-shop/pkg/payments/application"
)

type Runner struct {
	ch      <-chan OrderToProcess
	service application.PaymentsService
	wg      *sync.WaitGroup
	done    chan struct{}
}

func NewRunner(ch <-chan OrderToProcess, service application.PaymentsService) Runner {
	return Runner{
		ch,
		service,
		&sync.WaitGroup{},
		make(chan struct{}, 1),
	}
}

func (o Runner) Run() {
	defer func() {
		o.done <- struct{}{}
	}()
	for order := range o.ch {
		o.wg.Add(1)
		go func(orderToPay OrderToProcess) {
			defer o.wg.Done()

			if err := o.service.InitializeOrderPayment(orderToPay.ID, orderToPay.Price); err != nil {
				log.Print("Cannot initialize payment:", err)
			}
		}(order)
	}
}

func (o Runner) Stop() {
	o.wg.Wait()
	<-o.done
}
