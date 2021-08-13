package main

import (
	shop_app "github.com/ThreeDotsLabs/monolith-microservice-shop/pkg/shop/application"
	shop_repo "github.com/ThreeDotsLabs/monolith-microservice-shop/pkg/shop/infrastructure/repo"
	shop_ipc "github.com/ThreeDotsLabs/monolith-microservice-shop/pkg/shop/interfaces/private/ipc"
)

func buildShopService() (shop_app.ProductsService, shop_ipc.ProductInterface, *shop_repo.MemoryRepository) {
	repo := shop_repo.NewMemoryRepository()
	service := shop_app.NewProductsService(repo, repo)
	return service, shop_ipc.NewProductInterface(repo), repo
}
