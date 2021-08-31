package http

import (
	"monolith-microservice-shop/pkg/orders/application"

	"github.com/go-chi/chi"
)

func AddRoutes(router *chi.Mux, service application.OrdersService) {
	router.Post("/orders/{id}/paid", ordersEndpoint{service}.paid)
}
