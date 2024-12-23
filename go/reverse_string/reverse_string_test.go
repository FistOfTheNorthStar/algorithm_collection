package algorithms

import "testing"

func TestReverse(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Empty string", "", ""},
		{"Single char", "a", "a"},
		{"Simple word", "hello", "olleh"},
		{"With spaces", "hello world", "dlrow olleh"},
		{"Unicode", "hello世界", "界世olleh"},
		{"Test case 1", "2312cba", "abc2132"},
		{"Test case 2", "ab", "ba"},
		{"Test case 3", "1a c", "c a1"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := Reverse(tt.input); got != tt.expected {
				t.Errorf("Reverse() = %v, want %v", got, tt.expected)
			}
		})
	}

	t.Run("Empty string case", func(t *testing.T) {
		_, err := Reverse("")
		if err == nil {
			t.Error("Expected error for empty string, got nil")
		}
	})
}
