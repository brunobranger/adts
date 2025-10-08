package list

// List is a generic interface representing a list of elements of type T.
type List[T any] interface {

	// IsEmpty returns true if the list has no elements, false otherwise.
	IsEmpty() bool

	// InsertFirst inserts the element at the beginning of the list.
	InsertFirst(T)

	// InsertLast inserts the element at the end of the list.
	InsertLast(T)

	// RemoveFirst removes and returns the first element of the list.
	// Pre: The list is not empty.
	RemoveFirst() T

	// PeekFirst returns the first element without removing it.
	// Pre: The list is not empty.
	PeekFirst() T

	// PeekLast returns the last element without removing it.
	// Pre: The list is not empty.
	PeekLast() T

	// Length returns the number of elements in the list.
	Length() int

	// Iterate applies the visit function to each element until:
	// - all elements are visited, or
	// - visit returns false
	Iterate(visit func(T) bool)

	// Iterator returns an iterator for traversing the list.
	Iterator() ListIterator[T]
}

// ListIterator is an interface to iterate over a list and modify it.
type ListIterator[T any] interface {

	// Current returns the current element in the iteration.
	// Pre: There is a current element.
	Current() T

	// HasNext indicates if there is a next element to see.
	HasNext() bool

	// Next moves the iterator to the next element.
	// Pre: There is a next element.
	Next()

	// Insert adds an element at the current iterator position.
	Insert(T)

	// Remove deletes the current element and returns it.
	// Pre: There is a current element.
	Remove() T
}
