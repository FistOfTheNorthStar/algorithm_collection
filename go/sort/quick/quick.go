package sort

// QuickSort sorts an array using the quicksort algorithm
func QuickSort(numbers []int, lo, hi int) []int {
	// Make a copy of the input slice to avoid modifying the original
	result := make([]int, len(numbers))
	copy(result, numbers)

	quickSortHelper(result, lo, hi)
	return result
}

// quickSortHelper performs the recursive quicksort
func quickSortHelper(numbers []int, lo, hi int) {
	if lo < hi {
		partitionBorder := partition(numbers, lo, hi)
		// Sort elements recursively
		quickSortHelper(numbers, lo, partitionBorder-1)
		quickSortHelper(numbers, partitionBorder+1, hi)
	}
}

// partition helps to place the pivot element at its correct position
func partition(numbers []int, lo, hi int) int {
	// Element to be placed at right position
	pivot := numbers[hi]
	// Index of smaller element
	i := lo - 1

	// Traverse through all elements
	for j := lo; j < hi; j++ {
		// If current element is smaller than pivot
		if numbers[j] < pivot {
			i++
			// Swap elements
			numbers[i], numbers[j] = numbers[j], numbers[i]
		}
	}

	// Place pivot in its final position
	numbers[hi], numbers[i+1] = numbers[i+1], pivot
	return i + 1
}
