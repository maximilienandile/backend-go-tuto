package inter

import "github.com/maximilienandile/backend-go-tuto/internal/product"

type Storage interface {
	CreateProduct(product product.Product) error
}

type MySQLStorage struct {
}

func (m MySQLStorage) CreateProduct(product product.Product) error {
	// implement me
	return nil
}

type DynamoDBStorage struct {
}

func (m DynamoDBStorage) CreateProduct(product product.Product) error {
	// implement me
	return nil
}
