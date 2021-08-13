package orders_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/waffleboot/monolith-microservice-shop/pkg/common/price"
	order_domain "github.com/waffleboot/monolith-microservice-shop/pkg/orders/domain/orders"
	"github.com/waffleboot/monolith-microservice-shop/pkg/orders/infrastructure/repo"
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

func addOrder(t *testing.T, repo *repo.MemoryRepository, id string) *order_domain.Order {
	productPrice, err := price.NewPrice(10, "USD")
	assert.NoError(t, err)

	orderProduct, err := order_domain.NewProduct("1", "foo", productPrice)
	assert.NoError(t, err)

	orderAddress, err := order_domain.NewAddress("test", "test", "test", "test", "test")
	assert.NoError(t, err)

	p, err := order_domain.NewOrder(order_domain.ID(id), orderProduct, orderAddress)
	assert.NoError(t, err)

	err = repo.Save(p)
	assert.NoError(t, err)

	return p
}
