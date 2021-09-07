package application

import (
	"log"
	domain "monolith-microservice-shop/pkg/orders/domain/orders"

	"github.com/pkg/errors"
)

type PlaceOrderCommand struct {
	OrderID   domain.OrderID
	ProductID domain.ProductID
	Address   PlaceOrderCommandAddress
}

type PlaceOrderCommandAddress struct {
	Name     string
	Street   string
	City     string
	PostCode string
	Country  string
}

func (s OrdersService) PlaceOrder(cmd PlaceOrderCommand, repo domain.Repository) error {
	address, err := domain.NewAddress(
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

	newOrder, err := domain.NewOrder(cmd.OrderID, product, address)
	if err != nil {
		return errors.Wrap(err, "cannot create order")
	}

	if err := repo.Save(newOrder); err != nil {
		return errors.Wrap(err, "cannot save order")
	}

	if err := s.payments.InitializeOrderPayment(newOrder.ID(), newOrder.Product().Price()); err != nil {
		return errors.Wrap(err, "cannot initialize payment")
	}

	log.Printf("order %s placed", cmd.OrderID)

	return nil
}
