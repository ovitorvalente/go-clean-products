package main

import (
	"net/http"

	repo "github.com/ovitorvalente/go-clean-products/internal/infra/repository"
	handler "github.com/ovitorvalente/go-clean-products/internal/interfaces/http"
	usecase "github.com/ovitorvalente/go-clean-products/internal/usecase/product"
)

func main() {
	productRepo := repo.NewProductMemoryRepository()
	createUC := usecase.NewCreateProductUseCase(productRepo)
	listUC := usecase.NewListProductsUseCase(productRepo)

	productHandler := handler.NewProductHandler(createUC, listUC)

	http.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			productHandler.Create(w, r)
		case http.MethodGet:
			productHandler.List(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.ListenAndServe(":8080", nil)
}
