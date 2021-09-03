package main

import (
	"monolith-microservice-shop/pkg/common/cmd"
	orders_public_http "monolith-microservice-shop/pkg/orders/interfaces/public/http"
	payments_app "monolith-microservice-shop/pkg/payments/application"
	payments_orders "monolith-microservice-shop/pkg/payments/infrastructure/orders"
	payments_ipc "monolith-microservice-shop/pkg/payments/interfaces/ipc"
	"monolith-microservice-shop/pkg/shop"
	shop_http "monolith-microservice-shop/pkg/shop/interfaces/public/http"

	"github.com/go-chi/chi"
)

func createService(ch chan payments_ipc.OrderToProcess) (*chi.Mux, payments_ipc.Payments) {

	shopService, shopIpc, shopRepo := buildShopService()
	orders, ordersIpc, ordersRepo := buildOrderService(shopIpc, ch)

	payments := payments_app.NewPaymentsService(
		payments_orders.NewIPCService(ordersIpc),
	)
	paymentsInterface := payments_ipc.NewPayments(ch, payments)

	if err := shop.LoadShopFixtures(shopService); err != nil {
		panic(err)
	}

	r := cmd.CreateRouter()
	shop_http.AddRoutes(r, shopRepo)
	orders_public_http.AddRoutes(r, orders, ordersRepo)

	return r, paymentsInterface
}
