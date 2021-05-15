package cart

import (
	"testing"

	"github.com/maximilienandile/backend-go-tuto/internal/extMoney"

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
				UnitPriceVATInc:  extMoney.FromMoney(money.New(100, "EUR")),
				UnitPriceVATExc:  extMoney.FromMoney(money.New(50, "EUR")),
				UnitVAT:          extMoney.FromMoney(money.New(50, "EUR")),
				Quantity:         1,
			},
		}
		cart := Cart{
			CurrencyCode: "EUR",
			Items:        items,
		}
		// WHEN
		actualTotalPrice, err := cart.TotalPriceVATInc()

		// THEN
		assert.NoError(t, err, "impossible to compute total price UnitVAT included")
		expectedTotalPrice := money.New(100, "EUR")
		assert.Equal(t, expectedTotalPrice, actualTotalPrice)
	})
	t.Run("nominal with quantity greater than 1 and 2 items", func(t *testing.T) {
		// GIVEN
		items := map[string]Item{
			"42": {
				ID:               "42",
				ShortDescription: "A pair of socks",
				UnitPriceVATInc:  extMoney.FromMoney(money.New(100, "EUR")),
				UnitPriceVATExc:  extMoney.FromMoney(money.New(50, "EUR")),
				UnitVAT:          extMoney.FromMoney(money.New(50, "EUR")),
				Quantity:         1,
			},
			"43": {
				ID:               "43",
				ShortDescription: "A T-Shirt with a small gopher",
				UnitPriceVATInc:  extMoney.FromMoney(money.New(3480, "EUR")),
				UnitPriceVATExc:  extMoney.FromMoney(money.New(2900, "EUR")),
				UnitVAT:          extMoney.FromMoney(money.New(580, "EUR")),
				Quantity:         2,
			},
		}
		cart := Cart{
			CurrencyCode: "EUR",
			Items:        items,
		}

		// WHEN
		actualTotalPriceVATInc, err := cart.TotalPriceVATInc()

		// THEN
		assert.NoError(t, err, "should have no error when total price UnitVAT inc is computed")
		expectedTotalPriceVATINC := money.New(7060, "EUR")
		assert.Equal(t, expectedTotalPriceVATINC, actualTotalPriceVATInc)
	})
	t.Run("error case different currencies", func(t *testing.T) {
		items := map[string]Item{
			"42": {
				ID:               "42",
				ShortDescription: "A pair of socks",
				UnitPriceVATInc:  extMoney.FromMoney(money.New(100, "USD")),
				UnitPriceVATExc:  extMoney.FromMoney(money.New(50, "USD")),
				UnitVAT:          extMoney.FromMoney(money.New(50, "USD")),
				Quantity:         1,
			},
		}
		cart := Cart{
			CurrencyCode: "EUR",
			Items:        items,
		}

		// WHEN
		_, err := cart.TotalPriceVATInc()

		// THEN
		assert.Error(t, err, "when I add an item with a currency X to a basket of currency Y the method TotalPriceVATInc should fail")
	})
}
