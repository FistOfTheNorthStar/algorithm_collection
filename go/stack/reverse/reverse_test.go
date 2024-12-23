package stack

import "testing"

func TestReverse(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "simple word",
			input:    "cafe",
			expected: "efac",
		},
		{
			name:     "alphanumeric string",
			input:    "abc2132",
			expected: "2312cba",
		},
		// Additional test cases for robustness
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "unicode string",
			input:    "hello世界",
			expected: "界世olleh",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Reverse(tt.input)
			if result != tt.expected {
				t.Errorf("Reverse(%q) = %q; want %q",
					tt.input, result, tt.expected)
			}
		})
	}
}
