package http

import (
	"github.com/go-chi/chi"
	"github.com/waffleboot/monolith-microservice-shop/pkg/shop/domain/products"
)

func AddRoutes(router *chi.Mux, productsReadModel productsReadModel) {
	resource := productsResource{productsReadModel}
	router.Get("/products", resource.GetAll)
}

type productsReadModel interface {
	AllProducts() ([]products.Product, error)
}
