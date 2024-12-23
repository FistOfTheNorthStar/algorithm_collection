package stack

import "errors"

var ErrEmptyQueue = errors.New("queue is empty")

// queue/queue_test.go
package queue

import "testing"

func TestQueueViaStacks(t *testing.T) {
    t.Run("peek first element", func(t *testing.T) {
        // Setup
        queue := NewQueueViaStacks[int]()
        queue.Add(4)
        queue.Add(2)
        queue.Add(9)
        
        // Test
        value, err := queue.Peek()
        if err != nil {
            t.Errorf("Unexpected error: %v", err)
        }
        
        // Verify
        if value != 4 {
            t.Errorf("Peek() = %v; want 4", value)
        }
    })
    
    t.Run("peek empty queue", func(t *testing.T) {
        queue := NewQueueViaStacks[int]()
        _, err := queue.Peek()
        if err != ErrEmptyQueue {
            t.Errorf("Expected ErrEmptyQueue but got: %v", err)
        }
    })
    
    t.Run("peek multiple elements", func(t *testing.T) {
        queue := NewQueueViaStacks[int]()
        expected := []int{1, 2, 3}
        
        // Add elements
        for _, v := range expected {
            queue.Add(v)
        }
        
        // Peek and verify each element
        for _, want := range expected {
            got, err := queue.Peek()
            if err != nil {
                t.Errorf("Unexpected error: %v", err)
            }
            if got != want {
                t.Errorf("Peek() = %v; want %v", got, want)
            }
        }
    })
}