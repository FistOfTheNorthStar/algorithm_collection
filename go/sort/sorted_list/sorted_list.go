package sort

// MergeSorted merges two sorted slices into a single sorted slice
func MergeSorted(list1, list2 []int) []int {
	mergedList := make([]int, 0, len(list1)+len(list2))
	idx1 := 0
	idx2 := 0

	// Merge while both lists have elements
	for idx1 < len(list1) && idx2 < len(list2) {
		if list1[idx1] <= list2[idx2] {
			mergedList = append(mergedList, list1[idx1])
			idx1++
		} else {
			mergedList = append(mergedList, list2[idx2])
			idx2++
		}
	}

	// Append remaining elements from list1, if any
	if idx1 < len(list1) {
		mergedList = append(mergedList, list1[idx1:]...)
	}

	// Append remaining elements from list2, if any
	if idx2 < len(list2) {
		mergedList = append(mergedList, list2[idx2:]...)
	}

	return mergedList
}
