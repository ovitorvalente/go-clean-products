package product

type Repository interface {
	Save(product *Product) error
	FindAll() ([]*Product, error)
}
