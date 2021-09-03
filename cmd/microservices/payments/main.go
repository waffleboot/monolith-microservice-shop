package main

import (
	"fmt"
	"log"
	"os"

	"monolith-microservice-shop/pkg/common/cmd"
	payments_app "monolith-microservice-shop/pkg/payments/application"
	payments_orders "monolith-microservice-shop/pkg/payments/infrastructure/orders"
	"monolith-microservice-shop/pkg/payments/interfaces/amqp"
)

func main() {
	log.Println("Starting payments microservice")
	defer log.Println("Closing payments microservice")
	if err := createService().Run(cmd.Context()); err != nil {
		panic(err)
	}
}

func createService() amqp.Runner {
	cmd.WaitForService(os.Getenv("SHOP_RABBITMQ_ADDR"))

	service := payments_app.NewService(
		payments_orders.NewHTTPClient(os.Getenv("SHOP_ORDERS_SERVICE_ADDR")),
	)

	runner, err := amqp.NewRunner(
		fmt.Sprintf("amqp://%s/", os.Getenv("SHOP_RABBITMQ_ADDR")),
		os.Getenv("SHOP_RABBITMQ_ORDERS_TO_PAY_QUEUE"),
		service,
	)
	if err != nil {
		panic(err)
	}

	return runner
}
