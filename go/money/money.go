package money

import (
	"fmt"
	"math"
)

const (
	USD      = "USD"
	EUR      = "EUR"
	decimals = 2
)

// Decimal represents a fixed-point decimal number
type Decimal struct {
	value int64 // The number without decimal point
	scale uint  // Number of decimal places
}

// Money represents a monetary amount with its currency
type Money struct {
	amount   Decimal
	currency string
}

// NewDecimal creates a new Decimal from a float64
func NewDecimal(v float64) Decimal {
	scale := uint(decimals)
	value := int64(math.Round(v * math.Pow10(int(scale))))
	return Decimal{value: value, scale: scale}
}

// NewDecimalFromString creates a new Decimal from a string
func NewDecimalFromString(s string) (Decimal, error) {
	var f float64
	_, err := fmt.Sscanf(s, "%f", &f)
	if err != nil {
		return Decimal{}, fmt.Errorf("invalid decimal string: %v", err)
	}
	return NewDecimal(f), nil
}

// Mul multiplies two decimals
func (d Decimal) Mul(other Decimal) Decimal {
	newValue := d.value * other.value
	newScale := d.scale + other.scale
	// Adjust scale back to 2 decimal places
	if newScale > decimals {
		newValue = newValue / int64(math.Pow10(int(newScale-decimals)))
	}
	return Decimal{value: newValue, scale: decimals}
}

// Round rounds to specified decimal places
func (d Decimal) Round(places uint) Decimal {
	if d.scale == places {
		return d
	}
	value := d.value
	if d.scale > places {
		div := int64(math.Pow10(int(d.scale - places)))
		value = (value + div/2) / div
	} else {
		value *= int64(math.Pow10(int(places - d.scale)))
	}
	return Decimal{value: value, scale: places}
}

// String converts decimal to string
func (d Decimal) String() string {
	str := fmt.Sprintf("%d", d.value)
	l := len(str)
	if l <= int(d.scale) {
		// Pad with zeros
		str = fmt.Sprintf("%0*d", int(d.scale)+1, d.value)
		l = len(str)
	}
	decimalPos := l - int(d.scale)
	if d.scale == 0 {
		return str
	}
	return str[:decimalPos] + "." + str[decimalPos:]
}

// NewMoney creates a new Money instance
func NewMoney(amount Decimal, currency string) *Money {
	return &Money{
		amount:   amount,
		currency: currency,
	}
}

// Multiply converts the currency and applies the conversion factor
func (m *Money) Multiply(factor Decimal) *Money {
	newAmount := m.rounded(m.amount.Mul(factor))
	newCurrency := USD
	if m.currency == USD {
		newCurrency = EUR
	}
	return NewMoney(newAmount, newCurrency)
}

// MultiplySecure performs a safe multiplication after ensuring factor is a new decimal
func (m *Money) MultiplySecure(factor Decimal) *Money {
	// Create a new decimal from the string representation to ensure immutability
	safeFactor, err := NewDecimalFromString(factor.String())
	if err != nil {
		panic("invalid decimal factor")
	}
	return m.Multiply(safeFactor)
}

// rounded rounds to 2 decimal places
func (m *Money) rounded(amount Decimal) Decimal {
	return amount.Round(decimals)
}

// GetAmount returns the monetary amount
func (m *Money) GetAmount() Decimal {
	return m.amount
}

// GetCurrency returns the currency code
func (m *Money) GetCurrency() string {
	return m.currency
}
