package product

import (
	"github.com/google/uuid"
	domain "github.com/ovitorvalente/go-clean-products/internal/domain/product"
)

type CreateProductUseCase struct {
	repo domain.Repository
}

func NewCreateProductUseCase(repo domain.Repository) *CreateProductUseCase {
	return &CreateProductUseCase{repo: repo}
}

func (uc *CreateProductUseCase) Execute(id, name string, price float64) (*domain.Product, error) {
	id = uuid.New().String()

	product, err := domain.NewProduct(id, name, price)
	if err != nil {
		return nil, err
	}

	err = uc.repo.Save(product)
	if err != nil {
		return nil, err
	}

	return product, nil
}
