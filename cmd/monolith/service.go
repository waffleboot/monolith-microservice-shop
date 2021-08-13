package main

import (
	"monolith-microservice-shop/pkg/common/cmd"
	orders_public_http "monolith-microservice-shop/pkg/orders/interfaces/public/http"
	payments_app "monolith-microservice-shop/pkg/payments/application"
	payments_infra_orders "monolith-microservice-shop/pkg/payments/infrastructure/orders"
	payments_ipc "monolith-microservice-shop/pkg/payments/interfaces/ipc"
	"monolith-microservice-shop/pkg/shop"
	shop_interfaces_http "monolith-microservice-shop/pkg/shop/interfaces/public/http"

	"github.com/go-chi/chi"
)

func createService(paymentsChannel chan payments_ipc.OrderToProcess) (*chi.Mux, payments_ipc.PaymentsIPC) {

	shopService, shopIpc, shopRepo := buildShopService()
	orders, ordersIpc, ordersRepo := buildOrderService(shopIpc, paymentsChannel)

	paymentsService := payments_app.NewPaymentsService(
		payments_infra_orders.NewIPCService(ordersIpc),
	)
	paymentsIntraprocessInterface := payments_ipc.NewPaymentsIPC(paymentsChannel, paymentsService)

	if err := shop.LoadShopFixtures(shopService); err != nil {
		panic(err)
	}

	r := cmd.CreateRouter()
	shop_interfaces_http.AddRoutes(r, shopRepo)
	orders_public_http.AddRoutes(r, orders, ordersRepo)

	return r, paymentsIntraprocessInterface
}
