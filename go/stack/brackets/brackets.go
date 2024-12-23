package stack

// IsBalanced checks if the brackets in an expression are balanced
func IsBalanced(expression string) bool {
	if expression == "" {
		return false
	}

	// Create a stack using a slice
	stack := make([]rune, 0, len(expression))

	for _, char := range expression {
		switch char {
		case '{', '(', '[':
			stack = append(stack, char)

		case '}':
			if len(stack) == 0 || stack[len(stack)-1] != '{' {
				return false
			}
			stack = stack[:len(stack)-1]

		case ')':
			if len(stack) == 0 || stack[len(stack)-1] != '(' {
				return false
			}
			stack = stack[:len(stack)-1]

		case ']':
			if len(stack) == 0 || stack[len(stack)-1] != '[' {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
