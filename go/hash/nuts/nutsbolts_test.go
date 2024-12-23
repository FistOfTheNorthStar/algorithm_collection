package hashtable

import (
	"reflect"
	"testing"
)

func TestMatch(t *testing.T) {
	tests := []struct {
		name     string
		nuts     []rune
		bolts    []rune
		expected []rune
	}{
		{
			name:     "basic matching",
			nuts:     []rune{'$', '%', '&', 'x', '@'},
			bolts:    []rune{'%', '@', 'x', '$', '&'},
			expected: []rune{'%', '@', 'x', '$', '&'},
		},
		// Additional test cases
		{
			name:     "empty arrays",
			nuts:     []rune{},
			bolts:    []rune{},
			expected: []rune{},
		},
		{
			name:     "single element",
			nuts:     []rune{'@'},
			bolts:    []rune{'@'},
			expected: []rune{'@'},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Make copies of input arrays
			nuts := make([]rune, len(tt.nuts))
			bolts := make([]rune, len(tt.bolts))
			copy(nuts, tt.nuts)
			copy(bolts, tt.bolts)

			// Run the matching algorithm
			Match(nuts, bolts)

			// Verify nuts array matches the expected result
			if !reflect.DeepEqual(nuts, tt.expected) {
				t.Errorf("Match(%v, %v) resulted in nuts = %v; want %v",
					string(tt.nuts), string(tt.bolts), string(nuts), string(tt.expected))
			}

			// Verify bolts array wasn't modified
			if !reflect.DeepEqual(bolts, tt.bolts) {
				t.Errorf("Bolts array was modified: got %v; want %v",
					string(bolts), string(tt.bolts))
			}
		})
	}
}
