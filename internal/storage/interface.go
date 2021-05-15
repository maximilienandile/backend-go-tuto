package storage

import (
	"github.com/maximilienandile/backend-go-tuto/internal/cart"
	"github.com/maximilienandile/backend-go-tuto/internal/category"
	"github.com/maximilienandile/backend-go-tuto/internal/product"
)

type Storage interface {
	CreateProduct(product product.Product) error
	Products() ([]product.Product, error)
	CreateCategory(category category.Category) error
	Categories() ([]category.Category, error)
	UpdateInventory(productID string, delta int) error
	CreateOrUpdateCart(userID string, productID string, delta int) (cart.Cart, error)
	GetCart(userID string) (cart.Cart, error)
}
