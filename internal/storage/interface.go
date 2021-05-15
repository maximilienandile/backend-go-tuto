package storage

import "github.com/maximilienandile/backend-go-tuto/internal/product"

type Storage interface {
	CreateProduct(product product.Product) error
}
