package http

import (
	"net/http"

	httputils "monolith-microservice-shop/pkg/common/http"
	"monolith-microservice-shop/pkg/orders/application"
	domain "monolith-microservice-shop/pkg/orders/domain/orders"

	"github.com/go-chi/render"
	uuid "github.com/satori/go.uuid"
)

type PostOrderRequest struct {
	ProductID string           `json:"product_id"`
	Address   PostOrderAddress `json:"address"`
}

type PostOrderAddress struct {
	Name     string `json:"name"`
	Street   string `json:"street"`
	City     string `json:"city"`
	PostCode string `json:"post_code"`
	Country  string `json:"country"`
}

type PostOrdersResponse struct {
	OrderID string
}

func orders(service application.OrdersService, repo domain.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := PostOrderRequest{}
		if err := render.Decode(r, &req); err != nil {
			_ = render.Render(w, r, httputils.ErrBadRequest(err))
			return
		}
		cmd := application.PlaceOrderCommand{
			OrderID:   domain.OrderID(uuid.NewV1().String()),
			ProductID: domain.ProductID(req.ProductID),
			Address:   application.PlaceOrderCommandAddress(req.Address),
		}
		if err := service.PlaceOrder(cmd, repo); err != nil {
			_ = render.Render(w, r, httputils.ErrInternal(err))
			return
		}
		w.WriteHeader(http.StatusOK)
		render.JSON(w, r, PostOrdersResponse{
			OrderID: string(cmd.OrderID),
		})
	}
}
