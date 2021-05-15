package cart

import "github.com/Rhymond/go-money"

type Cart struct {
	ID    string
	Items []Item
}

func (c Cart) TotalPrice() (*money.Money, error) {

}

type Item struct {
	ID               string
	ShortDescription string
	Quantity         uint8
	UnitPriceVATExc  *money.Money
	VAT              *money.Money
	UnitPriceVATInc  *money.Money
}
