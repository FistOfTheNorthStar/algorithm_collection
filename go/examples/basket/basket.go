package basket

import (
	"fmt"
	"sort"
)

type BasketOptimizer struct{}

func NewBasketOptimizer() *BasketOptimizer {
	return &BasketOptimizer{}
}

// Fill finds the most valuable combination of products within budget
// Each product is represented as [id, price, value]
func (b *BasketOptimizer) Fill(products [][]int, budget int) [][]int {
	length := len(products)
	numIterations := 1 << length

	// Map to store combinations and their total values
	combinations := make(map[string]struct {
		products [][]int
		value    int
	})

	// Try all possible combinations
	for idx := 0; idx < numIterations; idx++ {
		subset := make([][]int, 0, length)
		sumPrice := 0
		sumValue := 0

		// Check each bit position
		for idx2 := 0; idx2 < length; idx2++ {
			if (idx & (1 << idx2)) == 0 {
				subset = append(subset, products[idx2])
				sumPrice += products[idx2][1]
				sumValue += products[idx2][2]
			}
		}

		// If combination is within budget, add to map
		if len(subset) > 0 && sumPrice <= budget {
			key := generateKey(subset)
			combinations[key] = struct {
				products [][]int
				value    int
			}{
				products: subset,
				value:    sumValue,
			}
		}
	}

	// Find combination with highest value
	var maxValue int
	var bestProducts [][]int
	for _, combo := range combinations {
		if combo.value > maxValue {
			maxValue = combo.value
			bestProducts = combo.products
		}
	}
	return bestProducts
}

// generateKey creates a unique string key for a product combination
func generateKey(products [][]int) string {
	// Sort products by ID to ensure consistent key generation
	sorted := make([][]int, len(products))
	copy(sorted, products)
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i][0] < sorted[j][0]
	})

	// Create key string
	result := ""
	for _, p := range sorted {
		result += fmt.Sprintf("%d-", p[0])
	}
	return result
}
