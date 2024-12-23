package stack

import "testing"

func TestIncorrectExpressions(t *testing.T) {
	tests := []struct {
		name       string
		expression string
	}{
		{"nil expression", ""},
		{"empty expression", ""},
		{"single bracket", "("},
		{"mismatched brackets", "(()a]"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if IsBalanced(tt.expression) {
				t.Errorf("IsBalanced(%q) = true; want false", tt.expression)
			}
		})
	}
}

func TestCorrectExpressions(t *testing.T) {
	tests := []struct {
		name       string
		expression string
	}{
		{"simple parentheses", "()"},
		{"nested brackets", "([])"},
		{"complex nesting", "{{([])}}"},
		{"with letters", "{{a([b])}c}dd"},
		{"mathematical expression", "(w*(x+y)/z−(p/(r−q)))"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !IsBalanced(tt.expression) {
				t.Errorf("IsBalanced(%q) = false; want true", tt.expression)
			}
		})
	}
}
