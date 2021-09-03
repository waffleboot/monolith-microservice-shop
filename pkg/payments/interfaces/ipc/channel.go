package ipc

import "monolith-microservice-shop/pkg/common/price"

type OrderToProcess struct {
	ID    string
	Price price.Price
}
