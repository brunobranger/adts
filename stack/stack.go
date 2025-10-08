package stack

// Stack represents an abstract data type for a stack (LIFO structure).
type Stack[T any] interface {
	// IsEmpty returns true if the stack has no elements, false otherwise.
	IsEmpty() bool

	// Top returns the element at the top of the stack.
	// If the stack is empty, it panics with "The stack is empty".
	Top() T

	// Push adds a new element to the top of the stack.
	Push(T)

	// Pop removes and returns the element at the top of the stack.
	// If the stack is empty, it panics with "The stack is empty".
	Pop() T
}
