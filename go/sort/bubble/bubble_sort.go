package sort

// BubbleSort sorts an array using the bubble sort algorithm
func BubbleSort(numbers []int) []int {
	if numbers == nil {
		panic("array not initialized")
	}

	// Make a copy of the input slice to avoid modifying the original
	result := make([]int, len(numbers))
	copy(result, numbers)

	for {
		numbersSwapped := false

		for i := 0; i < len(result)-1; i++ {
			if result[i] > result[i+1] {
				// Swap elements
				result[i], result[i+1] = result[i+1], result[i]
				numbersSwapped = true
			}
		}

		if !numbersSwapped {
			break
		}
	}

	return result
}
