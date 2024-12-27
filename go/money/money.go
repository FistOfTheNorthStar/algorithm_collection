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

// Money represents a monetary amount with its currency
type Money struct {
	amount   *big.Rat
	currency string
}

// NewMoney creates a new Money instance
func NewMoney(amount float64, currency string) *Money {
	// Convert float64 to *big.Rat
	rat := new(big.Rat).SetFloat64(amount)
	return &Money{
		amount:   rat,
		currency: currency,
	}
}

// Multiply converts the currency and applies the conversion factor
func (m *Money) Multiply(factor float64) *Money {
	// Convert factor to *big.Rat
	factorRat := new(big.Rat).SetFloat64(factor)

	// Multiply amounts
	newAmount := new(big.Rat).Mul(m.amount, factorRat)

	// Round to 2 decimal places
	newAmount = round(newAmount, decimals)

	newCurrency := USD
	if m.currency == USD {
		newCurrency = EUR
	}

	return &Money{
		amount:   newAmount,
		currency: newCurrency,
	}
}

// round to specified decimal places
func round(x *big.Rat, places int) *big.Rat {
	// Multiply by 10^places
	scale := new(big.Int).Exp(
		big.NewInt(10),
		big.NewInt(int64(places)),
		nil,
	)
	scaleRat := new(big.Rat).SetInt(scale)

	// Multiply, round, then divide
	result := new(big.Rat).Mul(x, scaleRat)
	rounded := new(big.Rat).SetFloat64(
		float64(result.Num().Int64()) / float64(result.Denom().Int64()),
	)
	return new(big.Rat).Quo(rounded, scaleRat)
}

// GetAmount returns the monetary amount as float64
func (m *Money) GetAmount() float64 {
	f, _ := m.amount.Float64()
	return f
}

// GetRawAmount returns the exact *big.Rat amount
func (m *Money) GetRawAmount() *big.Rat {
	return new(big.Rat).Set(m.amount)
}

// GetCurrency returns the currency code
func (m *Money) GetCurrency() string {
	return m.currency
}

// String returns formatted money string
func (m *Money) String() string {
	f, _ := m.amount.Float64()
	return fmt.Sprintf("%.2f %s", f, m.currency)
}
