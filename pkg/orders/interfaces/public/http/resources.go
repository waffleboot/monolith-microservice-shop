package http

import (
	"net/http"

	httputils "github.com/ThreeDotsLabs/monolith-microservice-shop/pkg/common/http"
	"github.com/ThreeDotsLabs/monolith-microservice-shop/pkg/orders/application"
	"github.com/ThreeDotsLabs/monolith-microservice-shop/pkg/orders/domain/orders"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	uuid "github.com/satori/go.uuid"
)

type ordersResource struct {
	service    application.OrdersService
	repository orders.Repository
}

func (o ordersResource) Post(w http.ResponseWriter, r *http.Request) {
	req := PostOrderRequest{}
	if err := render.Decode(r, &req); err != nil {
		_ = render.Render(w, r, httputils.ErrBadRequest(err))
		return
	}
	cmd := application.PlaceOrderCommand{
		OrderID:   orders.ID(uuid.NewV1().String()),
		ProductID: req.ProductID,
		Address:   application.PlaceOrderCommandAddress(req.Address),
	}
	if err := o.service.PlaceOrder(cmd); err != nil {
		_ = render.Render(w, r, httputils.ErrInternal(err))
		return
	}
	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, PostOrdersResponse{
		OrderID: string(cmd.OrderID),
	})
}

func (o ordersResource) GetPaid(w http.ResponseWriter, r *http.Request) {
	order, err := o.repository.ByID(orders.ID(chi.URLParam(r, "id")))
	if err != nil {
		_ = render.Render(w, r, httputils.ErrBadRequest(err))
		return
	}
	render.Respond(w, r, OrderPaidView{string(order.ID()), order.Paid()})
}
