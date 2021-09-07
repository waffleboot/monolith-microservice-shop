package http

import (
	"net/http"

	httputils "monolith-microservice-shop/pkg/common/http"
	domain "monolith-microservice-shop/pkg/orders/domain/orders"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type OrderPaidView struct {
	ID     string `json:"id"`
	IsPaid bool   `json:"is_paid"`
}

func paid(repo domain.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		order, err := repo.ByID(domain.OrderID(chi.URLParam(r, "id")))
		if err != nil {
			_ = render.Render(w, r, httputils.ErrBadRequest(err))
			return
		}
		render.Respond(w, r, OrderPaidView{string(order.ID()), order.Paid()})
	}
}
