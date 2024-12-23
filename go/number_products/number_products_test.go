package number_products

import "testing"

func TestValidProducts(t *testing.T) {
	tests := []struct {
		name     string
		x, y     int
		expected int
		wantThis bool // true if we expect this exact value, false if we expect any other value
	}{
		{"case 6-20", 6, 20, 3, true},
		{"case 1000-1130", 1000, 1130, 2, true},
		{"case 21-29", 21, 29, 1, false}, // expect anything but 1
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ValidProducts(tt.x, tt.y)
			if tt.wantThis && got != tt.expected {
				t.Errorf("ValidProducts(%d, %d) = %d; want %d",
					tt.x, tt.y, got, tt.expected)
			} else if !tt.wantThis && got == tt.expected {
				t.Errorf("ValidProducts(%d, %d) = %d; want any other value",
					tt.x, tt.y, got)
			}
		})
	}
}

func TestRightProducts(t *testing.T) {
	// Test first case
	if got := ValidProducts(6, 20); got != 3 {
		t.Errorf("ValidProducts(6, 20) = %d; want 3", got)
	}

	// Test second case
	if got := ValidProducts(1000, 1130); got != 2 {
		t.Errorf("ValidProducts(1000, 1130) = %d; want 2", got)
	}
}

func TestWrongProducts(t *testing.T) {
	if got := ValidProducts(21, 29); got == 1 {
		t.Errorf("ValidProducts(21, 29) = 1; expected different value")
	}
}
