package shop

import (
	"github.com/waffleboot/monolith-microservice-shop/pkg/orders/domain/orders"
	shop "github.com/waffleboot/monolith-microservice-shop/pkg/shop/interfaces/private/ipc"
)

type IPCService struct {
	shop shop.ProductInterface
}

func NewIPCService(shop shop.ProductInterface) IPCService {
	return IPCService{shop}
}

func (s IPCService) ProductByID(id orders.ProductID) (orders.Product, error) {
	product, err := s.shop.ProductByID(string(id))
	if err != nil {
		return orders.Product{}, err
	}

	return buildProductIPC(product)
}

func buildProductIPC(p shop.Product) (orders.Product, error) {
	return orders.NewProduct(
		orders.ProductID(p.ID),
		p.Name,
		p.Price)
}
