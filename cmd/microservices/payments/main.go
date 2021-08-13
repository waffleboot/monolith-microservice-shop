package main

import (
	"fmt"
	"log"
	"os"

	"monolith-microservice-shop/pkg/common/cmd"
	"monolith-microservice-shop/pkg/payments/application"
	payments_infra_orders "monolith-microservice-shop/pkg/payments/infrastructure/orders"
	"monolith-microservice-shop/pkg/payments/interfaces/amqp"
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
