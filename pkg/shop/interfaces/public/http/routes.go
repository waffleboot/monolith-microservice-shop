package http

import "github.com/go-chi/chi"

func AddRoutes(router *chi.Mux, model productsReadModel) {
	router.Get("/products", GetAll(model))
}
