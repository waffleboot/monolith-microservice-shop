package main

import (
	"fmt"
	"log"
	"os"

	"monolith-microservice-shop/pkg/common/cmd"
	"monolith-microservice-shop/pkg/orders/application"
	payments "monolith-microservice-shop/pkg/orders/infrastructure/payments"
	orders_repo "monolith-microservice-shop/pkg/orders/infrastructure/repo"
	shop "monolith-microservice-shop/pkg/orders/infrastructure/shop"
	private_http "monolith-microservice-shop/pkg/orders/interfaces/private/http"
	public_http "monolith-microservice-shop/pkg/orders/interfaces/public/http"

	"github.com/go-chi/chi"
)

func createService(router *chi.Mux) (done func()) {

	cmd.WaitForService(os.Getenv("SHOP_RABBITMQ_ADDR"))

	payments, err := payments.NewAMQPService(
		fmt.Sprintf("amqp://%s/", os.Getenv("SHOP_RABBITMQ_ADDR")),
		os.Getenv("SHOP_RABBITMQ_ORDERS_TO_PAY_QUEUE"),
	)
	if err != nil {
		panic(err)
	}

	repo := orders_repo.NewMemoryRepository()

	service := application.NewOrdersService(
		shop.NewHTTPClient(os.Getenv("SHOP_SHOP_SERVICE_ADDR")),
		payments,
		repo,
	)

	public_http.AddRoutes(router, service, repo)
	private_http.AddRoutes(router, service)

	return func() {
		if err := payments.Close(); err != nil {
			log.Printf("cannot close orders queue: %s", err)
		}
	}
}
