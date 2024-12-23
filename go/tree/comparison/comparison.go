package tree

import (
	"fmt"
	"math"
)

// Comparisons calculates the number of comparisons needed for N elements
func Comparisons(n int) int {
	accumElements := 0
	comparisons := 0

	for level := 0; level <= n/2; level++ {
		power := int(math.Pow(2, float64(level)))
		accumElements += power

		if accumElements >= n {
			comparisons = level + 1
			break
		}
	}

	fmt.Printf("comparisons -> %d\n", comparisons)
	return comparisons
}
