package list_test

import (
	ListModule "adts/list"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	_PANIC_LIST_MSG = "The list is empty"
	_PANIC_ITER_MSG = "Iterator reached the end"
)

// --------------------------------------------------------------------
// -------------------- LINKED LIST TESTS ------------------------------
// --------------------------------------------------------------------

// Test empty list behavior
func TestListIsEmpty(t *testing.T) {
	list := ListModule.CreateLinkedList[int]()
	require.True(t, list.IsEmpty())
	require.Equal(t, 0, list.Length())

	require.Panics(t, func() { list.PeekFirst() })
	require.Panics(t, func() { list.PeekLast() })
	require.Panics(t, func() { list.RemoveFirst() })
}

// Test insert and peek first element
func TestInsertAndPeekFirst(t *testing.T) {
	list := ListModule.CreateLinkedList[int]()

	list.InsertFirst(10)
	require.Equal(t, 10, list.PeekFirst())

	list.InsertFirst(20)
	require.Equal(t, 20, list.PeekFirst())
	require.Equal(t, 10, list.PeekLast()) // Previously first element moved to last
}

// Test insert and peek last element
func TestInsertAndPeekLast(t *testing.T) {
	list := ListModule.CreateLinkedList[int]()

	list.InsertLast(10)
	require.Equal(t, 10, list.PeekLast())

	list.InsertLast(20)
	require.Equal(t, 20, list.PeekLast())
	require.Equal(t, 10, list.PeekFirst()) // First element remains
}

// Test interleaved insertions
func TestInsertInterleaved(t *testing.T) {
	list := ListModule.CreateLinkedList[int]()

	for n := 0; n < 5; n++ {
		list.InsertFirst(5 - n)
		require.Equal(t, 5-n, list.PeekFirst())

		list.InsertLast(6 + n)
		require.Equal(t, 6+n, list.PeekLast())
	}

	require.Equal(t, 1, list.PeekFirst())
	require.Equal(t, 10, list.PeekLast())

	for n := 0; n < 10; n++ {
		require.Equal(t, n+1, list.RemoveFirst())
	}

	require.PanicsWithValue(t, _PANIC_LIST_MSG, func() { list.PeekFirst() })
	require.PanicsWithValue(t, _PANIC_LIST_MSG, func() { list.PeekLast() })
}

// Test remove first element
func TestRemoveFirstList(t *testing.T) {
	list := ListModule.CreateLinkedList[int]()

	for n := 0; n < 10; n++ {
		list.InsertLast(n)
	}

	for n := 0; n < 9; n++ {
		require.Equal(t, n, list.RemoveFirst())
		require.Equal(t, n+1, list.PeekFirst())
		require.Equal(t, 9-n, list.Length())
	}

	require.Equal(t, 9, list.RemoveFirst())

	require.PanicsWithValue(t, _PANIC_LIST_MSG, func() { list.PeekFirst() })
	require.PanicsWithValue(t, _PANIC_LIST_MSG, func() { list.PeekLast() })
	require.PanicsWithValue(t, _PANIC_LIST_MSG, func() { list.RemoveFirst() })
}

// Volume test
func TestVolume(t *testing.T) {
	list := ListModule.CreateLinkedList[int]()
	n := 10000

	for i := 0; i < n; i++ {
		list.InsertFirst(i)
		require.Equal(t, i, list.PeekFirst())
	}

	for i := n - 1; i >= 0; i-- {
		require.Equal(t, i, list.RemoveFirst())
	}

	require.True(t, list.IsEmpty())
	require.Panics(t, func() { list.PeekFirst() })
	require.Panics(t, func() { list.RemoveFirst() })
}

// -------------------- INTERNAL ITERATOR TESTS ------------------------

// Sum all elements
func TestSumAll(t *testing.T) {
	list := ListModule.CreateLinkedList[int]()
	arr := []int{0, 10, 20, 30, 40, -50}

	for _, n := range arr {
		list.InsertLast(n)
	}

	sum := 0
	list.Iterate(func(n int) bool {
		sum += n
		return true
	})

	require.Equal(t, 50, sum)
}

// Sum only even numbers
func TestSumEven(t *testing.T) {
	list := ListModule.CreateLinkedList[int]()
	arr := []int{0, 10, 15, 17, 20, 21, 29, -30, 50, -53}

	for _, n := range arr {
		list.InsertLast(n)
	}

	sum := 0
	list.Iterate(func(n int) bool {
		if n%2 == 0 {
			sum += n
		}
		return true
	})

	require.Equal(t, 50, sum)
}

// Sum until a condition
func TestSumUntilSeven(t *testing.T) {
	list := ListModule.CreateLinkedList[int]()
	arr := []int{0, 0, 1, 1, 2, 7, -4}

	for _, n := range arr {
		list.InsertLast(n)
	}

	sum := 0
	list.Iterate(func(n int) bool {
		if n != 7 {
			sum += n
			return true
		}
		return false
	})

	require.Equal(t, 4, sum)
}

// Sum first 5 even numbers
func TestSumFirstFiveEven(t *testing.T) {
	list := ListModule.CreateLinkedList[int]()
	arr := []int{0, 0, 1, 3, 5, 246, 7, -246, 100, -100, 13, 15}

	for _, n := range arr {
		list.InsertLast(n)
	}

	sum, count := 0, 0
	list.Iterate(func(n int) bool {
		if count < 5 {
			if n%2 == 0 {
				sum += n
				count++
			}
			return true
		}
		return false
	})

	require.Equal(t, 100, sum)
}

// -------------------- EXTERNAL ITERATOR TESTS ------------------------

// Iterates elements in order
func TestExternalIteratorIterates(t *testing.T) {
	arr := []int{5, 10, 15, 20, 25}
	list := ListModule.CreateLinkedList[int]()
	for _, v := range arr {
		list.InsertLast(v)
	}

	iter := list.Iterator()
	var result []int
	for iter.HasNext() {
		result = append(result, iter.Current())
		iter.Next()
	}

	require.Equal(t, arr, result)
}

// Test Current() method
func TestCurrent(t *testing.T) {
	list := ListModule.CreateLinkedList[int]()

	for n := 0; n < 10; n++ {
		list.InsertLast(n)
	}

	iter := list.Iterator()
	num := 0

	for iter.HasNext() {
		require.Equal(t, num, iter.Current())
		iter.Next()
		num++
	}

	require.PanicsWithValue(t, _PANIC_ITER_MSG, func() { iter.Current() })
	require.PanicsWithValue(t, _PANIC_ITER_MSG, func() { iter.Next() })
	require.PanicsWithValue(t, _PANIC_ITER_MSG, func() { iter.Remove() })
}

// Test HasNext() method
func TestHasNext(t *testing.T) {
	list := ListModule.CreateLinkedList[int]()
	list.InsertLast(1)
	list.InsertLast(2)

	iter := list.Iterator()
	require.True(t, iter.HasNext())

	iter.Next()
	require.True(t, iter.HasNext())

	iter.Next()
	require.False(t, iter.HasNext())
}

// Test Next() method
func TestNext(t *testing.T) {
	list := ListModule.CreateLinkedList[string]()
	list.InsertLast("A")
	list.InsertLast("B")

	iter := list.Iterator()
	require.Equal(t, "A", iter.Current())

	iter.Next()
	require.Equal(t, "B", iter.Current())

	iter.Next()
	require.Panics(t, func() { iter.Current() })
}

// Test insert at beginning using iterator
func TestIteratorInsertFirst(t *testing.T) {
	list := ListModule.CreateLinkedList[string]()
	iter := list.Iterator()

	iter.Insert("First")
	require.Equal(t, "First", list.PeekFirst())

	iter.Insert("Before First")
	require.Equal(t, "Before First", list.PeekFirst())
}

// Test insert in the middle using iterator
func TestIteratorInsertMiddle(t *testing.T) {
	list := ListModule.CreateLinkedList[string]()
	arr := []string{"First", "Fourth"}

	for _, s := range arr {
		list.InsertLast(s)
	}

	iter := list.Iterator()
	iter.Next()           // points to "Fourth"
	iter.Insert("Second") // inserted between "First" and "Fourth"
	require.Equal(t, "Second", iter.Current())

	iter.Next()
	iter.Insert("Third")
	require.Equal(t, "Third", iter.Current())
}

// Test insert at the end using iterator
func TestIteratorInsertLast(t *testing.T) {
	list := ListModule.CreateLinkedList[string]()
	arr := []string{"First", "Second", "Third"}

	for _, s := range arr {
		list.InsertLast(s)
	}

	iter := list.Iterator()
	for iter.HasNext() {
		iter.Next()
	}

	iter.Insert("Fourth")
	require.Equal(t, "Fourth", list.PeekLast())

	iter.Next()
	iter.Insert("After Fourth")
	require.Equal(t, "After Fourth", list.PeekLast())
}

// Test remove first element using iterator
func TestIteratorRemoveFirst(t *testing.T) {
	list := ListModule.CreateLinkedList[int]()
	list.InsertLast(100)
	list.InsertLast(200)

	iter := list.Iterator()
	require.Equal(t, 100, iter.Remove())

	require.Equal(t, 200, list.PeekFirst())
	require.Equal(t, 1, list.Length())
}

// Test remove middle element using iterator
func TestIteratorRemoveMiddle(t *testing.T) {
	list := ListModule.CreateLinkedList[int]()
	list.InsertLast(100)
	list.InsertLast(200)
	list.InsertLast(300)

	iter := list.Iterator()
	iter.Next() // move to 200

	require.Equal(t, 200, iter.Remove())
	require.Equal(t, 300, iter.Current())

	require.Equal(t, 100, list.PeekFirst())
	require.Equal(t, 300, list.PeekLast())
	require.Equal(t, 2, list.Length())
}

// Test remove last element using iterator
func TestIteratorRemoveLast(t *testing.T) {
	list := ListModule.CreateLinkedList[int]()
	list.InsertLast(100)
	list.InsertLast(200)

	iter := list.Iterator()
	iter.Next() // move to last

	require.Equal(t, 200, iter.Remove())
	require.False(t, iter.HasNext())

	require.Equal(t, 100, list.PeekFirst())
	require.Equal(t, 100, list.PeekLast())
	require.Equal(t, 1, list.Length())
}
