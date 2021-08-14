package shop

import (
	"monolith-microservice-shop/pkg/shop/application"
)

func LoadShopFixtures(productsService application.ProductsService) error {
	err := productsService.AddProduct(application.AddProductCommand{
		ID:            "1",
		Name:          "Product 1",
		Description:   "Some extra description",
		PriceCents:    422,
		PriceCurrency: "USD",
	})
	if err != nil {
		return err
	}

	return productsService.AddProduct(application.AddProductCommand{
		ID:            "2",
		Name:          "Product 2",
		Description:   "Another extra description",
		PriceCents:    333,
		PriceCurrency: "EUR",
	})
}
