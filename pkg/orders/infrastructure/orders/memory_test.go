package orders_test

import (
	"testing"

	"monolith-microservice-shop/pkg/common/price"
	domain "monolith-microservice-shop/pkg/orders/domain/orders"
	"monolith-microservice-shop/pkg/orders/infrastructure/repo"

	"github.com/stretchr/testify/assert"
)

func TestMemoryRepository(t *testing.T) {
	repo := repo.NewMemoryRepository()

	order1 := addOrder(t, repo, "1")
	// test idempotency
	_ = addOrder(t, repo, "1")

	repoOrder1, err := repo.ByID("1")
	assert.NoError(t, err)
	assert.EqualValues(t, *order1, *repoOrder1)

	order2 := addOrder(t, repo, "2")

	repoOrder2, err := repo.ByID("2")
	assert.NoError(t, err)
	assert.EqualValues(t, *order2, *repoOrder2)
}

func addOrder(t *testing.T, repo *repo.MemoryRepository, id string) *domain.Order {
	productPrice, err := price.NewPrice(10, "USD")
	assert.NoError(t, err)

	orderProduct, err := domain.NewProduct("1", "foo", productPrice)
	assert.NoError(t, err)

	orderAddress, err := domain.NewAddress("test", "test", "test", "test", "test")
	assert.NoError(t, err)

	p, err := domain.NewOrder(domain.OrderID(id), orderProduct, orderAddress)
	assert.NoError(t, err)

	err = repo.Save(p)
	assert.NoError(t, err)

	return p
}
