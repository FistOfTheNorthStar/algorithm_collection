package most_frequent

import "sort"

// most frequent char in the array
func MostFrequent(array []int) []int {
	frqMap := make(map[int]int)
	for _, elem := range array {
		frqMap[elem]++
	}

	maxFreq := 0
	for _, frq := range frqMap {
		if frq > maxFreq {
			maxFreq = frq
		}
	}

	var result []int
	for elem, freq := range frqMap {
		if freq == maxFreq {
			result = append(result, elem)
		}
	}

	sort.Ints(result)

	return result
}
