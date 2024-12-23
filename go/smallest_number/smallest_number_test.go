package numberutils

import (
	"strconv"
	"testing"
)

func TestSmallest(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
		shouldBe bool // true for correct values, false for wrong values
	}{
		// Correct values
		{"four digits positive", 4751, 1000, true},
		{"three digits positive", 189, 100, true},
		{"two digits positive", 37, 10, true},
		{"single digit positive", 1, 0, true},
		{"zero", 0, 0, true},
		{"negative one", -1, -9, true},
		{"negative two digits", -38, -99, true},

		// Wrong values
		{"single digit wrong", 8, 1, false},
		{"four digits wrong", 2891, 2000, false},

		// Additional edge cases
		{"large number", 999999, 100000, true},
		{"negative large", -9999, -9999, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Smallest(tt.input)
			if tt.shouldBe {
				if result != tt.expected {
					t.Errorf("Smallest(%d) = %d, want %d",
						tt.input, result, tt.expected)
				}
			} else {
				if result == tt.expected {
					t.Errorf("Smallest(%d) = %d, should not equal %d",
						tt.input, result, tt.expected)
				}
			}
		})
	}
}

// Benchmark to measure performance
func BenchmarkSmallest(b *testing.B) {
	testCases := []int{4751, -38, 0, 999999}

	for _, tc := range testCases {
		b.Run(strconv.Itoa(tc), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Smallest(tc)
			}
		})
	}
}
