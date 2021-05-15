package extMoney

import "github.com/Rhymond/go-money"

type ExtMoney struct {
	Amount   int64  `json:"amount"`
	Currency string `json:"currency"`
	Display  string `json:"display"`
}

func (m ExtMoney) ToMoney() *money.Money {
	return money.New(m.Amount, m.Currency)
}

func FromMoney(m *money.Money) ExtMoney {
	return ExtMoney{
		Amount:   m.Amount(),
		Currency: m.Currency().Code,
		Display:  m.Display(),
	}
}
