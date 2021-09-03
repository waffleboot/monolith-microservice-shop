package shop

import (
	"monolith-microservice-shop/pkg/orders/domain/orders"
	"monolith-microservice-shop/pkg/shop/interfaces/private/ipc"
)

type IPCService struct {
	shop ipc.ProductInterface
}

func NewIPCService(shop ipc.ProductInterface) IPCService {
	return IPCService{shop}
}

func (s IPCService) ProductByID(id orders.ProductID) (orders.Product, error) {
	product, err := s.shop.ProductByID(string(id))
	if err != nil {
		return orders.Product{}, err
	}

	return convert(product)
}

func convert(p ipc.Product) (orders.Product, error) {
	return orders.NewProduct(
		orders.ProductID(p.ID),
		p.Name,
		p.Price)
}
