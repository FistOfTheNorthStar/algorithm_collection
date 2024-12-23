package search

import "testing"

func TestBinarySearchNotFound(t *testing.T) {
	tests := []struct {
		name   string
		target string
		array  []string
	}{
		{
			name:   "target between elements",
			target: "fin",
			array:  []string{"ada", "fda"},
		},
		{
			name:   "target not in longer array",
			target: "eda",
			array:  []string{"ada", "bda", "cda", "dda"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Search(tt.target, tt.array); got {
				t.Errorf("Search(%v, %v) = true; want false",
					tt.target, tt.array)
			}
		})
	}
}

func TestBinarySearchFound(t *testing.T) {
	// String tests
	t.Run("string search", func(t *testing.T) {
		array := []string{"ada", "cal", "fda"}
		if !Search("cal", array) {
			t.Error("Search(\"cal\", [\"ada\", \"cal\", \"fda\"]) = false; want true")
		}
	})

	// Integer tests
	intTests := []struct {
		name   string
		target int
		array  []int
	}{
		{
			name:   "small sorted array",
			target: 21,
			array:  []int{1, 2, 3, 4, 5, 21},
		},
		{
			name:   "larger sorted array",
			target: 21,
			array:  []int{3, 7, 9, 13, 18, 21, 41, 52, 81, 97},
		},
	}

	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			if !Search(tt.target, tt.array) {
				t.Errorf("Search(%v, %v) = false; want true",
					tt.target, tt.array)
			}
		})
	}
}
