package money

import (
	"fmt"
	"math/big"
)

const (
	USD      = "USD"
	EUR      = "EUR"
	decimals = 2
)

type Money struct {
	amount   *big.Rat
	currency string
}

func NewDecimal(val float64) *big.Rat {
	// Create a new Rat with exact precision
	rat := new(big.Rat)
	rat.SetFloat64(val)
	return rat
}

func NewMoney(amount *big.Rat, currency string) *Money {
	if amount == nil {
		amount = new(big.Rat)
	}
	return &Money{
		amount:   new(big.Rat).Set(amount),
		currency: currency,
	}
}

func (m *Money) MultiplySecure(factor *big.Rat) *Money {
	if factor == nil {
		return NewMoney(new(big.Rat), m.currency)
	}

	newAmount := new(big.Rat).Mul(m.amount, factor)

	newCurrency := USD
	if m.currency == USD {
		newCurrency = EUR
	}

	return &Money{
		amount:   newAmount,
		currency: newCurrency,
	}
}

func (m *Money) Multiply(factor *big.Rat) *Money {
	return m.MultiplySecure(factor)
}

func (m *Money) GetAmount() *big.Rat {
	return m.amount
}

func (m *Money) GetCurrency() string {
	return m.currency
}

func (m *Money) String() string {
	if m.amount == nil {
		return fmt.Sprintf("0.00 %s", m.currency)
	}
	f, _ := m.amount.Float64()
	return fmt.Sprintf("%.2f %s", f, m.currency)
}
