package factorial

import "testing"

func TestFactorialRightResults(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{"factorial of 1", 1, 1},
		{"factorial of 0", 0, 1},
		{"factorial of 3", 3, 6},
		{"factorial of 6", 6, 720},
	}

	for _, tt := range tests {
		t.Run(tt.name+" recursive", func(t *testing.T) {
			if got := FactorialRecursive(tt.input); got != tt.expected {
				t.Errorf("FactorialRecursive(%d) = %d; want %d",
					tt.input, got, tt.expected)
			}
		})

		t.Run(tt.name+" iterative", func(t *testing.T) {
			if got := FactorialIterative(tt.input); got != tt.expected {
				t.Errorf("FactorialIterative(%d) = %d; want %d",
					tt.input, got, tt.expected)
			}
		})
	}
}

func TestFactorialWrongResults(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		notEqual int
	}{
		{"factorial of 3 not 5", 3, 5},
		{"factorial of 4 not 10", 4, 10},
	}

	for _, tt := range tests {
		t.Run(tt.name+" recursive", func(t *testing.T) {
			if got := FactorialRecursive(tt.input); got == tt.notEqual {
				t.Errorf("FactorialRecursive(%d) = %d; want different value",
					tt.input, got)
			}
		})

		t.Run(tt.name+" iterative", func(t *testing.T) {
			if got := FactorialIterative(tt.input); got == tt.notEqual {
				t.Errorf("FactorialIterative(%d) = %d; want different value",
					tt.input, got)
			}
		})
	}
}
