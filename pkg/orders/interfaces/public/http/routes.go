package http

import (
	"monolith-microservice-shop/pkg/orders/application"
	domain "monolith-microservice-shop/pkg/orders/domain/orders"

	"github.com/go-chi/chi"
)

func AddRoutes(router *chi.Mux,
	service application.OrdersService, repo domain.Repository) {
	router.Post("/orders", orders(service, repo))
	router.Get("/orders/{id}/paid", paid(repo))
}
