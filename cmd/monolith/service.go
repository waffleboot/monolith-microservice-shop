package main

import (
	. "monolith-microservice-shop/pkg/orders/application"
	. "monolith-microservice-shop/pkg/orders/infrastructure/payments"
	orders_repo "monolith-microservice-shop/pkg/orders/infrastructure/repo"
	. "monolith-microservice-shop/pkg/orders/infrastructure/shop"
	orders_interface "monolith-microservice-shop/pkg/orders/interfaces/private/ipc"
	orders_http "monolith-microservice-shop/pkg/orders/interfaces/public/http"
	payments "monolith-microservice-shop/pkg/payments/application"
	payments_orders "monolith-microservice-shop/pkg/payments/infrastructure/orders"
	payments_ipc "monolith-microservice-shop/pkg/payments/interfaces/ipc"
	. "monolith-microservice-shop/pkg/shop"
	shop_app "monolith-microservice-shop/pkg/shop/application"
	shop_repo "monolith-microservice-shop/pkg/shop/infrastructure/repo"
	shop_ipc "monolith-microservice-shop/pkg/shop/interfaces/private/ipc"
	shop_http "monolith-microservice-shop/pkg/shop/interfaces/public/http"

	"github.com/go-chi/chi"
)

func build(router *chi.Mux, ch chan payments_ipc.OrderToProcess) payments_ipc.Runner {

	shopRepo := shop_repo.NewMemoryRepository()
	if err := LoadShopFixtures(shop_app.NewService(shopRepo, shopRepo)); err != nil {
		panic(err)
	}

	ordersRepo := orders_repo.NewMemoryRepository()
	ordersService := NewOrdersService(
		WithShop(shop_ipc.NewProductInterface(shopRepo)),
		WithPaymentsOverChannel(ch),
		ordersRepo,
	)

	paymentsService := payments.NewService(
		payments_orders.WithOrders(orders_interface.WithOrders(ordersService)),
	)

	shop_http.AddRoutes(router, shopRepo)
	orders_http.AddRoutes(router, ordersService, ordersRepo)

	return payments_ipc.NewRunner(ch, paymentsService)
}
