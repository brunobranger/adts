package queue_test

import (
	QueuePkg "adts/queue"
	StackPkg "adts/stack"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	ELEM_COUNT  = 100
	VOLUME_SIZE = 10000
)

type Person struct {
	Name string
	Age  int
}

type Animal struct {
	Name    string
	Age     int
	Species string
}

func TestEmptyQueue(t *testing.T) {
	queue := QueuePkg.NewLinkedQueue[int]()
	require.True(t, queue.IsEmpty())

	require.PanicsWithValue(t, "The queue is empty", func() { queue.Dequeue() })
	require.PanicsWithValue(t, "The queue is empty", func() { queue.Front() })
}

func TestSingleElement(t *testing.T) {
	queue := QueuePkg.NewLinkedQueue[int]()
	queue.Enqueue(1)

	require.Equal(t, 1, queue.Front())
	require.Equal(t, 1, queue.Dequeue())
}

func TestMultipleElements(t *testing.T) {
	queue := QueuePkg.NewLinkedQueue[int]()
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	require.Equal(t, 1, queue.Front())
	require.Equal(t, 1, queue.Dequeue())

	require.Equal(t, 2, queue.Front())
	require.Equal(t, 2, queue.Dequeue())

	require.Equal(t, 3, queue.Front())
	require.Equal(t, 3, queue.Dequeue())
}

func TestFIFO(t *testing.T) {
	queue := QueuePkg.NewLinkedQueue[int]()
	for i := 1; i < ELEM_COUNT; i++ {
		queue.Enqueue(i)
		require.Equal(t, i, queue.Dequeue())
	}
}

func TestFIFOWithAuxiliaryStructure(t *testing.T) {
	stack := StackPkg.NewDynamicStack[int]()
	queue := QueuePkg.NewLinkedQueue[int]()

	for i := 1; i <= ELEM_COUNT; i++ {
		queue.Enqueue(i)
	}

	for i := 1; i <= ELEM_COUNT; i++ {
		stack.Push(queue.Dequeue())
	}

	for i := 1; i <= ELEM_COUNT; i++ {
		queue.Enqueue(stack.Pop())
	}

	require.Equal(t, ELEM_COUNT, queue.Front())
}

func TestVolume(t *testing.T) {
	queue := QueuePkg.NewLinkedQueue[int]()
	for i := 1; i <= VOLUME_SIZE; i++ {
		queue.Enqueue(i)
		require.Equal(t, 1, queue.Front())
	}
	for i := 1; i <= VOLUME_SIZE; i++ {
		require.Equal(t, i, queue.Front())
		require.Equal(t, i, queue.Dequeue())
	}
	require.True(t, queue.IsEmpty())
	require.Panics(t, func() { queue.Front() })
	require.Panics(t, func() { queue.Dequeue() })
}

func TestReusableQueue(t *testing.T) {
	queue := QueuePkg.NewLinkedQueue[int]()

	queue.Enqueue(1)
	queue.Enqueue(2)

	val := queue.Dequeue()
	require.Equal(t, 1, val)

	queue.Enqueue(3)

	val = queue.Dequeue()
	require.Equal(t, 2, val)

	val = queue.Dequeue()
	require.Equal(t, 3, val)

	require.True(t, queue.IsEmpty())
}

func TestQueueWithStructs(t *testing.T) {
	testQueue(t, []Person{
		{Name: "Bruno", Age: 19},
		{Name: "Abril", Age: 18},
	})

	testQueue(t, []Animal{
		{Name: "Rocco", Age: 12, Species: "Dog"},
		{Name: "Mia", Age: 4, Species: "Cat"},
		{Name: "Akira", Age: 1, Species: "Cat"},
	})
}

func TestQueueWithPointers(t *testing.T) {
	a, b := 10, 20
	testQueue(t, []*int{&a, &b})
}

func TestQueueWithFloats(t *testing.T) {
	testQueue(t, []float64{3.14, 2.71, 1.41})
}

func TestQueueWithBool(t *testing.T) {
	testQueue(t, []bool{true, false})
}

func TestQueueWithNil(t *testing.T) {
	testQueue(t, []*int{nil})
}

// Helper function to test a queue with generic elements
func testQueue[T comparable](t *testing.T, elements []T) {
	queue := QueuePkg.NewLinkedQueue[T]()

	for _, e := range elements {
		queue.Enqueue(e)
		require.Equal(t, elements[0], queue.Front())
	}

	for _, expected := range elements {
		val := queue.Dequeue()
		require.Equal(t, expected, val)
	}

	require.True(t, queue.IsEmpty())
}
