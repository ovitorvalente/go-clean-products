package product

import "errors"

type Product struct {
	ID    string
	Name  string
	Price float64
}

func NewProduct(id, name string, price float64) (*Product, error) {
	if name == "" {
		return nil, errors.New("Nome do produto é obrigatorio.")
	}

	if price <= 0 {
		return nil, errors.New("O preço do produto não pode ser menor ou igual a zero.")
	}

	return &Product{
		ID:    id,
		Name:  name,
		Price: price,
	}, nil
}
