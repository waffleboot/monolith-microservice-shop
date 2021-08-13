package ipc_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/waffleboot/monolith-microservice-shop/pkg/common/price"
	"github.com/waffleboot/monolith-microservice-shop/pkg/shop/domain/products"
	"github.com/waffleboot/monolith-microservice-shop/pkg/shop/interfaces/private/ipc"
)

func TestProductFromDomainProduct(t *testing.T) {
	productPrice := price.NewPriceP(42, "USD")
	domainProduct, err := products.NewProduct("123", "name", "desc", productPrice)
	assert.NoError(t, err)

	p := ipc.ProductFromDomainProduct(*domainProduct)

	assert.EqualValues(t, ipc.Product{
		"123",
		"name",
		"desc",
		productPrice,
	}, p)
}
