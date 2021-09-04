package shop

import (
	domain "monolith-microservice-shop/pkg/orders/domain/orders"
	shop_ipc "monolith-microservice-shop/pkg/shop/interfaces/private/ipc"
)

type wrapper struct {
	shop shop_ipc.ProductInterface
}

func WithShop(shop shop_ipc.ProductInterface) wrapper {
	return wrapper{shop}
}

func (s wrapper) ProductByID(id domain.ProductID) (domain.Product, error) {
	product, err := s.shop.ProductByID(string(id))
	if err != nil {
		return domain.Product{}, err
	}

	return convert(product)
}

func convert(p shop_ipc.Product) (domain.Product, error) {
	return domain.NewProduct(
		domain.ProductID(p.ID),
		p.Name,
		p.Price)
}
