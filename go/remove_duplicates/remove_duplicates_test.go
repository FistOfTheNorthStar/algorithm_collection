package listutils

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicatesAndOrder(t *testing.T) {
	t.Run("strings test", func(t *testing.T) {
		input := []string{"c", "b", "b", "d", "c", "a"}
		expected := []string{"a", "b", "c", "d"}

		result := RemoveDuplicatesAndOrder(input)

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("integers test", func(t *testing.T) {
		input := []int{3, 3, 4, 1, 7, 1, 2, 1}
		expected := []int{1, 2, 3, 4, 7}

		result := RemoveDuplicatesAndOrder(input)

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
}
