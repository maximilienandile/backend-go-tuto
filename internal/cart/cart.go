package cart

import (
	"fmt"

	"github.com/Rhymond/go-money"
)

type Cart struct {
	ID           string
	CurrencyCode string
	Items        []Item
}

func (c Cart) TotalPriceVATInc() (*money.Money, error) {
	totalPrice := money.New(0, c.CurrencyCode)
	for _, item := range c.Items {
		itemPrice := item.UnitPriceVATInc.Multiply(int64(item.Quantity))
		var err error
		totalPrice, err = totalPrice.Add(itemPrice)
		if err != nil {
			return nil, fmt.Errorf("impossible to add item price to total price: %w", err)
		}
	}
	return totalPrice, nil
}

func (c *Cart) UpsertItem(productID string, delta int) error {
	// handle the logic to add an item to a cart

	return nil
}

type Item struct {
	ID               string
	ShortDescription string
	Quantity         uint8
	UnitPriceVATExc  *money.Money
	VAT              *money.Money
	UnitPriceVATInc  *money.Money
}
