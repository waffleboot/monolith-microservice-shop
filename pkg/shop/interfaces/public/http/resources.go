package http

import (
	"net/http"

	httputils "monolith-microservice-shop/pkg/common/http"
	"monolith-microservice-shop/pkg/common/price"

	"github.com/go-chi/render"
)

type productsResource struct {
	readModel productsReadModel
}

func (p productsResource) GetAll(w http.ResponseWriter, r *http.Request) {
	products, err := p.readModel.AllProducts()
	if err != nil {
		_ = render.Render(w, r, httputils.ErrInternal(err))
		return
	}

	view := []productView{}
	for _, product := range products {
		view = append(view, productView{
			string(product.ID()),
			product.Name(),
			product.Description(),
			priceViewFromPrice(product.Price()),
		})
	}

	render.Respond(w, r, view)
}

func priceViewFromPrice(p price.Price) priceView {
	return priceView{p.Cents(), p.Currency()}
}
