package are_unique_chars

func AreUniqueChars(str string) bool {
	if len(str) > 128 {
		return false
	}

	chars := make([]bool, 128)
	for _, char := range str {
		if chars[char] {
			return false
		}
		chars[char] = true
	}
	return true
}
