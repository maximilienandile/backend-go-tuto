package product

import "github.com/Rhymond/go-money"

type Product struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	Image            string `json:"image"`
	ShortDescription string `json:"shortDescription"`
	Description      string `json:"description"`
	PriceVATExcluded Amount `json:"priceVatExcluded"`
	VAT              Amount `json:"vat"`
	TotalPrice       Amount `json:"totalPrice"`
	// inventory
	Stock    uint `json:"stock"`
	Reserved uint `json:"reserved"`
	Version  uint `json:"version"`
}

type Amount struct {
	Money   *money.Money `json:"money"`
	Display string       `json:"display"`
}
