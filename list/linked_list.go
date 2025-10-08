package list

const (
	_PANIC_LIST_MSG = "The list is empty"
	_PANIC_ITER_MSG = "Iterator reached the end"
)

type node[T any] struct {
	data T
	next *node[T]
}

type linkedList[T any] struct {
	first *node[T]
	last  *node[T]
	size  int
}

type linkedListIterator[T any] struct {
	prev    *node[T]
	current *node[T]
	list    *linkedList[T]
}

// CreateLinkedList creates and returns a new linked list.
func CreateLinkedList[T any]() List[T] {
	return &linkedList[T]{}
}

func (list *linkedList[T]) IsEmpty() bool {
	return list.first == nil
}

func (list *linkedList[T]) InsertFirst(element T) {
	newNode := newNodeLinkedList(element)
	newNode.next = list.first

	if list.IsEmpty() {
		list.last = newNode
	}

	list.first = newNode
	list.size++
}

func (list *linkedList[T]) InsertLast(element T) {
	newNode := newNodeLinkedList(element)

	if list.IsEmpty() {
		list.first = newNode
	} else {
		list.last.next = newNode
	}

	list.last = newNode
	list.size++
}

func (list *linkedList[T]) RemoveFirst() T {
	first := list.PeekFirst()

	list.first = list.first.next
	list.size--

	if list.IsEmpty() {
		list.last = nil
	}

	return first
}

func (list *linkedList[T]) PeekFirst() T {
	if list.IsEmpty() {
		panic(_PANIC_LIST_MSG)
	}
	return list.first.data
}

func (list *linkedList[T]) PeekLast() T {
	if list.IsEmpty() {
		panic(_PANIC_LIST_MSG)
	}
	return list.last.data
}

func (list *linkedList[T]) Length() int {
	return list.size
}

func (list *linkedList[T]) Iterate(visit func(T) bool) {
	for current := list.first; current != nil; current = current.next {
		if !visit(current.data) {
			return
		}
	}
}

func (list *linkedList[T]) Iterator() ListIterator[T] {
	return &linkedListIterator[T]{current: list.first, list: list}
}

func (iter *linkedListIterator[T]) Current() T {
	if !iter.HasNext() {
		panic(_PANIC_ITER_MSG)
	}
	return iter.current.data
}

func (iter *linkedListIterator[T]) HasNext() bool {
	return iter.current != nil
}

func (iter *linkedListIterator[T]) Next() {
	if !iter.HasNext() {
		panic(_PANIC_ITER_MSG)
	}
	iter.prev = iter.current
	iter.current = iter.current.next
}

func (iter *linkedListIterator[T]) Insert(element T) {
	newNode := newNodeLinkedList(element)
	newNode.next = iter.current

	if iter.current == iter.list.first {
		iter.list.first = newNode
	} else {
		iter.prev.next = newNode
	}

	if iter.prev == iter.list.last {
		iter.list.last = newNode
	}

	iter.current = newNode
	iter.list.size++
}

func (iter *linkedListIterator[T]) Remove() T {
	if !iter.HasNext() {
		panic(_PANIC_ITER_MSG)
	}

	data := iter.current.data
	next := iter.current.next

	if iter.current == iter.list.first {
		iter.list.first = next
	} else {
		iter.prev.next = next
	}

	if next == nil {
		iter.list.last = iter.prev
	}

	iter.current = next
	iter.list.size--

	return data
}

func newNodeLinkedList[T any](data T) *node[T] {
	return &node[T]{data: data}
}
