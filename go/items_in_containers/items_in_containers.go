package algorithms

import "errors"

// NumberOfItems counts items between pipes in given ranges of a string
func NumberOfItemsInContainer(s string, startIndices, endIndices []int) ([]int, error) {
	// Validate input sizes
	if len(startIndices) < 1 || len(startIndices) > 100000 {
		return nil, errors.New("wrong size in startIndices")
	}
	if len(endIndices) < 1 || len(endIndices) > 100000 {
		return nil, errors.New("wrong size in endIndices")
	}
	if len(startIndices) != len(endIndices) {
		return nil, errors.New("indices lengths must match")
	}

	numberOfItemsInClosedCompartments := make([]int, 0, len(startIndices))

	for idx := 0; idx < len(startIndices); idx++ {
		start := startIndices[idx]
		end := endIndices[idx]

		// Validate index values
		if start < 1 || start > 100000 {
			return nil, errors.New("wrong value at startIndices")
		}
		if end < 1 || end > 100000 {
			return nil, errors.New("wrong value at endIndices")
		}

		// Adjust for 0-based indexing
		substring := s[start-1 : end]

		numOfAsterisk := 0
		wasFirstPipeFound := false
		numOfAsteriskAccumulated := 0

		for _, c := range substring {
			switch c {
			case '|':
				if wasFirstPipeFound {
					numOfAsteriskAccumulated += numOfAsterisk
					numOfAsterisk = 0
				} else {
					wasFirstPipeFound = true
					numOfAsterisk = 0
				}
			case '*':
				numOfAsterisk++
			default:
				return nil, errors.New("wrong character")
			}
		}

		numberOfItemsInClosedCompartments = append(numberOfItemsInClosedCompartments, numOfAsteriskAccumulated)
	}

	return numberOfItemsInClosedCompartments, nil
}
