package main

import (
	shop_app "monolith-microservice-shop/pkg/shop/application"
	shop_repo "monolith-microservice-shop/pkg/shop/infrastructure/repo"
	shop_ipc "monolith-microservice-shop/pkg/shop/interfaces/private/ipc"
)

func buildShopService() (shop_app.ProductsService, shop_ipc.ProductInterface, *shop_repo.MemoryRepository) {
	repo := shop_repo.NewMemoryRepository()
	service := shop_app.NewProductsService(repo, repo)
	return service, shop_ipc.NewProductInterface(repo), repo
}
