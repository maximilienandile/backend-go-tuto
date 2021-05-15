package cart

import (
	"fmt"

	"github.com/Rhymond/go-money"
)

type Cart struct {
	ID           string `json:"id"`
	CurrencyCode string `json:"currencyCode"`
	// key: productID
	// value : the item in the cart
	Items   map[string]Item `json:"items"`
	Version uint            `json:"version"`
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
	if c.Items == nil {
		c.Items = make(map[string]Item)
	}
	itemFound, found := c.Items[productID]
	if !found {
		// item not in the cart we have to add it
		if delta <= 0 {
			return fmt.Errorf("item not found in the cart, but delta is less or equal to zero, (we cannot add an item with a negative or zero quantity): %d", delta)
		}
		c.Items[productID] = Item{
			ID:       productID,
			Quantity: uint8(delta),
		}
	} else {
		// a product with this id is already in the cart
		// we found an entry in the map
		newQuantity := int(itemFound.Quantity) + delta
		if newQuantity < 0 {
			return fmt.Errorf("new quantity cannot be less than zero")
		} else if newQuantity > 0 {
			itemFound.Quantity = uint8(newQuantity)
			c.Items[productID] = itemFound
		} else {
			// equal to zero.
			// it means that I want to remove that from my cart
			delete(c.Items, productID)
		}

	}
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
