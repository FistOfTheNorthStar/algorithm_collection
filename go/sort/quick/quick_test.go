package sort

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "small array",
			input:    []int{13, 12, 14, 6, 7},
			expected: []int{6, 7, 12, 13, 14},
		},
		{
			name:     "array with negative and duplicates",
			input:    []int{7, 9, 1, 4, 9, 12, 4, 13, -2, 9},
			expected: []int{-2, 1, 4, 4, 7, 9, 9, 9, 12, 13},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Make a copy of input to verify it's not modified
			original := make([]int, len(tt.input))
			copy(original, tt.input)

			result := QuickSort(tt.input, 0, len(tt.input)-1)

			// Check if result is correct
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("QuickSort(%v) = %v; want %v",
					original, result, tt.expected)
			}

			// Verify original slice wasn't modified
			if !reflect.DeepEqual(tt.input, original) {
				t.Errorf("Original slice was modified: got %v; want %v",
					tt.input, original)
			}
		})
	}
}
