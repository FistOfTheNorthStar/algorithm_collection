package hashtable

// Match rearranges nuts array to match the order of bolts array
func Match(nuts, bolts []rune) {
	// Create a map to store nut positions
	nutMap := make(map[rune]int)
	for i, nut := range nuts {
		nutMap[nut] = i
	}

	// Rearrange nuts to match bolts
	for i, bolt := range bolts {
		if _, exists := nutMap[bolt]; exists {
			nuts[i] = bolt
		}
	}
}
