package shop

import (
	"testing"

	"monolith-microservice-shop/pkg/common/price"
	"monolith-microservice-shop/pkg/orders/domain/orders"
	"monolith-microservice-shop/pkg/shop/interfaces/private/ipc"

	"github.com/stretchr/testify/assert"
)

func TestOrderProductFromShopProduct(t *testing.T) {
	shopProduct := ipc.Product{
		ID:          "123",
		Name:        "name",
		Description: "desc",
		Price:       price.NewPriceP(42, "EUR"),
	}
	orderProduct, err := convert(shopProduct)
	assert.NoError(t, err)

	expectedOrderProduct, err := orders.NewProduct("123", "name", price.NewPriceP(42, "EUR"))
	assert.NoError(t, err)

	assert.EqualValues(t, expectedOrderProduct, orderProduct)
}
