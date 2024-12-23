package hashtable

import "testing"

func TestHashTable(t *testing.T) {
	t.Run("basic operations", func(t *testing.T) {
		ht := NewHashTable[string, string]()

		// Test putting and getting values
		testCases := []struct {
			key   string
			value string
		}{
			{"DAB", "Atlanta"},
			{"ABC", "Chicago"},
			{"CAD", "Dallas"},
			{"BAD", "Detroit"},
		}

		// Insert all test cases
		for _, tc := range testCases {
			ht.Put(tc.key, tc.value)
		}

		// Verify all values can be retrieved
		for _, tc := range testCases {
			got, exists := ht.Get(tc.key)
			if !exists {
				t.Errorf("Get(%q) returned not exists, want exists", tc.key)
				continue
			}
			if got != tc.value {
				t.Errorf("Get(%q) = %q, want %q", tc.key, got, tc.value)
			}
		}
	})

	t.Run("nonexistent key", func(t *testing.T) {
		ht := NewHashTable[string, string]()
		_, exists := ht.Get("nonexistent")
		if exists {
			t.Error("Get(nonexistent) returned exists, want not exists")
		}
	})

	t.Run("collision handling", func(t *testing.T) {
		ht := NewHashTable[string, string]()

		// These keys should produce the same hash
		ht.Put("ABC", "First")
		ht.Put("CBA", "Second")

		// Both values should be retrievable
		if v, _ := ht.Get("ABC"); v != "First" {
			t.Errorf("Get(ABC) = %q, want First", v)
		}
		if v, _ := ht.Get("CBA"); v != "Second" {
			t.Errorf("Get(CBA) = %q, want Second", v)
		}
	})
}
