package http

import (
	"encoding/json"
	"net/http"

	usecase "github.com/ovitorvalente/go-clean-products/internal/usecase/product"
)

type ProductHandler struct {
	createUC *usecase.CreateProductUseCase
	listUC   *usecase.ListProductsUseCase
}

func NewProductHandler(createUC *usecase.CreateProductUseCase, listUC *usecase.ListProductsUseCase) *ProductHandler {
	return &ProductHandler{
		createUC: createUC,
		listUC:   listUC,
	}
}

func (h *ProductHandler) Create(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name  string  `json:"name"`
		Price float64 `json:"price"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	product, err := h.createUC.Execute(input.Name, input.Price)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}

func (h *ProductHandler) List(w http.ResponseWriter, r *http.Request) {
	products, err := h.listUC.Execute()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}
