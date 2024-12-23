package search

type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64 | ~string
}

// Search performs binary search on a sorted slice of ordered elements
func Search[T Ordered](target T, array []T) bool {
	if len(array) == 0 {
		return false
	}

	min := 0
	max := len(array) - 1

	for min <= max {
		mid := (min + max) / 2
		if target < array[mid] {
			max = mid - 1
		} else if target > array[mid] {
			min = mid + 1
		} else {
			return true
		}
	}
	return false
}
