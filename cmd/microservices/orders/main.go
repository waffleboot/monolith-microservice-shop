package main

import (
	"log"
	"net/http"
	"os"

	"github.com/waffleboot/monolith-microservice-shop/pkg/common/cmd"
)

func main() {

	log.Println("Starting orders microservice")

	ctx := cmd.Context()

	r, cancel := createService()
	defer cancel()

	server := &http.Server{Addr: os.Getenv("SHOP_ORDERS_SERVICE_BIND_ADDR"), Handler: r}
	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			panic(err)
		}
	}()

	<-ctx.Done()

	log.Println("Closing orders microservice")

	if err := server.Close(); err != nil {
		panic(err)
	}
}
