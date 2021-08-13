package main

import (
	"fmt"
	"log"
	"os"

	"github.com/waffleboot/monolith-microservice-shop/pkg/common/cmd"
	"github.com/waffleboot/monolith-microservice-shop/pkg/payments/application"
	payments_infra_orders "github.com/waffleboot/monolith-microservice-shop/pkg/payments/infrastructure/orders"
	"github.com/waffleboot/monolith-microservice-shop/pkg/payments/interfaces/amqp"
)

func main() {
	log.Println("Starting payments microservice")
	defer log.Println("Closing payments microservice")

	ctx := cmd.Context()

	paymentsInterface := createService()
	if err := paymentsInterface.Run(ctx); err != nil {
		panic(err)
	}
}

func createService() amqp.PaymentsAMQP {
	cmd.WaitForService(os.Getenv("SHOP_RABBITMQ_ADDR"))

	paymentsService := application.NewPaymentsService(
		payments_infra_orders.NewHTTPClient(os.Getenv("SHOP_ORDERS_SERVICE_ADDR")),
	)

	paymentsInterface, err := amqp.NewPaymentsAMQP(
		fmt.Sprintf("amqp://%s/", os.Getenv("SHOP_RABBITMQ_ADDR")),
		os.Getenv("SHOP_RABBITMQ_ORDERS_TO_PAY_QUEUE"),
		paymentsService,
	)
	if err != nil {
		panic(err)
	}

	return paymentsInterface
}
