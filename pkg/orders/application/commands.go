package application

import "github.com/ThreeDotsLabs/monolith-microservice-shop/pkg/orders/domain/orders"

type PlaceOrderCommand struct {
	OrderID   orders.ID
	ProductID orders.ProductID

	Address PlaceOrderCommandAddress
}

type PlaceOrderCommandAddress struct {
	Name     string
	Street   string
	City     string
	PostCode string
	Country  string
}

type MarkOrderAsPaidCommand struct {
	OrderID orders.ID
}
