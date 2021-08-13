package http

import (
	"github.com/go-chi/chi"
	"github.com/waffleboot/monolith-microservice-shop/pkg/orders/application"
)

func AddRoutes(router *chi.Mux, service application.OrdersService) {
	resource := ordersResource{service}
	router.Post("/orders/{id}/paid", resource.PostPaid)
}
