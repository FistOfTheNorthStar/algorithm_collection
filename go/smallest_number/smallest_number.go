package numberutils

import (
	"math"
	"strconv"
)

// Smallest finds the smallest number with same length as N, or one digit less for negative numbers
func Smallest(N int) int {
	// Get number of digits, handling negative numbers
	numberOfDigits := len(strconv.Itoa(int(math.Abs(float64(N)))))

	if N >= 0 {
		if numberOfDigits == 1 {
			return 0
		}
		// For positive numbers: 10^(number of digits - 1)
		return int(math.Pow10(numberOfDigits - 1))
	}

	// For negative numbers: -(10^number of digits - 1)
	return 1 - int(math.Pow10(numberOfDigits))
}
