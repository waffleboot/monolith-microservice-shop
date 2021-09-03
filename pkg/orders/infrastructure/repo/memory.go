package repo

import domain "monolith-microservice-shop/pkg/orders/domain/orders"

type MemoryRepository struct {
	orders []domain.Order
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{[]domain.Order{}}
}

func (m *MemoryRepository) Save(order *domain.Order) error {
	for i, p := range m.orders {
		if p.ID() == order.ID() {
			m.orders[i] = *order
			return nil
		}
	}
	m.orders = append(m.orders, *order)
	return nil
}

func (m MemoryRepository) ByID(id domain.OrderID) (*domain.Order, error) {
	for _, p := range m.orders {
		if p.ID() == id {
			return &p, nil
		}
	}
	return nil, domain.ErrNotFound
}
