package number_products

import (
	"fmt"
	"math"
)

// ValidProducts counts how many products of consecutive integers N*(N+1)
// exist between X and Y inclusive
func ValidProducts(x, y int) int {
	if x < 1 {
		panic(fmt.Sprintf("invalid input: %d", x))
	}

	n := int(math.Sqrt(float64(x)))
	validCount := 0

	for n*(n+1) <= y {
		product := n * (n + 1)
		if product >= x && product <= y {
			validCount++
		}
		n++
	}

	return validCount
}
