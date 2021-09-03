package main

import (
	"monolith-microservice-shop/pkg/orders/application"
	"monolith-microservice-shop/pkg/orders/infrastructure/repo"
	"monolith-microservice-shop/pkg/orders/interfaces/private/ipc"

	orders_payments "monolith-microservice-shop/pkg/orders/infrastructure/payments"
	orders_product "monolith-microservice-shop/pkg/orders/infrastructure/shop"
	payments_ipc "monolith-microservice-shop/pkg/payments/interfaces/ipc"
	shop_ipc "monolith-microservice-shop/pkg/shop/interfaces/private/ipc"
)

func buildOrderService(

	products shop_ipc.ProductInterface,
	ch chan payments_ipc.OrderToProcess) (

	application.OrdersService,
	ipc.OrdersService,
	*repo.MemoryRepository) {

	repo := repo.NewMemoryRepository()

	service := application.NewOrdersService(
		orders_product.NewIPCService(products),
		orders_payments.NewIPCService(ch),
		repo,
	)
	return service, ipc.NewOrdersService(service), repo
}
