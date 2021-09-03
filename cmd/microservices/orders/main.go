package main

import (
	"log"
	"net/http"
	"os"

	"monolith-microservice-shop/pkg/common/cmd"
)

func main() {

	log.Println("Starting orders microservice")

	router := cmd.CreateRouter()

	defer createService(router)

	server := &http.Server{Addr: os.Getenv("SHOP_ORDERS_SERVICE_BIND_ADDR"), Handler: router}
	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			panic(err)
		}
	}()

	<-cmd.Context().Done()

	log.Println("Closing orders microservice")

	if err := server.Close(); err != nil {
		panic(err)
	}
}
