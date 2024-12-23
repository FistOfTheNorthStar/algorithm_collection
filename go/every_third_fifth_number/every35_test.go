package numberutils

import (
	"strings"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	tests := []struct {
		name     string
		N        int
		expected string
	}{
		{
			name:     "up to 15",
			N:        15,
			expected: "1, 2, Fizz, 4, Buzz, Fizz, 7, 8, Fizz, Buzz, 11, Fizz, 13, 14, Fizz-Buzz",
		},
		{
			name:     "up to 5",
			N:        5,
			expected: "1, 2, Fizz, 4, Buzz",
		},
		{
			name:     "single number",
			N:        1,
			expected: "1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FizzBuzz(tt.N)
			if result != tt.expected {
				t.Errorf("FizzBuzz(%d) =\n%v\nwant:\n%v",
					tt.N, result, tt.expected)
			}
		})
	}
}

func TestFizzBuzzProperties(t *testing.T) {
	result := FizzBuzz(100)
	elements := strings.Split(result, ", ")

	t.Run("check length", func(t *testing.T) {
		if len(elements) != 100 {
			t.Errorf("Expected 100 elements, got %d", len(elements))
		}
	})

	t.Run("check Fizz-Buzz positions", func(t *testing.T) {
		for i, element := range elements {
			num := i + 1
			switch {
			case num%15 == 0:
				if element != "Fizz-Buzz" {
					t.Errorf("Position %d should be Fizz-Buzz, got %s", num, element)
				}
			case num%3 == 0:
				if element != "Fizz" {
					t.Errorf("Position %d should be Fizz, got %s", num, element)
				}
			case num%5 == 0:
				if element != "Buzz" {
					t.Errorf("Position %d should be Buzz, got %s", num, element)
				}
			}
		}
	})
}

// Benchmark the FizzBuzz function
func BenchmarkFizzBuzz(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FizzBuzz(100)
	}
}
