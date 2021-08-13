package shop_test

import (
	"testing"

	"monolith-microservice-shop/pkg/common/price"
	"monolith-microservice-shop/pkg/orders/domain/orders"
	"monolith-microservice-shop/pkg/orders/infrastructure/shop"
	"monolith-microservice-shop/pkg/shop/interfaces/private/ipc"

	"github.com/stretchr/testify/assert"
)

func TestOrderProductFromShopProduct(t *testing.T) {
	shopProduct := ipc.Product{
		"123",
		"name",
		"desc",
		price.NewPriceP(42, "EUR"),
	}
	orderProduct, err := shop.BuildProductIPC(shopProduct)
	assert.NoError(t, err)

	expectedOrderProduct, err := orders.NewProduct("123", "name", price.NewPriceP(42, "EUR"))
	assert.NoError(t, err)

	assert.EqualValues(t, expectedOrderProduct, orderProduct)
}
