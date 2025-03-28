package shop

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"monolith-microservice-shop/pkg/common/price"
	"monolith-microservice-shop/pkg/orders/domain/orders"
	shop "monolith-microservice-shop/pkg/shop/interfaces/private/http"

	"github.com/pkg/errors"
)

type httpClient struct {
	address string
}

func WithHttp(address string) httpClient {
	return httpClient{address}
}

func (h httpClient) ProductByID(id orders.ProductID) (orders.Product, error) {
	resp, err := http.Get(fmt.Sprintf("%s/products/%s", h.address, id))
	if err != nil {
		return orders.Product{}, errors.Wrap(err, "request to shop failed")
	}

	defer func() {
		_ = resp.Body.Close()
	}()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return orders.Product{}, errors.Wrap(err, "cannot read response")
	}

	productView := shop.ProductView{}
	if err := json.Unmarshal(b, &productView); err != nil {
		return orders.Product{}, errors.Wrapf(err, "cannot decode response: %s", b)
	}

	return buildProductHttp(productView)
}

func buildProductHttp(v shop.ProductView) (orders.Product, error) {
	price, err := productPrice(v.Price)
	if err != nil {
		return orders.Product{}, errors.Wrap(err, "cannot decode price")
	}
	return orders.NewProduct(
		orders.ProductID(v.ID),
		v.Name,
		price)
}

func productPrice(v shop.PriceView) (price.Price, error) {
	return price.NewPrice(
		v.Cents,
		v.Currency)
}
