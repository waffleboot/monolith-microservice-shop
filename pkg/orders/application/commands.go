package application

import domain "monolith-microservice-shop/pkg/orders/domain/orders"

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

type MarkOrderAsPaidCommand struct {
	OrderID domain.OrderID
}
