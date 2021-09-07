package http

import (
	"monolith-microservice-shop/pkg/shop/domain/products"
)

type productsReadModel interface {
	AllProducts() ([]products.Product, error)
}
