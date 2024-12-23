package are_unique_chars

import "testing"

func TestAreUniqueChars(t *testing.T) {
	t.Run("Non-unique characters", func(t *testing.T) {
		tests := []struct {
			input string
		}{
			{"29s2"},
			{"1903aio9p"},
		}

		for _, tt := range tests {
			if AreUniqueChars(tt.input) {
				t.Errorf("AreUniqueChars(%s) = true, want false", tt.input)
			}
		}
	})

	t.Run("Unique characters", func(t *testing.T) {
		tests := []struct {
			input string
		}{
			{"29s13"},
			{"2813450769"},
		}

		for _, tt := range tests {
			if !AreUniqueChars(tt.input) {
				t.Errorf("AreUniqueChars(%s) = false, want true", tt.input)
			}
		}
	})
}
