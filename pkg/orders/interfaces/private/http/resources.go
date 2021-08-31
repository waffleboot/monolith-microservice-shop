package http

import (
	"net/http"

	httputils "monolith-microservice-shop/pkg/common/http"
	"monolith-microservice-shop/pkg/orders/application"
	"monolith-microservice-shop/pkg/orders/domain/orders"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type ordersEndpoint struct {
	service application.OrdersService
}

func (o ordersEndpoint) paid(w http.ResponseWriter, r *http.Request) {
	cmd := application.MarkOrderAsPaidCommand{
		OrderID: orders.ID(chi.URLParam(r, "id")),
	}
	if err := o.service.MarkOrderAsPaid(cmd); err != nil {
		_ = render.Render(w, r, httputils.ErrInternal(err))
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
