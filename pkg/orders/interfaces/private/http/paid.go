package http

import (
	"net/http"

	httputils "monolith-microservice-shop/pkg/common/http"
	"monolith-microservice-shop/pkg/orders/application"
	"monolith-microservice-shop/pkg/orders/domain/orders"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

func paid(service application.OrdersService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cmd := application.MarkOrderAsPaidCommand{
			OrderID: orders.OrderID(chi.URLParam(r, "id")),
		}
		if err := service.MarkOrderAsPaid(cmd); err != nil {
			_ = render.Render(w, r, httputils.ErrInternal(err))
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}
