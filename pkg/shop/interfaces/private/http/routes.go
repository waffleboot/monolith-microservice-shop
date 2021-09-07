package http

import (
	shop "monolith-microservice-shop/pkg/shop/domain/products"

	"github.com/go-chi/chi"
)

func AddRoutes(router *chi.Mux, repo shop.Repository) {
	router.Get("/products/{id}", products(repo))
}
