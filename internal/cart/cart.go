package cart

import (
	"fmt"

	"github.com/Rhymond/go-money"
)

type Cart struct {
	ID           string `json:"id"`
	CurrencyCode string `json:"currencyCode"`
	// key : productID
	// Value : Item
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
	itemRetrieved, found := c.Items[productID]
	if !found {
		if delta > 0 {
			c.Items[productID] = Item{
				ID:               productID,
				ShortDescription: "",
				Quantity:         uint8(delta),
				UnitPriceVATExc:  nil,
				VAT:              nil,
				UnitPriceVATInc:  nil,
			}
		} else {
			return fmt.Errorf("delta is less than zero but item do not exists")
		}
	} else {
		// item is in the map
		newQuantity := int(itemRetrieved.Quantity) + delta
		if newQuantity < 0 {
			return fmt.Errorf("impossible to have a quantity less than zero")
		} else if newQuantity > 0 {
			itemRetrieved.Quantity = uint8(newQuantity)
			c.Items[productID] = itemRetrieved
		} else {
			// quantity is 0
			// remove from cart
			delete(c.Items, productID)
		}
	}
	return nil
}

type Item struct {
	ID               string       `json:"id"`
	ShortDescription string       `json:"shortDescription"`
	Quantity         uint8        `json:"quantity"`
	UnitPriceVATExc  *money.Money `json:"unitPriceVATExc"`
	VAT              *money.Money `json:"vat"`
	UnitPriceVATInc  *money.Money `json:"unitPriceVATInc"`
}
