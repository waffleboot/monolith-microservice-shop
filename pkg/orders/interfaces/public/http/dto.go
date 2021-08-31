package http

import "monolith-microservice-shop/pkg/orders/domain/orders"

type PostOrderRequest struct {
	ProductID orders.ProductID `json:"product_id"`
	Address   PostOrderAddress `json:"address"`
}

type PostOrderAddress struct {
	Name     string `json:"name"`
	Street   string `json:"street"`
	City     string `json:"city"`
	PostCode string `json:"post_code"`
	Country  string `json:"country"`
}

type PostOrdersResponse struct {
	OrderID string
}

type OrderPaidView struct {
	ID     string `json:"id"`
	IsPaid bool   `json:"is_paid"`
}
