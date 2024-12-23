package algorithms

import "testing"

func TestNotPrimeNumbers(t *testing.T) {
	tests := []struct {
		name   string
		input  int
		expect bool
	}{
		{"negative number", -1, false},
		{"perfect square", 625, false},
		{"small composite", 4, false},
		{"larger composite", 100, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPrimeNumber(tt.input); got != tt.expect {
				t.Errorf("IsPrimeNumber(%d) = %v, want %v", tt.input, got, tt.expect)
			}
		})
	}
}

func TestPrimeNumbers(t *testing.T) {
	tests := []struct {
		name   string
		input  int
		expect bool
	}{
		{"smallest prime", 2, true},
		{"small prime", 3, true},
		{"prime 5", 5, true},
		{"prime 7", 7, true},
		{"larger prime", 73, true},
	}

	for _, tt := range tests {
		t.Run(tt.Run, func(t *testing.T) {
			if got := IsPrimeNumber(tt.input); got != tt.expect {
				t.Errorf("IsPrimeNumber(%d) = %v, want %v", tt.input, got, tt.expect)
			}
		})
	}
}
