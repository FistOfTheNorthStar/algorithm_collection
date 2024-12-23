package algorithms

import "testing"

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		target   int
		expected int
	}{
		{"Found", []int{1, 2, 3, 4, 5}, 3, 2},
		{"Not Found", []int{1, 2, 3, 4, 5}, 6, -1},
		{"Empty Array", []int{}, 1, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinarySearch(tt.arr, tt.target); got != tt.expected {
				t.Errorf("BinarySearch() = %v, want %v", got, tt.expected)
			}
		})
	}
}
