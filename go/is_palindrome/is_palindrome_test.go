// palindrome_test.go
package algorithms

import "testing"

func TestIsPalindrome(t *testing.T) {
	// Test non-palindrome cases
	t.Run("Not palindromes", func(t *testing.T) {
		tests := []struct {
			name  string
			input string
		}{
			{"Three chars", "2f1"},
			{"With hyphen", "-101"},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if IsPalindrome(tt.input) {
					t.Errorf("IsPalindrome(%s) = true, want false", tt.input)
				}
			})
		}
	})

	// Test palindrome cases
	t.Run("Valid palindromes", func(t *testing.T) {
		tests := []struct {
			name  string
			input string
		}{
			{"Five chars", "2f1f2"},
			{"With hyphens", "-101-"},
			{"Single digit", "9"},
			{"Two digits", "99"},
			{"Word", "madam"},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if !IsPalindrome(tt.input) {
					t.Errorf("IsPalindrome(%s) = false, want true", tt.input)
				}
			})
		}
	})
}
