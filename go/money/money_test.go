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
	gotFloat, _ := moneyInDollar.GetAmount().Float64()
	expectedFloat, _ := expected.Float64()

	if !floatEquals(gotFloat, expectedFloat, 0.01) {
		t.Errorf("Expected %v but got %v",
			expectedFloat,
			gotFloat)
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
	gotFloat, _ := moneyInEuros.GetAmount().Float64()
	expectedFloat, _ := expected.Float64()

	if !floatEquals(gotFloat, expectedFloat, 0.01) {
		t.Errorf("Expected %v but got %v",
			expectedFloat,
			gotFloat)
	}

	if moneyInEuros.GetCurrency() != EUR {
		t.Errorf("Expected EUR currency but got %v",
			moneyInEuros.GetCurrency())
	}
}

// Helper function to compare floats with tolerance
func floatEquals(a, b, tolerance float64) bool {
	diff := a - b
	if diff < 0 {
		diff = -diff
	}
	return diff <= tolerance
}
