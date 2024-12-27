package money

import (
	"testing"
)

func TestConvertEuroToDollar(t *testing.T) {
	// Setup
	amount := NewDecimal(67.89)
	factor := NewDecimal(1.454706142288997)
	expected := NewDecimal(98.76)

	// Execute
	moneyInEuros := NewMoney(amount, "EUR")
	moneyInDollar := moneyInEuros.Multiply(factor)

	// Verify
	if moneyInDollar.GetAmount().String() != expected.String() {
		t.Errorf("Expected %v but got %v",
			expected.String(),
			moneyInDollar.GetAmount().String())
	}
	if moneyInDollar.GetCurrency() != USD {
		t.Errorf("Expected USD currency but got %v",
			moneyInDollar.GetCurrency())
	}
}

func TestConvertDollarToEuro(t *testing.T) {
	// Setup
	amount := NewDecimal(98.76)
	factor := NewDecimal(0.6874240583232078)
	expected := NewDecimal(67.89)

	// Execute
	moneyInDollar := NewMoney(amount, "USD")
	moneyInEuros := moneyInDollar.MultiplySecure(factor)

	// Verify
	if moneyInEuros.GetAmount().String() != expected.String() {
		t.Errorf("Expected %v but got %v",
			expected.String(),
			moneyInEuros.GetAmount().String())
	}
	if moneyInEuros.GetCurrency() != EUR {
		t.Errorf("Expected EUR currency but got %v",
			moneyInEuros.GetCurrency())
	}
}
