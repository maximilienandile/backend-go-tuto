package storage

import (
	"github.com/maximilienandile/backend-go-tuto/internal/category"
	"github.com/maximilienandile/backend-go-tuto/internal/product"
)

type UpdateProductInput struct {
	ProductID        string
	Name             string
	Image            string
	ShortDescription string
	Description      string
	PriceVATExcluded product.Amount
	VAT              product.Amount
	TotalPrice       product.Amount
}

type Storage interface {
	CreateProduct(product product.Product) error
	Products() ([]product.Product, error)
	CreateCategory(category category.Category) error
	Categories() ([]category.Category, error)
	UpdateInventory(productID string, delta int) error
	UpdateProduct(input UpdateProductInput) error
}
