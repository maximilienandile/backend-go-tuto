package storage

import "github.com/maximilienandile/backend-go-tuto/internal/product"

type MySQL struct {
}

func (m *MySQL) CreateProduct(product product.Product) error {
	panic("implement me")
}

func (m MySQL) String() string {
	return "MYSQL"
}
