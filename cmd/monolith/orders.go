package main

import (
	orders_app "monolith-microservice-shop/pkg/orders/application"
	orders_repo "monolith-microservice-shop/pkg/orders/infrastructure/repo"
	orders_ipc "monolith-microservice-shop/pkg/orders/interfaces/private/ipc"

	orders_payments "monolith-microservice-shop/pkg/orders/infrastructure/payments"
	orders_product "monolith-microservice-shop/pkg/orders/infrastructure/shop"
	payments_ipc "monolith-microservice-shop/pkg/payments/interfaces/ipc"
	shop_ipc "monolith-microservice-shop/pkg/shop/interfaces/private/ipc"
)

func buildOrderService(

	products shop_ipc.ProductInterface,
	payments chan payments_ipc.OrderToProcess) (

	orders_app.OrdersService,
	orders_ipc.OrdersIPC,
	*orders_repo.MemoryRepository) {

	repo := orders_repo.NewMemoryRepository()

	service := orders_app.NewOrdersService(
		orders_product.NewIPCService(products),
		orders_payments.NewIPCService(payments),
		repo,
	)
	return service, orders_ipc.NewOrdersIPC(service), repo
}
