package cart

import (
	"github.com/Rhymond/go-money"
	"github.com/maximilienandile/backend-go-tuto/internal/extMoney"
)

type Item struct {
	ID               string            `json:"id"`
	Title            string            `json:"title"`
	ShortDescription string            `json:"shortDescription"`
	Quantity         uint8             `json:"quantity"`
	UnitPriceVATExc  extMoney.ExtMoney `json:"unitPriceVATExc"`
	UnitVAT          extMoney.ExtMoney `json:"unitVAT"`
	UnitPriceVATInc  extMoney.ExtMoney `json:"unitPriceVATInc"`
	// total price for the line in the cart
	TotalPriceVATExc extMoney.ExtMoney `json:"totalPriceVATExc"`
	TotalVAT         extMoney.ExtMoney `json:"totalVat"`
	TotalPriceVATInc extMoney.ExtMoney `json:"totalPriceVATInc"`
}

// ComputeTotalVAT will compute the total VAT for an item in the cart
func (i Item) ComputeTotalVAT() *money.Money {
	return i.UnitVAT.ToMoney().Multiply(int64(i.Quantity))
}

func (i Item) ComputeTotalPriceVATInc() *money.Money {
	return i.UnitPriceVATInc.ToMoney().Multiply(int64(i.Quantity))
}
