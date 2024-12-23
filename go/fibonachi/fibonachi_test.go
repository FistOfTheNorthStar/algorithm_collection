package fibonacci

import "testing"

func TestFibonacci(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{"fibonacci of 3", 3, 2},
		{"fibonacci of 7", 7, 13},
		{"fibonacci of 10", 10, 55},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Fibonacci(tt.input)
			if result != tt.expected {
				t.Errorf("Fibonacci(%d) = %d; want %d",
					tt.input, result, tt.expected)
			}
		})
	}
}
