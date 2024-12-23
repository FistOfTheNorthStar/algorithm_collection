package linkedlist

import (
	"testing"
)

func TestLinkedList(t *testing.T) {
	t.Run("add and reverse strings", func(t *testing.T) {
		list := NewLinkedList[string]()
		list.Add("s1")
		list.Add("s2")
		list.Add("s3")
		list.Add("s4")

		if got := list.String(); got != "[s1 -> s2 -> s3 -> s4]" {
			t.Errorf("Before reverse = %v, want [s1 -> s2 -> s3 -> s4]", got)
		}

		list.Reverse()

		if got := list.String(); got != "[s4 -> s3 -> s2 -> s1]" {
			t.Errorf("After reverse = %v, want [s4 -> s3 -> s2 -> s1]", got)
		}
	})

	t.Run("add and reverse integers", func(t *testing.T) {
		list := NewLinkedList[int]()
		list.Add(1)
		list.Add(2)
		list.Add(3)
		list.Add(5)

		if got := list.String(); got != "[1 -> 2 -> 3 -> 5]" {
			t.Errorf("Before reverse = %v, want [1 -> 2 -> 3 -> 5]", got)
		}

		list.Reverse()

		if got := list.String(); got != "[5 -> 3 -> 2 -> 1]" {
			t.Errorf("After reverse = %v, want [5 -> 3 -> 2 -> 1]", got)
		}
	})

	t.Run("reverse empty list", func(t *testing.T) {
		list := NewLinkedList[string]()
		list.Reverse() // Should not panic
		if got := list.String(); got != "[]" {
			t.Errorf("Empty list = %v, want []", got)
		}
	})

	t.Run("reverse single element", func(t *testing.T) {
		list := NewLinkedList[int]()
		list.Add(1)
		list.Reverse()
		if got := list.String(); got != "[1]" {
			t.Errorf("Single element = %v, want [1]", got)
		}
	})
}
