package main

import (
	"log"
	"net/http"
	"os"

	"monolith-microservice-shop/pkg/common/cmd"
	. "monolith-microservice-shop/pkg/shop"
	. "monolith-microservice-shop/pkg/shop/application"
	. "monolith-microservice-shop/pkg/shop/infrastructure/repo"
	private_http "monolith-microservice-shop/pkg/shop/interfaces/private/http"
	public_http "monolith-microservice-shop/pkg/shop/interfaces/public/http"

	"github.com/go-chi/chi"
)

func main() {
	log.Println("Starting shop microservice")

	router := cmd.CreateRouter()
	createService(router)
	server := &http.Server{Addr: os.Getenv("SHOP_SHOP_SERVICE_BIND_ADDR"), Handler: router}
	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			panic(err)
		}
	}()

	<-cmd.Context().Done()
	log.Println("Closing shop microservice")

	if err := server.Close(); err != nil {
		panic(err)
	}
}

func createService(router *chi.Mux) {
	repo := NewMemoryRepository()
	service := NewService(repo, repo)
	if err := LoadShopFixtures(service); err != nil {
		panic(err)
	}
	public_http.AddRoutes(router, repo)
	private_http.AddRoutes(router, repo)
}
