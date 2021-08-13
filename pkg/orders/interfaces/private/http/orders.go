package http

import (
	"monolith-microservice-shop/pkg/orders/application"

	"github.com/go-chi/chi"
)

func AddRoutes(router *chi.Mux, service application.OrdersService) {
	resource := ordersResource{service}
	router.Post("/orders/{id}/paid", resource.PostPaid)
}
