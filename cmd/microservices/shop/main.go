package main

import (
	"log"
	"net/http"
	"os"

	"monolith-microservice-shop/pkg/common/cmd"
	"monolith-microservice-shop/pkg/shop"
	"monolith-microservice-shop/pkg/shop/application"
	shop_repo "monolith-microservice-shop/pkg/shop/infrastructure/repo"
	shop_private_http "monolith-microservice-shop/pkg/shop/interfaces/private/http"
	shop_public_http "monolith-microservice-shop/pkg/shop/interfaces/public/http"

	"github.com/go-chi/chi"
)

func main() {
	log.Println("Starting shop microservice")

	ctx := cmd.Context()

	r := createService()
	server := &http.Server{Addr: os.Getenv("SHOP_SHOP_SERVICE_BIND_ADDR"), Handler: r}
	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			panic(err)
		}
	}()

	<-ctx.Done()
	log.Println("Closing shop microservice")

	if err := server.Close(); err != nil {
		panic(err)
	}
}

func createService() *chi.Mux {
	repo := shop_repo.NewMemoryRepository()
	service := application.NewProductsService(repo, repo)

	if err := shop.LoadShopFixtures(service); err != nil {
		panic(err)
	}

	r := cmd.CreateRouter()

	shop_public_http.AddRoutes(r, repo)
	shop_private_http.AddRoutes(r, repo)

	return r
}
