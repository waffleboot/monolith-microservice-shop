package repo

import "monolith-microservice-shop/pkg/orders/domain/orders"

var _ orders.Repository = (*MemoryRepository)(nil)

type MemoryRepository struct {
	orders []orders.Order
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{[]orders.Order{}}
}

func (m *MemoryRepository) Save(order *orders.Order) error {
	for i, p := range m.orders {
		if p.ID() == order.ID() {
			m.orders[i] = *order
			return nil
		}
	}
	m.orders = append(m.orders, *order)
	return nil
}

func (m MemoryRepository) ByID(id orders.ID) (*orders.Order, error) {
	for _, p := range m.orders {
		if p.ID() == id {
			return &p, nil
		}
	}
	return nil, orders.ErrNotFound
}
