package numberutils

import (
	"fmt"
	"testing"
)

func TestIsNumberAValidPowerOfBase(t *testing.T) {
	tests := []struct {
		name     string
		number   int
		base     int
		expected bool
	}{
		// Invalid cases
		{
			name:     "6 is not a power of 2",
			number:   6,
			base:     2,
			expected: false,
		},
		{
			name:     "16 is not a power of 5",
			number:   16,
			base:     5,
			expected: false,
		},
		{
			name:     "14 is not a power of 7",
			number:   14,
			base:     7,
			expected: false,
		},

		// Valid cases
		{
			name:     "243 is a power of 3",
			number:   243,
			base:     3,
			expected: true,
		},
		{
			name:     "16 is a power of 4",
			number:   16,
			base:     4,
			expected: true,
		},
		{
			name:     "125 is a power of 5",
			number:   125,
			base:     5,
			expected: true,
		},

		// Edge cases
		{
			name:     "0 is considered a power of any base",
			number:   0,
			base:     2,
			expected: true,
		},
		{
			name:     "1 is a power of any base",
			number:   1,
			base:     7,
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsNumberAValidPowerOfBase(tt.number, tt.base)
			if result != tt.expected {
				t.Errorf("IsNumberAValidPowerOfBase(%d, %d) = %v, want %v",
					tt.number, tt.base, result, tt.expected)
			}
		})
	}
}

// Benchmark to test performance
func BenchmarkIsNumberAValidPowerOfBase(b *testing.B) {
	testCases := []struct {
		number int
		base   int
	}{
		{243, 3},
		{16, 4},
		{6, 2},
		{125, 5},
	}

	for _, tc := range testCases {
		b.Run(fmt.Sprintf("num=%d,base=%d", tc.number, tc.base), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				IsNumberAValidPowerOfBase(tc.number, tc.base)
			}
		})
	}
}
