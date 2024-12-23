package linkedlist

import (
	"fmt"
	"strings"
)

// Node represents a node in the linked list
type Node[T any] struct {
	data T
	next *Node[T]
}

// LinkedList represents a generic linked list
type LinkedList[T any] struct {
	head *Node[T]
}

// NewLinkedList creates a new empty linked list
func NewLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{
		head: nil,
	}
}

// Add adds a new node to the end of the list
func (ll *LinkedList[T]) Add(data T) {
	node := &Node[T]{
		data: data,
		next: nil,
	}

	if ll.head == nil {
		ll.head = node
		return
	}

	last := ll.head
	for last.next != nil {
		last = last.next
	}
	last.next = node
}

// String returns a string representation of the linked list
func (ll *LinkedList[T]) String() string {
	var builder strings.Builder
	builder.WriteString("[")

	currentNode := ll.head
	for currentNode != nil {
		builder.WriteString(fmt.Sprintf("%v", currentNode.data))
		if currentNode.next != nil {
			builder.WriteString(" -> ")
		}
		currentNode = currentNode.next
	}
	builder.WriteString("]")

	return builder.String()
}

// Reverse reverses the linked list in place
func (ll *LinkedList[T]) Reverse() {
	if ll.head == nil {
		return
	}

	var prev *Node[T]
	current := ll.head
	var next *Node[T]

	for current != nil {
		next = current.next
		current.next = prev
		prev = current
		current = next
	}
	ll.head = prev
}
