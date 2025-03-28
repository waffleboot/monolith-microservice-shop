package orders

import "errors"

var ErrNotFound = errors.New("order not found")

type Repository interface {
	Save(*Order) error
	ByID(OrderID) (*Order, error)
}
