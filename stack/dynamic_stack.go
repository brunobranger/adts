package stack

const (
	INITIAL_CAPACITY = 10
	RESIZE_FACTOR    = 2
	SHRINK_THRESHOLD = 4
)

type dynamicStack[T any] struct {
	data []T
	size int
}

// ------------ FUNCTION TO CREATE AND RETURN THE STACK ------------ //

func NewDynamicStack[T any]() Stack[T] {
	return &dynamicStack[T]{
		data: make([]T, INITIAL_CAPACITY),
		size: 0,
	}
}

// ------------ STACK PRIMITIVES ------------ //

func (s *dynamicStack[T]) IsEmpty() bool {
	return s.size == 0
}

func (s *dynamicStack[T]) Top() T {
	if s.IsEmpty() {
		panic("The stack is empty")
	}
	return s.data[s.size-1]
}

func (s *dynamicStack[T]) Push(element T) {
	if s.size == len(s.data) {
		newCapacity := len(s.data) * RESIZE_FACTOR
		s.resize(newCapacity)
	}
	s.data[s.size] = element
	s.size++
}

func (s *dynamicStack[T]) Pop() T {
	if s.IsEmpty() {
		panic("The stack is empty")
	}
	s.size--
	top := s.data[s.size]

	if s.size*SHRINK_THRESHOLD <= len(s.data) {
		newCapacity := len(s.data) / RESIZE_FACTOR
		s.resize(newCapacity)
	}

	return top
}

// ------------ INTERNAL HELPER METHODS ------------ //

func (s *dynamicStack[T]) resize(newCapacity int) {
	if newCapacity < 1 {
		newCapacity = 1
	}
	newData := make([]T, newCapacity)
	copy(newData, s.data[:s.size])
	s.data = newData
}
