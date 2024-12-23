package stack

// Reverse returns the reversed version of the input string
func Reverse(text string) string {
	// Convert string to rune slice to handle Unicode characters correctly
	runes := []rune(text)

	// Create a stack using slice
	stack := make([]rune, 0, len(runes))

	// Push all characters onto stack
	for _, r := range runes {
		stack = append(stack, r)
	}

	// Pop characters from stack to create reversed string
	for i := 0; i < len(runes); i++ {
		// Pop from stack (take last element)
		lastIdx := len(stack) - 1
		runes[i] = stack[lastIdx]
		stack = stack[:lastIdx]
	}

	return string(runes)
}
