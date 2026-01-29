package repository

import (
	domain "github.com/ovitorvalente/go-clean-products/internal/domain/product"
)

type ProductMemoryRepository struct {
	products []*domain.Product
}

func NewProductMemoryRepository() *ProductMemoryRepository {
	return &ProductMemoryRepository{
		products: []*domain.Product{},
	}
}

func (r *ProductMemoryRepository) Save(product *domain.Product) error {
	r.products = append(r.products, product)
	return nil
}

func (r *ProductMemoryRepository) FindAll() ([]*domain.Product, error) {
	return r.products, nil
}
