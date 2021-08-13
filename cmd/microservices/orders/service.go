package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-chi/chi"
	"github.com/waffleboot/monolith-microservice-shop/pkg/common/cmd"
	"github.com/waffleboot/monolith-microservice-shop/pkg/orders/application"
	payments "github.com/waffleboot/monolith-microservice-shop/pkg/orders/infrastructure/payments"
	orders_repo "github.com/waffleboot/monolith-microservice-shop/pkg/orders/infrastructure/repo"
	shop "github.com/waffleboot/monolith-microservice-shop/pkg/orders/infrastructure/shop"
	private_http "github.com/waffleboot/monolith-microservice-shop/pkg/orders/interfaces/private/http"
	public_http "github.com/waffleboot/monolith-microservice-shop/pkg/orders/interfaces/public/http"
)

func createService() (router *chi.Mux, done func()) {

	cmd.WaitForService(os.Getenv("SHOP_RABBITMQ_ADDR"))

	shop := shop.NewHTTPClient(os.Getenv("SHOP_SHOP_SERVICE_ADDR"))

	payments, err := payments.NewAMQPService(
		fmt.Sprintf("amqp://%s/", os.Getenv("SHOP_RABBITMQ_ADDR")),
		os.Getenv("SHOP_RABBITMQ_ORDERS_TO_PAY_QUEUE"),
	)
	if err != nil {
		panic(err)
	}

	repo := orders_repo.NewMemoryRepository()

	service := application.NewOrdersService(
		shop,
		payments,
		repo,
	)

	r := cmd.CreateRouter()

	public_http.AddRoutes(r, service, repo)
	private_http.AddRoutes(r, service)

	return r, func() {
		if err := payments.Close(); err != nil {
			log.Printf("cannot close orders queue: %s", err)
		}
	}
}
