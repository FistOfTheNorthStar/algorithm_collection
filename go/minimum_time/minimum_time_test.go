package assembler

import "testing"

func TestMinimumTime(t *testing.T) {
	tests := []struct {
		name       string
		numOfParts int
		list       []int
		expected   int
	}{
		{
			name:       "four parts",
			numOfParts: 4,
			list:       []int{8, 4, 6, 12},
			expected:   58,
		},
		{
			name:       "five parts",
			numOfParts: 5,
			list:       []int{3, 7, 2, 10, 5},
			expected:   59,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MinimumTime(tt.numOfParts, tt.list)
			if result != tt.expected {
				t.Errorf("MinimumTime(%d, %v) = %d; want %d",
					tt.numOfParts, tt.list, result, tt.expected)
			}
		})
	}
}
