package listutils

import "sort"

// Ordered is an interface for types that can be ordered
type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64 | ~string
}

func RemoveDuplicatesAndOrder[T Ordered](list []T) []T {
	seen := make(map[T]struct{})
	result := make([]T, 0)

	// Remove duplicates
	for _, v := range list {
		if _, ok := seen[v]; !ok {
			seen[v] = struct{}{}
			result = append(result, v)
		}
	}

	// Sort using the natural ordering of T
	sort.Slice(result, func(i, j int) bool {
		return result[i] < result[j]
	})

	return result
}
