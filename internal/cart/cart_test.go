package cart

import (
	"testing"

	"github.com/Rhymond/go-money"
)

func TestCart_TotalPriceVATInc(t *testing.T) {
	// GIVEN
	items := []Item{
		{
			ID:               "42",
			ShortDescription: "A pair of socks",
			UnitPriceVATInc:  money.New(100, "EUR"),
			UnitPriceVATExc:  money.New(50, "EUR"),
			VAT:              money.New(50, "EUR"),
			Quantity:         1,
		},
	}
	cart := Cart{
		ID:           "42",
		CurrencyCode: "EUR",
		Items:        items,
	}
	// WHEN
	actualTotalPrice, err := cart.TotalPriceVATInc()

	// THEN
	if err != nil {
		t.Fail()
	}
	expectedTotalPrice := money.New(100, "EUR")
	isEqual, err := expectedTotalPrice.Equals(actualTotalPrice)
	if err != nil {
		t.Fail()
	}
	if !isEqual {
		t.Fail()
	}

}
