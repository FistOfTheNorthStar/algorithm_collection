package basket

import (
	"testing"
)

func sumValues(products [][]int) int {
	sum := 0
	for _, p := range products {
		sum += p[2]
	}
	return sum
}

func TestBasketOptimizer(t *testing.T) {
	tests := []struct {
		name     string
		products [][]int
		budget   int
		expected int
	}{
		{
			name: "products ordered by value",
			products: [][]int{
				{1, 98, 230}, // Multiplied prices by 100 to convert to cents
				{2, 98, 230},
				{3, 48, 75},
				{4, 129, 55},
				{5, 129, 47},
				{6, 486, 14},
				{7, 169, 12},
			},
			budget:   400, // 4.00 converted to cents
			expected: 590,
		},
		{
			name: "products not ordered by value",
			products: [][]int{
				{1, 98, 230},
				{2, 129, 47},
				{3, 169, 12},
				{4, 129, 55},
				{5, 98, 230},
				{6, 486, 14},
				{7, 48, 75},
			},
			budget:   400,
			expected: 590,
		},
		{
			name: "different products",
			products: [][]int{
				{1, 98, 230},
				{2, 51, 30},
				{3, 49, 28},
				{4, 129, 55},
				{5, 98, 230},
				{6, 486, 14},
				{7, 48, 75},
			},
			budget:   400,
			expected: 593,
		},
	}

	optimizer := NewBasketOptimizer()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := optimizer.Fill(tt.products, tt.budget)
			totalValue := sumValues(result)
			if totalValue != tt.expected {
				t.Errorf("Fill() total value = %d; want %d",
					totalValue, tt.expected)
			}
		})
	}
}
