package product

import (
	"github.com/Rhymond/go-money"
	"github.com/maximilienandile/backend-go-tuto/internal/extMoney"
)

type Product struct {
	ID               string            `json:"id"`
	CategoryID       string            `json:"categoryId"`
	Name             string            `json:"name"`
	Image            string            `json:"image"`
	ShortDescription string            `json:"shortDescription"`
	Description      string            `json:"description"`
	PriceVATExcluded extMoney.ExtMoney `json:"priceVatExcluded"`
	VAT              extMoney.ExtMoney `json:"vat"`
	TotalPrice       extMoney.ExtMoney `json:"totalPrice"`
	// inventory
	Stock    uint `json:"stock"`
	Reserved uint `json:"reserved"`
	Version  uint `json:"version"`
}

type Amount struct {
	Money   *money.Money `json:"money"`
	Display string       `json:"display"`
}
