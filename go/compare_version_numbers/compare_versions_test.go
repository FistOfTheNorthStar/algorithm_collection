package algorithms

import "testing"

func TestCompareVersions(t *testing.T) {
	tests := []struct {
		name     string
		v1       string
		v2       string
		expected int
	}{
		{"Same version with zero", "15", "15.0", 0},
		{"Same version with trailing zero", "10.1", "10.1.0", 0},
		{"Lower version", "10.1", "10.1.1", -1},
		{"Higher version", "10.1.2", "10.1.1", 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CompareVersions(tt.v1, tt.v2)
			if result != tt.expected {
				t.Errorf("Compare(%s, %s) = %d, want %d",
					tt.v1, tt.v2, result, tt.expected)
			}
		})
	}
}
