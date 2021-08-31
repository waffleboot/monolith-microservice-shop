package http

import (
	"monolith-microservice-shop/pkg/orders/application"
	"monolith-microservice-shop/pkg/orders/domain/orders"

	"github.com/go-chi/chi"
)

func AddRoutes(router *chi.Mux,
	service application.OrdersService, repository orders.Repository) {
	endpoint := ordersEndpoint{service, repository}
	router.Post("/orders", endpoint.orders)
	router.Get("/orders/{id}/paid", endpoint.paid)
}
