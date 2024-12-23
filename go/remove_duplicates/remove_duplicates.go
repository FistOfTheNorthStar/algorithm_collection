package listutils

import (
	"sort"
)

func RemoveDuplicatesAndOrder[T comparable](list []T) []T {
	seen := make(map[T]struct{})
	result := make([]T, 0)

	for _, v := range list {
		if _, ok := seen[v]; !ok {
			seen[v] = struct{}{}
			result = append(result, v)
		}
	}

	sort.Slice(result, func(i, j int) bool {
		switch v := any(result[i]).(type) {
		case int:
			return v < result[j].(int)
		case string:
			return v < result[j].(string)
		default:
			return false
		}
	})

	return result
}
