package assembler

import "sort"

// MinimumTime calculates the minimum time required to assemble all parts
func MinimumTime(numOfParts int, list []int) int {
	// Create a copy of the slice to avoid modifying the input
	sizes := make([]int, len(list))
	copy(sizes, list)

	// Sort the array
	sort.Ints(sizes)

	accumulatedTime := 0

	// Iterate through the array up to the second-to-last element
	for i := 0; i < len(sizes)-1; i++ {
		// Calculate the time to assemble current and next part
		currentTime := sizes[i] + sizes[i+1]
		accumulatedTime += currentTime

		// Update the next element with the accumulated time
		// This represents the combined part for the next iteration
		sizes[i+1] = currentTime
	}

	return accumulatedTime
}
