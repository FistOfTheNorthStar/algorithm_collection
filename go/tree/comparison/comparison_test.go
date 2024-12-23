package tree

import "testing"

func TestComparisons(t *testing.T) {
	tests := []struct {
		name     string
		elements int
		maxComp  int
	}{
		{
			name:     "small tree (15 elements)",
			elements: 15,
			maxComp:  4,
		},
		{
			name:     "medium tree (31 elements)",
			elements: 31,
			maxComp:  5,
		},
		{
			name:     "large tree (1000 elements)",
			elements: 1000,
			maxComp:  10,
		},
		{
			name:     "very large tree (1B elements)",
			elements: 1000000000,
			maxComp:  30,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Comparisons(tt.elements)
			if result > tt.maxComp {
				t.Errorf("Comparisons(%d) = %d; want <= %d",
					tt.elements, result, tt.maxComp)
			}
		})
	}
}

// Additional test cases
func TestEdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		elements int
	}{
		{"zero elements", 0},
		{"one element", 1},
		{"negative elements", -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Comparisons(tt.elements)
			if result < 0 {
				t.Errorf("Comparisons(%d) = %d; want >= 0",
					tt.elements, result)
			}
		})
	}
}
