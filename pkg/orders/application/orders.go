package application

import (
	"log"

	"monolith-microservice-shop/pkg/common/price"
	"monolith-microservice-shop/pkg/orders/domain/orders"

	"github.com/pkg/errors"
)

type productService interface {
	ProductByID(id orders.ProductID) (orders.Product, error)
}

type paymentService interface {
	InitializeOrderPayment(id orders.ID, price price.Price) error
}

///////////////////////////////////////////////////

type OrdersService struct {
	products productService
	payments paymentService
	repo     orders.Repository
}

func NewOrdersService(products productService, payments paymentService, repo orders.Repository) OrdersService {
	return OrdersService{products, payments, repo}
}

func (s OrdersService) PlaceOrder(cmd PlaceOrderCommand) error {
	address, err := orders.NewAddress(
		cmd.Address.Name,
		cmd.Address.Street,
		cmd.Address.City,
		cmd.Address.PostCode,
		cmd.Address.Country,
	)
	if err != nil {
		return errors.Wrap(err, "invalid address")
	}

	product, err := s.products.ProductByID(cmd.ProductID)
	if err != nil {
		return errors.Wrap(err, "cannot get product")
	}

	newOrder, err := orders.NewOrder(cmd.OrderID, product, address)
	if err != nil {
		return errors.Wrap(err, "cannot create order")
	}

	if err := s.repo.Save(newOrder); err != nil {
		return errors.Wrap(err, "cannot save order")
	}

	if err := s.payments.InitializeOrderPayment(newOrder.ID(), newOrder.Product().Price()); err != nil {
		return errors.Wrap(err, "cannot initialize payment")
	}

	log.Printf("order %s placed", cmd.OrderID)

	return nil
}

func (s OrdersService) MarkOrderAsPaid(cmd MarkOrderAsPaidCommand) error {
	o, err := s.repo.ByID(cmd.OrderID)
	if err != nil {
		return errors.Wrapf(err, "cannot get order %s", cmd.OrderID)
	}

	o.MarkAsPaid()

	if err := s.repo.Save(o); err != nil {
		return errors.Wrap(err, "cannot save order")
	}

	log.Printf("marked order %s as paid", cmd.OrderID)

	return nil
}

// func (s OrdersService) OrderByID(id orders.ID) (orders.Order, error) {
// 	o, err := s.ordersRepository.ByID(id)
// 	if err != nil {
// 		return orders.Order{}, errors.Wrapf(err, "cannot get order %s", id)
// 	}

// 	return *o, nil
// }
