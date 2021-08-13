package http

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	httputils "github.com/waffleboot/monolith-microservice-shop/pkg/common/http"
	"github.com/waffleboot/monolith-microservice-shop/pkg/common/price"
	shop "github.com/waffleboot/monolith-microservice-shop/pkg/shop/domain/products"
)

func AddRoutes(router *chi.Mux, repo shop.Repository) {
	resource := productsResource{repo}
	router.Get("/products/{id}", resource.Get)
}

type productsResource struct {
	repo shop.Repository
}

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

func (p productsResource) Get(w http.ResponseWriter, r *http.Request) {
	product, err := p.repo.ByID(shop.ID(chi.URLParam(r, "id")))

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
