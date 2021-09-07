package application

import (
	"monolith-microservice-shop/pkg/common/price"
	domain "monolith-microservice-shop/pkg/orders/domain/orders"
)

type productService interface {
	ProductByID(id domain.ProductID) (domain.Product, error)
}

type paymentService interface {
	InitializeOrderPayment(id domain.OrderID, price price.Price) error
}

///////////////////////////////////////////////////

type OrdersService struct {
	products productService
	payments paymentService
	repo     domain.Repository
}

func NewOrdersService(products productService, payments paymentService, repo domain.Repository) OrdersService {
	return OrdersService{products, payments, repo}
}

// func (s OrdersService) OrderByID(id orders.ID) (orders.Order, error) {
// 	o, err := s.ordersRepository.ByID(id)
// 	if err != nil {
// 		return orders.Order{}, errors.Wrapf(err, "cannot get order %s", id)
// 	}

// 	return *o, nil
// }
