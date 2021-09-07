package http

import (
	"net/http"

	httputils "monolith-microservice-shop/pkg/common/http"
	"monolith-microservice-shop/pkg/common/price"
	shop "monolith-microservice-shop/pkg/shop/domain/products"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type ProductView struct {
	ID string `json:"id"`

	Name        string `json:"name"`
	Description string `json:"description"`

	Price PriceView `json:"price"`
}

type PriceView struct {
	Cents    uint   `json:"cents"`
	Currency string `json:"currency"`
}

func priceViewFromPrice(p price.Price) PriceView {
	return PriceView{p.Cents(), p.Currency()}
}

func products(repo shop.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		product, err := repo.ByID(shop.ID(chi.URLParam(r, "id")))

		if err != nil {
			_ = render.Render(w, r, httputils.ErrInternal(err))
			return
		}

		render.Respond(w, r, ProductView{
			string(product.ID()),
			product.Name(),
			product.Description(),
			priceViewFromPrice(product.Price()),
		})
	}
}
