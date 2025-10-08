package queue

// Queue represents an abstract data type for a queue (FIFO structure).
type Queue[T any] interface {
	// IsEmpty returns true if the queue has no elements, false otherwise.
	IsEmpty() bool

	// Front returns the element at the front of the queue.
	// If the queue is empty, it panics with "The queue is empty".
	Front() T

	// Enqueue adds a new element to the end of the queue.
	Enqueue(T)

	// Dequeue removes and returns the element at the front of the queue.
	// If the queue is empty, it panics with "The queue is empty".
	Dequeue() T
}
