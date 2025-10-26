package dictionary_test

import (
	TDADictionary "adts/dictionary"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

var VOLUME_SIZES = []int{12500, 25000, 50000, 100000, 200000, 400000}

func stringEquality(a, b string) bool {
	return a == b
}

func intEquality(a, b int) bool {
	return a == b
}

func TestEmptyDictionary(t *testing.T) {
	t.Log("Check that empty Dictionary has no keys")
	dict := TDADictionary.CreateHash[string, string](stringEquality)
	require.EqualValues(t, 0, dict.Count())
	require.False(t, dict.Belongs("A"))
	require.PanicsWithValue(t, "The key does not belong to the dictionary", func() { dict.Get("A") })
	require.PanicsWithValue(t, "The key does not belong to the dictionary", func() { dict.Delete("A") })
}

func TestDictionaryDefaultKey(t *testing.T) {
	t.Log("Test on an empty Hash that if we search for the default value of the data type, " +
		"it still doesn't exist")
	dict := TDADictionary.CreateHash[string, string](stringEquality)
	require.False(t, dict.Belongs(""))
	require.PanicsWithValue(t, "The key does not belong to the dictionary", func() { dict.Get("") })
	require.PanicsWithValue(t, "The key does not belong to the dictionary", func() { dict.Delete("") })

	dictNum := TDADictionary.CreateHash[int, string](intEquality)
	require.False(t, dictNum.Belongs(0))
	require.PanicsWithValue(t, "The key does not belong to the dictionary", func() { dictNum.Get(0) })
	require.PanicsWithValue(t, "The key does not belong to the dictionary", func() { dictNum.Delete(0) })
}

func TestOneElement(t *testing.T) {
	t.Log("Check that Dictionary with one element has only that Key")
	dict := TDADictionary.CreateHash[string, int](stringEquality)
	dict.Save("A", 10)
	require.EqualValues(t, 1, dict.Count())
	require.True(t, dict.Belongs("A"))
	require.False(t, dict.Belongs("B"))
	require.EqualValues(t, 10, dict.Get("A"))
	require.PanicsWithValue(t, "The key does not belong to the dictionary", func() { dict.Get("B") })
}

func TestDictionarySave(t *testing.T) {
	t.Log("Save some elements in the dictionary, and check that it works correctly at all times")
	key1 := "Cat"
	key2 := "Dog"
	key3 := "Cow"
	value1 := "meow"
	value2 := "woof"
	value3 := "moo"
	keys := []string{key1, key2, key3}
	values := []string{value1, value2, value3}

	dict := TDADictionary.CreateHash[string, string](stringEquality)
	require.False(t, dict.Belongs(keys[0]))
	require.False(t, dict.Belongs(keys[0]))
	dict.Save(keys[0], values[0])
	require.EqualValues(t, 1, dict.Count())
	require.True(t, dict.Belongs(keys[0]))
	require.True(t, dict.Belongs(keys[0]))
	require.EqualValues(t, values[0], dict.Get(keys[0]))
	require.EqualValues(t, values[0], dict.Get(keys[0]))

	require.False(t, dict.Belongs(keys[1]))
	require.False(t, dict.Belongs(keys[2]))
	dict.Save(keys[1], values[1])
	require.True(t, dict.Belongs(keys[0]))
	require.True(t, dict.Belongs(keys[1]))
	require.EqualValues(t, 2, dict.Count())
	require.EqualValues(t, values[0], dict.Get(keys[0]))
	require.EqualValues(t, values[1], dict.Get(keys[1]))

	require.False(t, dict.Belongs(keys[2]))
	dict.Save(keys[2], values[2])
	require.True(t, dict.Belongs(keys[0]))
	require.True(t, dict.Belongs(keys[1]))
	require.True(t, dict.Belongs(keys[2]))
	require.EqualValues(t, 3, dict.Count())
	require.EqualValues(t, values[0], dict.Get(keys[0]))
	require.EqualValues(t, values[1], dict.Get(keys[1]))
	require.EqualValues(t, values[2], dict.Get(keys[2]))
}

func TestValueReplacement(t *testing.T) {
	t.Log("Save a couple of keys, and then save again, checking that the value has been replaced")
	key := "Cat"
	key2 := "Dog"
	dict := TDADictionary.CreateHash[string, string](stringEquality)
	dict.Save(key, "meow")
	dict.Save(key2, "woof")
	require.True(t, dict.Belongs(key))
	require.True(t, dict.Belongs(key2))
	require.EqualValues(t, "meow", dict.Get(key))
	require.EqualValues(t, "woof", dict.Get(key2))
	require.EqualValues(t, 2, dict.Count())

	dict.Save(key, "mew")
	dict.Save(key2, "bark")
	require.True(t, dict.Belongs(key))
	require.True(t, dict.Belongs(key2))
	require.EqualValues(t, 2, dict.Count())
	require.EqualValues(t, "mew", dict.Get(key))
	require.EqualValues(t, "bark", dict.Get(key2))
}

func TestValueReplacementHopscotch(t *testing.T) {
	t.Log("Save many keys, and then replace their values. Then validate that all values are " +
		"correct. For a Hopscotch implementation, detects errors when making space or saving elements.")

	dict := TDADictionary.CreateHash[int, int](intEquality)
	for i := 0; i < 500; i++ {
		dict.Save(i, i)
	}
	for i := 0; i < 500; i++ {
		dict.Save(i, 2*i)
	}
	ok := true
	for i := 0; i < 500 && ok; i++ {
		ok = dict.Get(i) == 2*i
	}
	require.True(t, ok, "Elements were not updated correctly")
}

func TestDictionaryDelete(t *testing.T) {
	t.Log("Save some elements in the dictionary, and delete them, checking that at all times " +
		"the dictionary behaves appropriately")
	key1 := "Cat"
	key2 := "Dog"
	key3 := "Cow"
	value1 := "meow"
	value2 := "woof"
	value3 := "moo"
	keys := []string{key1, key2, key3}
	values := []string{value1, value2, value3}
	dict := TDADictionary.CreateHash[string, string](stringEquality)

	require.False(t, dict.Belongs(keys[0]))
	require.False(t, dict.Belongs(keys[0]))
	dict.Save(keys[0], values[0])
	dict.Save(keys[1], values[1])
	dict.Save(keys[2], values[2])

	require.True(t, dict.Belongs(keys[2]))
	require.EqualValues(t, values[2], dict.Delete(keys[2]))
	require.PanicsWithValue(t, "The key does not belong to the dictionary", func() { dict.Delete(keys[2]) })
	require.EqualValues(t, 2, dict.Count())
	require.False(t, dict.Belongs(keys[2]))

	require.True(t, dict.Belongs(keys[0]))
	require.EqualValues(t, values[0], dict.Delete(keys[0]))
	require.PanicsWithValue(t, "The key does not belong to the dictionary", func() { dict.Delete(keys[0]) })
	require.EqualValues(t, 1, dict.Count())
	require.False(t, dict.Belongs(keys[0]))
	require.PanicsWithValue(t, "The key does not belong to the dictionary", func() { dict.Get(keys[0]) })

	require.True(t, dict.Belongs(keys[1]))
	require.EqualValues(t, values[1], dict.Delete(keys[1]))
	require.PanicsWithValue(t, "The key does not belong to the dictionary", func() { dict.Delete(keys[1]) })
	require.EqualValues(t, 0, dict.Count())
	require.False(t, dict.Belongs(keys[1]))
	require.PanicsWithValue(t, "The key does not belong to the dictionary", func() { dict.Get(keys[1]) })
}

func TestReuseOfDeleted(t *testing.T) {
	t.Log("White box test: checks, for the case of a ClosedHash, that there is no problem " +
		"reinserting a deleted element")
	dict := TDADictionary.CreateHash[string, string](stringEquality)
	key := "hello"
	dict.Save(key, "world!")
	dict.Delete(key)
	require.EqualValues(t, 0, dict.Count())
	require.False(t, dict.Belongs(key))
	dict.Save(key, "world!!")
	require.True(t, dict.Belongs(key))
	require.EqualValues(t, 1, dict.Count())
	require.EqualValues(t, "world!!", dict.Get(key))
}

func TestWithNumericKeys(t *testing.T) {
	t.Log("Validates that it doesn't only work with strings")
	dict := TDADictionary.CreateHash[int, string](intEquality)
	key := 10
	value := "Kitten"

	dict.Save(key, value)
	require.EqualValues(t, 1, dict.Count())
	require.True(t, dict.Belongs(key))
	require.EqualValues(t, value, dict.Get(key))
	require.EqualValues(t, value, dict.Delete(key))
	require.False(t, dict.Belongs(key))
}

func TestWithStructKeys(t *testing.T) {
	t.Log("Validates that it also works with more complex structures")
	type basic struct {
		a string
		b int
	}
	type advanced struct {
		w int
		x basic
		y basic
		z string
	}

	dict := TDADictionary.CreateHash[advanced, int](func(a, b advanced) bool {
		return a.w == b.w && a.z == b.z && a.x.a == b.x.a && a.x.b == b.x.b && a.y.a == b.y.a && a.y.b == b.y.b
	})

	a1 := advanced{w: 10, z: "hello", x: basic{a: "world", b: 8}, y: basic{a: "!", b: 10}}
	a2 := advanced{w: 10, z: "aloh", x: basic{a: "dlrow", b: 14}, y: basic{a: "!", b: 5}}
	a3 := advanced{w: 10, z: "hello", x: basic{a: "world", b: 8}, y: basic{a: "!", b: 4}}

	dict.Save(a1, 0)
	dict.Save(a2, 1)
	dict.Save(a3, 2)

	require.True(t, dict.Belongs(a1))
	require.True(t, dict.Belongs(a2))
	require.True(t, dict.Belongs(a3))
	require.EqualValues(t, 0, dict.Get(a1))
	require.EqualValues(t, 1, dict.Get(a2))
	require.EqualValues(t, 2, dict.Get(a3))
	dict.Save(a1, 5)
	require.EqualValues(t, 5, dict.Get(a1))
	require.EqualValues(t, 2, dict.Get(a3))
	require.EqualValues(t, 5, dict.Delete(a1))
	require.False(t, dict.Belongs(a1))
	require.EqualValues(t, 2, dict.Get(a3))

}

func TestEmptyKey(t *testing.T) {
	t.Log("We save an empty key (i.e. \"\") and it should work without problems")
	dict := TDADictionary.CreateHash[string, string](stringEquality)
	key := ""
	dict.Save(key, key)
	require.True(t, dict.Belongs(key))
	require.EqualValues(t, 1, dict.Count())
	require.EqualValues(t, key, dict.Get(key))
}

func TestNilValue(t *testing.T) {
	t.Log("We test that the value can be nil without problems")
	dict := TDADictionary.CreateHash[string, *int](stringEquality)
	key := "Fish"
	dict.Save(key, nil)
	require.True(t, dict.Belongs(key))
	require.EqualValues(t, 1, dict.Count())
	require.EqualValues(t, (*int)(nil), dict.Get(key))
	require.EqualValues(t, (*int)(nil), dict.Delete(key))
	require.False(t, dict.Belongs(key))
}

func TestLongStringParticular(t *testing.T) {
	t.Log("Problematic cases have been seen when using the K&R hashing function, so " +
		"a test is added with that hashing function and a very long string")
	// The '~' character is the highest value in ASCII (126).
	keys := make([]string, 10)
	str := "%d~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~" +
		"~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~"
	dict := TDADictionary.CreateHash[string, string](stringEquality)
	values := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"}
	for i := 0; i < 10; i++ {
		keys[i] = fmt.Sprintf(str, i)
		dict.Save(keys[i], values[i])
	}
	require.EqualValues(t, 10, dict.Count())

	ok := true
	for i := 0; i < 10 && ok; i++ {
		ok = dict.Get(keys[i]) == values[i]
	}

	require.True(t, ok, "Getting long key works")
}

func TestRepeatedSaveAndDelete(t *testing.T) {
	t.Log("This test saves and deletes repeatedly. We do this because a common error is not considering " +
		"deletions for resizing in a Closed Hash. If it doesn't resize, it will very likely get stuck in an infinite " +
		"loop")

	dict := TDADictionary.CreateHash[int, int](intEquality)
	for i := 0; i < 1000; i++ {
		dict.Save(i, i)
		require.True(t, dict.Belongs(i))
		dict.Delete(i)
		require.False(t, dict.Belongs(i))
	}
}

func search(key string, keys []string) int {
	for i, k := range keys {
		if k == key {
			return i
		}
	}
	return -1
}

func TestInternalIteratorKeys(t *testing.T) {
	t.Log("Validates that all keys are traversed (and only once) with the internal iterator")
	key1 := "Cat"
	key2 := "Dog"
	key3 := "Cow"
	keys := []string{key1, key2, key3}
	dict := TDADictionary.CreateHash[string, *int](stringEquality)
	dict.Save(keys[0], nil)
	dict.Save(keys[1], nil)
	dict.Save(keys[2], nil)

	cs := []string{"", "", ""}
	count := 0
	countPtr := &count

	dict.Iterate(func(key string, value *int) bool {
		cs[count] = key
		*countPtr = *countPtr + 1
		return true
	})

	require.EqualValues(t, 3, count)
	require.NotEqualValues(t, -1, search(cs[0], keys))
	require.NotEqualValues(t, -1, search(cs[1], keys))
	require.NotEqualValues(t, -1, search(cs[2], keys))
	require.NotEqualValues(t, cs[0], cs[1])
	require.NotEqualValues(t, cs[0], cs[2])
	require.NotEqualValues(t, cs[2], cs[1])
}

func TestInternalIteratorValues(t *testing.T) {
	t.Log("Validates that data is traversed correctly (and only once) with the internal iterator")
	key1 := "Cat"
	key2 := "Dog"
	key3 := "Cow"
	key4 := "LittleDonkey"
	key5 := "Hamster"

	dict := TDADictionary.CreateHash[string, int](stringEquality)
	dict.Save(key1, 6)
	dict.Save(key2, 2)
	dict.Save(key3, 3)
	dict.Save(key4, 4)
	dict.Save(key5, 5)

	factorial := 1
	ptrFactorial := &factorial
	dict.Iterate(func(_ string, value int) bool {
		*ptrFactorial *= value
		return true
	})

	require.EqualValues(t, 720, factorial)
}

func TestInternalIteratorValuesWithDeletions(t *testing.T) {
	t.Log("Validates that data is traversed correctly (and only once) with the internal iterator, without traversing deleted data")
	key0 := "Elephant"
	key1 := "Cat"
	key2 := "Dog"
	key3 := "Cow"
	key4 := "LittleDonkey"
	key5 := "Hamster"

	dict := TDADictionary.CreateHash[string, int](stringEquality)
	dict.Save(key0, 7)
	dict.Save(key1, 6)
	dict.Save(key2, 2)
	dict.Save(key3, 3)
	dict.Save(key4, 4)
	dict.Save(key5, 5)

	dict.Delete(key0)

	factorial := 1
	ptrFactorial := &factorial
	dict.Iterate(func(_ string, value int) bool {
		*ptrFactorial *= value
		return true
	})

	require.EqualValues(t, 720, factorial)
}

func executeVolumeTest(b *testing.B, n int) {
	dict := TDADictionary.CreateHash[string, int](stringEquality)

	keys := make([]string, n)
	values := make([]int, n)

	/* Insert 'n' pairs into the hash */
	for i := 0; i < n; i++ {
		values[i] = i
		keys[i] = fmt.Sprintf("%08d", i)
		dict.Save(keys[i], values[i])
	}

	require.EqualValues(b, n, dict.Count(), "The number of elements is incorrect")

	/* Verify that it returns the correct values */
	ok := true
	for i := 0; i < n; i++ {
		ok = dict.Belongs(keys[i])
		if !ok {
			break
		}
		ok = dict.Get(keys[i]) == values[i]
		if !ok {
			break
		}
	}

	require.True(b, ok, "Belongs and Get with many elements does not work correctly")
	require.EqualValues(b, n, dict.Count(), "The number of elements is incorrect")

	/* Verify that it deletes and returns the correct values */
	for i := 0; i < n; i++ {
		ok = dict.Delete(keys[i]) == values[i]
		if !ok {
			break
		}
		ok = !dict.Belongs(keys[i])
		if !ok {
			break
		}
	}

	require.True(b, ok, "Deleting many elements does not work correctly")
	require.EqualValues(b, 0, dict.Count())
}

func BenchmarkDictionary(b *testing.B) {
	b.Log("Dictionary stress test. Tests saving different amounts of elements (very large), " +
		"executing the tests many times to generate a benchmark. Validates that the count " +
		"is appropriate. Then we validate that we can get and check if each generated key belongs, " +
		"and that we can then delete without problems")
	for _, n := range VOLUME_SIZES {
		b.Run(fmt.Sprintf("Test %d elements", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				executeVolumeTest(b, n)
			}
		})
	}
}

func TestIterateEmptyDictionary(t *testing.T) {
	t.Log("Iterating over empty dictionary simply has it at the end")
	dict := TDADictionary.CreateHash[string, int](stringEquality)
	iter := dict.Iterator()
	require.False(t, iter.HasNext())
	require.PanicsWithValue(t, "The iterator has finished iterating", func() { iter.Current() })
	require.PanicsWithValue(t, "The iterator has finished iterating", func() { iter.Next() })
}

func TestDictionaryIteration(t *testing.T) {
	t.Log("We save 3 values in a Dictionary, and iterate validating that the keys are all different " +
		"but belong to the dictionary. Also the values of Current and Next are correct with each other")
	key1 := "Cat"
	key2 := "Dog"
	key3 := "Cow"
	value1 := "meow"
	value2 := "woof"
	value3 := "moo"
	keys := []string{key1, key2, key3}
	values := []string{value1, value2, value3}
	dict := TDADictionary.CreateHash[string, string](stringEquality)
	dict.Save(keys[0], values[0])
	dict.Save(keys[1], values[1])
	dict.Save(keys[2], values[2])
	iter := dict.Iterator()

	require.True(t, iter.HasNext())
	first, _ := iter.Current()
	require.NotEqualValues(t, -1, search(first, keys))

	iter.Next()
	second, second_value := iter.Current()
	require.NotEqualValues(t, -1, search(second, keys))
	require.EqualValues(t, values[search(second, keys)], second_value)
	require.NotEqualValues(t, first, second)
	require.True(t, iter.HasNext())

	iter.Next()
	require.True(t, iter.HasNext())
	third, _ := iter.Current()
	require.NotEqualValues(t, -1, search(third, keys))
	require.NotEqualValues(t, first, third)
	require.NotEqualValues(t, second, third)
	iter.Next()

	require.False(t, iter.HasNext())
	require.PanicsWithValue(t, "The iterator has finished iterating", func() { iter.Current() })
	require.PanicsWithValue(t, "The iterator has finished iterating", func() { iter.Next() })
}

func TestIteratorDoesNotReachEnd(t *testing.T) {
	t.Log("Create an iterator and don't advance it. Then create another iterator and advance it.")
	dict := TDADictionary.CreateHash[string, string](stringEquality)
	keys := []string{"A", "B", "C"}
	dict.Save(keys[0], "")
	dict.Save(keys[1], "")
	dict.Save(keys[2], "")

	dict.Iterator()
	iter2 := dict.Iterator()
	iter2.Next()
	iter3 := dict.Iterator()
	first, _ := iter3.Current()
	iter3.Next()
	second, _ := iter3.Current()
	iter3.Next()
	third, _ := iter3.Current()
	iter3.Next()
	require.False(t, iter3.HasNext())
	require.NotEqualValues(t, first, second)
	require.NotEqualValues(t, third, second)
	require.NotEqualValues(t, first, third)
	require.NotEqualValues(t, -1, search(first, keys))
	require.NotEqualValues(t, -1, search(second, keys))
	require.NotEqualValues(t, -1, search(third, keys))
}

func TestIterationAfterDeletions(t *testing.T) {
	t.Log("White box test: This test tries to verify the behavior of the open hash when " +
		"it has empty lists in its table. The iterator should ignore empty lists, advancing until " +
		"it finds a real element.")

	key1 := "Cat"
	key2 := "Dog"
	key3 := "Cow"

	dict := TDADictionary.CreateHash[string, string](stringEquality)
	dict.Save(key1, "")
	dict.Save(key2, "")
	dict.Save(key3, "")
	dict.Delete(key1)
	dict.Delete(key2)
	dict.Delete(key3)
	iter := dict.Iterator()

	require.False(t, iter.HasNext())
	require.PanicsWithValue(t, "The iterator has finished iterating", func() { iter.Current() })
	require.PanicsWithValue(t, "The iterator has finished iterating", func() { iter.Next() })
	dict.Save(key1, "A")
	iter = dict.Iterator()

	require.True(t, iter.HasNext())
	c1, v1 := iter.Current()
	require.EqualValues(t, key1, c1)
	require.EqualValues(t, "A", v1)
	iter.Next()
	require.False(t, iter.HasNext())
}

func executeVolumeIteratorTests(b *testing.B, n int) {
	dict := TDADictionary.CreateHash[string, *int](stringEquality)

	keys := make([]string, n)
	values := make([]int, n)

	/* Insert 'n' pairs into the hash */
	for i := 0; i < n; i++ {
		keys[i] = fmt.Sprintf("%08d", i)
		values[i] = i
		dict.Save(keys[i], &values[i])
	}

	// Test iteration over stored keys.
	iter := dict.Iterator()
	require.True(b, iter.HasNext())

	ok := true
	var i int
	var key string
	var value *int

	for i = 0; i < n; i++ {
		if !iter.HasNext() {
			ok = false
			break
		}
		c1, v1 := iter.Current()
		key = c1
		if key == "" {
			ok = false
			break
		}
		value = v1
		if value == nil {
			ok = false
			break
		}
		*value = n
		iter.Next()
	}
	require.True(b, ok, "Iteration in volume does not work correctly")
	require.EqualValues(b, n, i, "The entire length was not traversed")
	require.False(b, iter.HasNext(), "The iterator should be at the end after traversing")

	ok = true
	for i = 0; i < n; i++ {
		if values[i] != n {
			ok = false
			break
		}
	}
	require.True(b, ok, "Not all elements were changed")
}

func BenchmarkIterator(b *testing.B) {
	b.Log("Stress test of the Dictionary Iterator. Tests saving different amounts of elements " +
		"(very large) b.N elements, iterating all of them without problems. Each test is executed b.N times to generate " +
		"a benchmark")
	for _, n := range VOLUME_SIZES {
		b.Run(fmt.Sprintf("Test %d elements", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				executeVolumeIteratorTests(b, n)
			}
		})
	}
}

func TestVolumeIteratorCutoff(t *testing.T) {
	t.Log("Volume test of internal iterator, to validate that whenever it is indicated to cut " +
		"the iteration with the visit function, it cuts")

	dict := TDADictionary.CreateHash[int, int](intEquality)

	/* Insert 'n' pairs into the hash */
	for i := 0; i < 10000; i++ {
		dict.Save(i, i)
	}

	continueExecuting := true
	continuedExecutingWhenItShouldnt := false

	dict.Iterate(func(k int, v int) bool {
		if !continueExecuting {
			continuedExecutingWhenItShouldnt = true
			return false
		}
		if k%100 == 0 {
			continueExecuting = false
			return false
		}
		return true
	})

	require.False(t, continueExecuting, "An element that generates the cutoff should have been found")
	require.False(t, continuedExecutingWhenItShouldnt,
		"It should not have continued executing if we found an element that made the iteration cut")
}
