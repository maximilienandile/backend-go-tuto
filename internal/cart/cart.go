package cart

import (
	"fmt"

	"github.com/maximilienandile/backend-go-tuto/internal/extMoney"

	"github.com/maximilienandile/backend-go-tuto/internal/product"

	"github.com/Rhymond/go-money"
)

type Cart struct {
	CurrencyCode string `json:"currencyCode"`
	// key: productID
	// value : the item in the cart
	Items       map[string]Item   `json:"items"`
	Version     uint              `json:"version"`
	TotalVATInc extMoney.ExtMoney `json:"totalPriceVATInc"`
	TotalVAT    extMoney.ExtMoney `json:"totalVAT"`
	TotalVATExc extMoney.ExtMoney `json:"totalPriceVATExc"`
}

func (c *Cart) ComputePrices() error {
	if c.CurrencyCode == "" {
		c.CurrencyCode = "EUR"
	}
	// compute total price VAT INC
	totalPriceVATInc, err := c.TotalPriceVATInc()
	if err != nil {
		return fmt.Errorf("impossible to compute total VAT Inc: %w", err)
	}
	c.TotalVATInc = extMoney.FromMoney(totalPriceVATInc)
	// VAT
	totalVAT, err := c.VAT()
	if err != nil {
		return fmt.Errorf("impossible to compute total VAT: %w", err)
	}
	c.TotalVAT = extMoney.FromMoney(totalVAT)
	// total price VAT Exc = Total VAT INC - VAT
	totalPriceVATExc, err := totalPriceVATInc.Subtract(totalVAT)
	if err != nil {
		return fmt.Errorf("impossible to compute total VAT exc, cannot substract: %w", err)
	}
	c.TotalVATExc = extMoney.FromMoney(totalPriceVATExc)
	return nil
}

func (c Cart) TotalPriceVATInc() (*money.Money, error) {
	totalPrice := money.New(0, c.CurrencyCode)
	for _, item := range c.Items {
		itemPrice := item.UnitPriceVATInc.ToMoney().Multiply(int64(item.Quantity))
		var err error
		totalPrice, err = totalPrice.Add(itemPrice)
		if err != nil {
			return nil, fmt.Errorf("impossible to add item price to total price: %w", err)
		}
	}
	return totalPrice, nil
}

func (c Cart) VAT() (*money.Money, error) {
	totalVAT := money.New(0, c.CurrencyCode)
	for _, item := range c.Items {
		itemVAT := item.VAT.ToMoney().Multiply(int64(item.Quantity))
		var err error
		totalVAT, err = totalVAT.Add(itemVAT)
		if err != nil {
			return nil, fmt.Errorf("impossible to add item price to total price: %w", err)
		}
	}
	return totalVAT, nil
}

func (c *Cart) UpsertItem(product product.Product, delta int) error {
	if c.Items == nil {
		c.Items = make(map[string]Item)
	}
	itemFound, found := c.Items[product.ID]
	if !found {
		// item not in the cart we have to add it
		if delta <= 0 {
			return fmt.Errorf("item not found in the cart, but delta is less or equal to zero, (we cannot add an item with a negative or zero quantity): %d", delta)
		}
		c.Items[product.ID] = Item{
			ID:               product.ID,
			Quantity:         uint8(delta),
			ShortDescription: product.ShortDescription,
			UnitPriceVATExc:  product.PriceVATExcluded,
			VAT:              product.VAT,
			UnitPriceVATInc:  product.PriceVATExcluded,
		}
	} else {
		// a product with this id is already in the cart
		// we found an entry in the map
		newQuantity := int(itemFound.Quantity) + delta
		if newQuantity < 0 {
			return fmt.Errorf("new quantity cannot be less than zero")
		} else if newQuantity > 0 {
			itemFound.Quantity = uint8(newQuantity)
			c.Items[product.ID] = itemFound
		} else {
			// equal to zero.
			// it means that I want to remove that from my cart
			delete(c.Items, product.ID)
		}

	}
	return nil
}

type Item struct {
	ID               string            `json:"id"`
	ShortDescription string            `json:"shortDescription"`
	Quantity         uint8             `json:"quantity"`
	UnitPriceVATExc  extMoney.ExtMoney `json:"unitPriceVATExc"`
	VAT              extMoney.ExtMoney `json:"vat"`
	UnitPriceVATInc  extMoney.ExtMoney `json:"UnitPriceVATInc"`
}
