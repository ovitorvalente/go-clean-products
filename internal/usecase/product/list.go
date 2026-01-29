package product

import domain "github.com/ovitorvalente/go-clean-products/internal/domain/product"

type ListProductsUseCase struct {
	repo domain.Repository
}

func NewListProductsUseCase(repo domain.Repository) *ListProductsUseCase {
	return &ListProductsUseCase{repo: repo}
}

func (uc *ListProductsUseCase) Execute() ([]*domain.Product, error) {
	return uc.repo.FindAll()
}
