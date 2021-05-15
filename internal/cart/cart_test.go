package cart

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Rhymond/go-money"
)

func TestCart_TotalPriceVATInc(t *testing.T) {
	t.Run("nominal", func(t *testing.T) {
		// GIVEN
		items := map[string]Item{
			"42": {
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
		assert.NoError(t, err, "impossible to compute total price VAT included")
		expectedTotalPrice := money.New(100, "EUR")
		assert.Equal(t, expectedTotalPrice, actualTotalPrice)
	})
	t.Run("nominal with quantity greater than 1 and 2 items", func(t *testing.T) {
		// GIVEN
		items := map[string]Item{
			"42": {
				ID:               "42",
				ShortDescription: "A pair of socks",
				UnitPriceVATInc:  money.New(100, "EUR"),
				UnitPriceVATExc:  money.New(50, "EUR"),
				VAT:              money.New(50, "EUR"),
				Quantity:         1,
			},
			"43": {
				ID:               "43",
				ShortDescription: "A T-Shirt with a small gopher",
				UnitPriceVATInc:  money.New(3480, "EUR"),
				UnitPriceVATExc:  money.New(2900, "EUR"),
				VAT:              money.New(580, "EUR"),
				Quantity:         2,
			},
		}
		cart := Cart{
			ID:           "42",
			CurrencyCode: "EUR",
			Items:        items,
		}

		// WHEN
		actualTotalPriceVATInc, err := cart.TotalPriceVATInc()

		// THEN
		assert.NoError(t, err, "should have no error when total price VAT inc is computed")
		expectedTotalPriceVATINC := money.New(7060, "EUR")
		assert.Equal(t, expectedTotalPriceVATINC, actualTotalPriceVATInc)
	})
	t.Run("error case different currencies", func(t *testing.T) {
		items := map[string]Item{
			"42": {
				ID:               "42",
				ShortDescription: "A pair of socks",
				UnitPriceVATInc:  money.New(100, "USD"),
				UnitPriceVATExc:  money.New(50, "USD"),
				VAT:              money.New(50, "USD"),
				Quantity:         1,
			},
		}
		cart := Cart{
			ID:           "42",
			CurrencyCode: "EUR",
			Items:        items,
		}

		// WHEN
		_, err := cart.TotalPriceVATInc()

		// THEN
		assert.Error(t, err, "when I add an item with a currency X to a basket of currency Y the method TotalPriceVATInc should fail")
	})
}
