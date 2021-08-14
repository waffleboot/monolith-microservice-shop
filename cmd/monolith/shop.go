package main

import (
	"monolith-microservice-shop/pkg/shop/application"
	"monolith-microservice-shop/pkg/shop/infrastructure/repo"
	"monolith-microservice-shop/pkg/shop/interfaces/private/ipc"
)

func buildShopService() (application.ProductsService, ipc.ProductInterface, *repo.MemoryRepository) {
	repo := repo.NewMemoryRepository()
	service := application.NewProductsService(repo, repo)
	return service, ipc.NewProductInterface(repo), repo
}
