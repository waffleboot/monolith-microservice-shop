package http

import (
	"github.com/go-chi/chi"
	"github.com/waffleboot/monolith-microservice-shop/pkg/orders/application"
	"github.com/waffleboot/monolith-microservice-shop/pkg/orders/domain/orders"
)

func AddRoutes(router *chi.Mux,
	service application.OrdersService, repository orders.Repository) {
	resource := ordersResource{service, repository}
	router.Post("/orders", resource.Post)
	router.Get("/orders/{id}/paid", resource.GetPaid)
}
