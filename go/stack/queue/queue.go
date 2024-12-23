package stack

// QueueViaStacks implements a queue using two stacks
type QueueViaStacks[T any] struct {
	inbox  []T // for pushing new elements
	outbox []T // for popping elements in FIFO order
}

// NewQueueViaStacks creates a new QueueViaStacks instance
func NewQueueViaStacks[T any]() *QueueViaStacks[T] {
	return &QueueViaStacks[T]{
		inbox:  make([]T, 0),
		outbox: make([]T, 0),
	}
}

// Add pushes a new value onto the queue
func (q *QueueViaStacks[T]) Add(value T) {
	// Add to inbox stack (newest elements on top)
	q.inbox = append(q.inbox, value)
}

// Peek returns and removes the oldest element from the queue
func (q *QueueViaStacks[T]) Peek() (T, error) {
	if len(q.outbox) == 0 {
		// Transfer elements from inbox to outbox in reverse order
		for len(q.inbox) > 0 {
			lastIdx := len(q.inbox) - 1
			value := q.inbox[lastIdx]
			q.inbox = q.inbox[:lastIdx]
			q.outbox = append(q.outbox, value)
		}
	}

	if len(q.outbox) == 0 {
		var zero T
		return zero, ErrEmptyQueue
	}

	// Remove and return the last element from outbox
	lastIdx := len(q.outbox) - 1
	value := q.outbox[lastIdx]
	q.outbox = q.outbox[:lastIdx]
	return value, nil
}
