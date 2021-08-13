package http

import (
	"monolith-microservice-shop/pkg/shop/domain/products"

	"github.com/go-chi/chi"
)

func AddRoutes(router *chi.Mux, productsReadModel productsReadModel) {
	resource := productsResource{productsReadModel}
	router.Get("/products", resource.GetAll)
}

type productsReadModel interface {
	AllProducts() ([]products.Product, error)
}
