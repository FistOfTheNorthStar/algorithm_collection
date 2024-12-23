package most_frequent

import (
	"reflect"
	"testing"
)

func TestMostFrequent(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "two most frequent",
			input:    []int{3, 2, 0, 3, 1, 2},
			expected: []int{2, 3},
		},
		{
			name:     "single most frequent",
			input:    []int{3, 5, 0, 5, 5, 1, 2},
			expected: []int{5},
		},
		{
			name:     "single element array",
			input:    []int{7},
			expected: []int{7},
		},
		{
			name:     "all same frequency",
			input:    []int{1, 2, 3},
			expected: []int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Make a copy of input to verify it's not modified
			input := make([]int, len(tt.input))
			copy(input, tt.input)

			result := MostFrequent(tt.input)

			// Check if result is correct
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("MostFrequent(%v) = %v; want %v",
					input, result, tt.expected)
			}

			// Verify input wasn't modified
			if !reflect.DeepEqual(tt.input, input) {
				t.Errorf("Input was modified: got %v; want %v",
					tt.input, input)
			}
		})
	}
}
