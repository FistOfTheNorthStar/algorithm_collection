package sort

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "small array",
			input:    []int{6, 4, 9, 5},
			expected: []int{4, 5, 6, 9},
		},
		{
			name:     "larger array with duplicates",
			input:    []int{7, 9, 1, 4, 9, 12, 4, 13, 9},
			expected: []int{1, 4, 4, 7, 9, 9, 9, 12, 13},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Make a copy of input to verify it's not modified
			original := make([]int, len(tt.input))
			copy(original, tt.input)

			result := BubbleSort(tt.input)

			// Check if result is correct
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("BubbleSort(%v) = %v; want %v",
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

func TestBubbleSortNilSlice(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic for nil slice, but it didn't happen")
		}
	}()

	BubbleSort(nil)
}
