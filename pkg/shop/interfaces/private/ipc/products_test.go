package ipc_test

import (
	"testing"

	"monolith-microservice-shop/pkg/common/price"
	"monolith-microservice-shop/pkg/shop/domain/products"
	"monolith-microservice-shop/pkg/shop/interfaces/private/ipc"

	"github.com/stretchr/testify/assert"
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
