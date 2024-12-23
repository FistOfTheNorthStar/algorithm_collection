package money

import (
	"testing"

	"github.com/shopspring/decimal"
)

func TestConvertEuroToDollar(t *testing.T) {
	// Setup
	amount, _ := decimal.NewFromString("67.89")
	factor, _ := decimal.NewFromString("1.454706142288997")
	expected, _ := decimal.NewFromString("98.76")

	// Execute
	moneyInEuros := NewMoney(amount, "EUR")
	moneyInDollar := moneyInEuros.Multiply(factor)

	// Verify
	if !moneyInDollar.GetAmount().Equal(expected) {
		t.Errorf("Expected %v but got %v", expected, moneyInDollar.GetAmount())
	}
	if moneyInDollar.GetCurrency() != USD {
		t.Errorf("Expected USD currency but got %v", moneyInDollar.GetCurrency())
	}
}

func TestConvertDollarToEuro(t *testing.T) {
	// Setup
	amount, _ := decimal.NewFromString("98.76")
	factor, _ := decimal.NewFromString("0.6874240583232078")
	expected, _ := decimal.NewFromString("67.89")

	// Execute
	moneyInDollar := NewMoney(amount, "USD")
	moneyInEuros := moneyInDollar.MultiplySecure(factor)

	// Verify
	if !moneyInEuros.GetAmount().Equal(expected) {
		t.Errorf("Expected %v but got %v", expected, moneyInEuros.GetAmount())
	}
	if moneyInEuros.GetCurrency() != EUR {
		t.Errorf("Expected EUR currency but got %v", moneyInEuros.GetCurrency())
	}
}
