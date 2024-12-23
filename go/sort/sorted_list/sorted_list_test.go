package sort

import (
	"reflect"
	"testing"
)

func TestMergeSorted(t *testing.T) {
	tests := []struct {
		name     string
		list1    []int
		list2    []int
		expected []int
	}{
		{
			name:     "merge with duplicates",
			list1:    []int{1, 1, 2, 5, 8},
			list2:    []int{3, 4, 6},
			expected: []int{1, 1, 2, 3, 4, 5, 6, 8},
		},
		{
			name:     "merge without duplicates",
			list1:    []int{2, 4, 5},
			list2:    []int{1, 3, 6},
			expected: []int{1, 2, 3, 4, 5, 6},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Make copies of input slices to verify they're not modified
			list1Copy := make([]int, len(tt.list1))
			list2Copy := make([]int, len(tt.list2))
			copy(list1Copy, tt.list1)
			copy(list2Copy, tt.list2)

			result := MergeSorted(tt.list1, tt.list2)

			// Check if result is correct
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("MergeSorted(%v, %v) = %v; want %v",
					list1Copy, list2Copy, result, tt.expected)
			}

			// Verify input slices weren't modified
			if !reflect.DeepEqual(tt.list1, list1Copy) {
				t.Errorf("First input slice was modified: got %v; want %v",
					tt.list1, list1Copy)
			}
			if !reflect.DeepEqual(tt.list2, list2Copy) {
				t.Errorf("Second input slice was modified: got %v; want %v",
					tt.list2, list2Copy)
			}
		})
	}
}

func TestMergeSortedLists2(t *testing.T) {
	// Test inputs
	list1 := []int{2, 4, 5}
	list2 := []int{1, 3, 6}
	expected := []int{1, 2, 3, 4, 5, 6}

	// Get merged result
	result := MergeSorted(list1, list2)

	// Compare with expected output
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("MergeSorted(%v, %v) = %v; want %v",
			list1, list2, result, expected)
	}
}
