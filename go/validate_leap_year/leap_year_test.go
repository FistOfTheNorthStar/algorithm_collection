package dateutils

import "testing"

func TestIsLeapYear(t *testing.T) {
	tests := []struct {
		name     string
		year     int
		expected bool
	}{
		// Leap years
		{name: "divisible by 400", year: 400, expected: true},
		{name: "year 2000", year: 2000, expected: true},
		{name: "year 2020", year: 2020, expected: true},
		{name: "year 2024", year: 2024, expected: true},

		// Non-leap years
		{name: "year 401", year: 401, expected: false},
		{name: "year 2018", year: 2018, expected: false},
		{name: "year 2100", year: 2100, expected: false}, // divisible by 100 but not 400
		{name: "year 2023", year: 2023, expected: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsLeapYear(tt.year)
			if got != tt.expected {
				t.Errorf("IsLeapYear(%d) = %v, want %v", tt.year, got, tt.expected)
			}
		})
	}
}
