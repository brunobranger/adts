package stack_test

import (
	"adts/stack" // ajusta la ruta según tu repositorio
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEmptyStack(t *testing.T) {
	stack := stack.NewDynamicStack[int]()
	require.True(t, stack.IsEmpty())

	require.Panics(t, func() { stack.Top() })
	require.Panics(t, func() { stack.Pop() })
}

func TestPushPop(t *testing.T) {
	stack := stack.NewDynamicStack[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	require.Equal(t, 3, stack.Top())
	require.Equal(t, 3, stack.Pop())
	require.Equal(t, 2, stack.Top())
}

func TestVolume(t *testing.T) {
	stack := stack.NewDynamicStack[int]()
	n := 10000
	for i := 0; i < n; i++ {
		stack.Push(i)
		require.Equal(t, i, stack.Top())
	}

	for i := n - 1; i >= 0; i-- {
		require.Equal(t, i, stack.Pop())
	}

	require.True(t, stack.IsEmpty())
}

func TestAlternatingPushPop(t *testing.T) {
	stack := stack.NewDynamicStack[int]()
	for i := 0; i < 10; i++ {
		stack.Push(i)
		require.Equal(t, i, stack.Pop())
		require.True(t, stack.IsEmpty())
	}
}

func TestReverseOrder(t *testing.T) {
	elements := []int{1, 2, 3, 4, 5}
	stack := stack.NewDynamicStack[int]()

	for _, elem := range elements {
		stack.Push(elem)
	}

	for i := len(elements) - 1; i >= 0; i-- {
		require.Equal(t, elements[i], stack.Pop())
	}
	require.True(t, stack.IsEmpty())
}

func TestStackResizing(t *testing.T) {
	stack := stack.NewDynamicStack[int]()

	for i := 0; i < 20; i++ {
		stack.Push(i)
		require.Equal(t, i, stack.Top())
	}

	for i := 19; i >= 5; i-- {
		require.Equal(t, i, stack.Pop())
	}

	require.False(t, stack.IsEmpty())

	for i := 4; i >= 0; i-- {
		require.Equal(t, i, stack.Pop())
	}

	require.True(t, stack.IsEmpty())
}

func TestStackWithStrings(t *testing.T) {
	stack := stack.NewDynamicStack[string]()
	stack.Push("Linkin")
	stack.Push("Park")
	stack.Push("!")

	require.Equal(t, "!", stack.Pop())
	require.Equal(t, "Park", stack.Pop())
	require.Equal(t, "Linkin", stack.Pop())

	require.True(t, stack.IsEmpty())
}

type Person struct {
	Name string
	Age  int
}

func TestStackWithStructs(t *testing.T) {
	stack := stack.NewDynamicStack[Person]()
	stack.Push(Person{"Bruno", 19})
	stack.Push(Person{"Abril", 18})

	require.Equal(t, Person{"Abril", 18}, stack.Pop())
	require.Equal(t, Person{"Bruno", 19}, stack.Pop())

	require.True(t, stack.IsEmpty())
}
