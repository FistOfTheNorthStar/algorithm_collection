package algorithms

import (
	"strconv"
	"strings"
)

func CompareVersions(v1, v2 string) int {
	a1 := parseVersion(v1)
	a2 := parseVersion(v2)

	idx := 0
	for idx < len(a1) || idx < len(a2) {
		if idx < len(a1) && idx < len(a2) {
			if a1[idx] < a2[idx] {
				return -1
			} else if a1[idx] > a2[idx] {
				return 1
			}
		} else if idx < len(a1) {
			if a1[idx] != 0 {
				return 1
			}
		} else if idx < len(a2) {
			if a2[idx] != 0 {
				return -1
			}
		}
		idx++
	}
	return 0
}

func parseVersion(v string) []int {
	parts := strings.Split(v, ".")
	result := make([]int, len(parts))

	for i, part := range parts {
		num, _ := strconv.Atoi(strings.TrimSpace(part))
		result[i] = num
	}

	return result
}
