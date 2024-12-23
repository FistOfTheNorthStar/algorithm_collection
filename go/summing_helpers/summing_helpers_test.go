package numberutils

import "testing"

func TestSumOfEvenNumbers(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{
			name:     "sum up to 12",
			input:    12,
			expected: 42,
		},
		{
			name:     "sum up to 21",
			input:    21,
			expected: 110,
		},
		{
			name:     "sum up to 0",
			input:    0,
			expected: 0,
		},
		{
			name:     "sum up to 2",
			input:    2,
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name+" (for loop)", func(t *testing.T) {
			if got := SumOfEvenNumbersWithFor(tt.input); got != tt.expected {
				t.Errorf("SumOfEvenNumbersWithFor(%d) = %v, want %v",
					tt.input, got, tt.expected)
			}
		})

		t.Run(tt.name+" (step by 2)", func(t *testing.T) {
			if got := SumOfEvenNumbersWithStep(tt.input); got != tt.expected {
				t.Errorf("SumOfEvenNumbersWithStep(%d) = %v, want %v",
					tt.input, got, tt.expected)
			}
		})
	}
}

// Benchmark both implementations
func BenchmarkSumOfEvenNumbers(b *testing.B) {
	b.Run("for loop", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			SumOfEvenNumbersWithFor(1000)
		}
	})

	b.Run("step by 2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			SumOfEvenNumbersWithStep(1000)
		}
	})
}
