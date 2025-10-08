package queue

type linkedQueue[T any] struct {
	front *queueNode[T]
	rear  *queueNode[T]
}

type queueNode[T any] struct {
	data T
	next *queueNode[T]
}

// ------------ QUEUE OPERATIONS ------------ //

func (q *linkedQueue[T]) IsEmpty() bool {
	return q.front == nil
}

func (q *linkedQueue[T]) Front() T {
	if q.IsEmpty() {
		panic("The queue is empty")
	}
	return q.front.data
}

func (q *linkedQueue[T]) Enqueue(element T) {
	newNode := createNode(element, nil)
	if q.front == nil {
		q.front = newNode
	}
	if q.rear != nil {
		q.rear.next = newNode
	}
	q.rear = newNode
}

func (q *linkedQueue[T]) Dequeue() T {
	if q.front == nil {
		panic("The queue is empty")
	}

	data := q.front.data
	q.front = q.front.next

	if q.front == nil {
		q.rear = nil
	}

	return data
}

// ------------ HELPER FUNCTIONS ------------ //

func NewLinkedQueue[T any]() Queue[T] {
	return &linkedQueue[T]{}
}

func createNode[T any](data T, next *queueNode[T]) *queueNode[T] {
	return &queueNode[T]{data: data, next: next}
}
