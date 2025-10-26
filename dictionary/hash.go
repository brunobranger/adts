package dictionary

import (
	TDAList "adts/list"
	"fmt"
)

const (
	_PANIC_MESSAGE_DICTIONARY = "The key does not belong to the dictionary"
	_PANIC_MESSAGE_ITER       = "The iterator has finished iterating"
	_INITIAL_SIZE             = 7
	_MAX_LOAD_FACTOR          = 3.0
	_MIN_LOAD_FACTOR          = 2.0
	_RESIZE_FACTOR            = 2
)

type (
	keyValuePairList[K, V any]     TDAList.List[*keyValuePair[K, V]]
	iterKeyValuePairList[K, V any] TDAList.ListIterator[*keyValuePair[K, V]]
)

type keyValuePair[K, V any] struct {
	key   K
	value V
}

type openHash[K, V any] struct {
	table []keyValuePairList[K, V]
	size  int
	count int
	cmp   func(K, K) bool
}

type iterOpenHash[K, V any] struct {
	hash       *openHash[K, V]
	currentPos int
	current    iterKeyValuePairList[K, V]
}

func CreateHash[K, V any](cmp func(K, K) bool) Dictionary[K, V] {
	table := createTable[K, V](_INITIAL_SIZE)
	return &openHash[K, V]{table: table, size: _INITIAL_SIZE, cmp: cmp}
}

// -------------------- DICTIONARY PRIMITIVES --------------------

func (hash *openHash[K, V]) Save(key K, value V) {
	iter := hash.hashSearch(key)

	if iter.HasNext() {
		iter.Remove()
		hash.count--
	}
	iter.Insert(createPair(key, value))

	hash.count++
	if float32(hash.count)/float32(hash.size) >= _MAX_LOAD_FACTOR {
		hash.rehash(hash.size * _RESIZE_FACTOR)
	}
}

func (hash *openHash[K, V]) Belongs(key K) bool {
	iter := hash.hashSearch(key)
	return iter.HasNext()
}

func (hash *openHash[K, V]) Get(key K) V {
	iter := hash.hashSearch(key)
	if iter.HasNext() {
		return iter.Current().value
	}
	panic(_PANIC_MESSAGE_DICTIONARY)
}

func (hash *openHash[K, V]) Delete(key K) V {
	iter := hash.hashSearch(key)
	if iter.HasNext() {
		pair := iter.Remove()

		hash.count--
		if float32(hash.count)/float32(hash.size) <= _MIN_LOAD_FACTOR && hash.size > _INITIAL_SIZE {
			hash.rehash(hash.size / _RESIZE_FACTOR)
		}

		return pair.value
	}
	panic(_PANIC_MESSAGE_DICTIONARY)
}

func (hash *openHash[K, V]) Count() int {
	return hash.count
}

func (hash *openHash[K, V]) Iterate(visit func(key K, value V) bool) {
	for _, list := range hash.table {
		iterateNextList := true

		list.Iterate(func(pair *keyValuePair[K, V]) bool {
			if !visit(pair.key, pair.value) {
				iterateNextList = false
				return false
			}
			return true
		})

		if !iterateNextList {
			return
		}
	}
}

func (hash *openHash[K, V]) Iterator() DictionaryIterator[K, V] {
	iter := iterOpenHash[K, V]{hash: hash}
	iter.findList()
	return &iter
}

// ----------------- EXTERNAL ITERATOR PRIMITIVES -----------------

func (iter *iterOpenHash[K, V]) HasNext() bool {
	return iter.currentPos != iter.hash.size
}

func (iter *iterOpenHash[K, V]) Current() (K, V) {
	if !iter.HasNext() {
		panic(_PANIC_MESSAGE_ITER)
	}
	pair := iter.current.Current()
	return pair.key, pair.value
}

func (iter *iterOpenHash[K, V]) Next() {
	if !iter.HasNext() {
		panic(_PANIC_MESSAGE_ITER)
	}

	iter.current.Next()
	if iter.current.HasNext() {
		return
	}

	iter.currentPos++
	iter.findList()
}

// ----------------------- AUXILIARY FUNCTIONS -----------------------

// Creation functions

func createTable[K, V any](size int) []keyValuePairList[K, V] {
	table := make([]keyValuePairList[K, V], size)
	for i := range table {
		table[i] = TDAList.CreateLinkedList[*keyValuePair[K, V]]()
	}
	return table
}

func createPair[K, V any](key K, value V) *keyValuePair[K, V] {
	return &keyValuePair[K, V]{key, value}
}

// Dictionary functions

func (hash *openHash[K, V]) hashSearch(key K) iterKeyValuePairList[K, V] {
	pos := convertToPosition(key, hash.size)
	list := hash.table[pos]

	var iter iterKeyValuePairList[K, V]
	for iter = list.Iterator(); iter.HasNext(); iter.Next() {
		pair := iter.Current()
		if hash.cmp(pair.key, key) {
			return iter
		}
	}

	return iter
}

func (hash *openHash[K, V]) rehash(newSize int) {
	newTable := createTable[K, V](newSize)

	for _, list := range hash.table {
		for iter := list.Iterator(); iter.HasNext(); iter.Next() {
			pair := iter.Current()
			pos := convertToPosition(pair.key, newSize)
			newTable[pos].InsertLast(pair)
		}
	}

	hash.table = newTable
	hash.size = newSize
}

// External iterator functions

func (iter *iterOpenHash[K, V]) findList() {
	for iter.HasNext() {
		list := iter.hash.table[iter.currentPos]
		if !list.IsEmpty() {
			iter.current = list.Iterator()
			return
		}
		iter.currentPos++
	}
}

// Hashing functions

func convertToPosition[K any](key K, size int) int {
	keyBytes := convertToBytes(key)
	return hashingFNV(keyBytes, size)
}

func convertToBytes[K any](key K) []byte {
	return fmt.Appendf(nil, "%v", key)
}

func hashingFNV(key []byte, size int) int {
	var h uint64 = 14695981039346656037
	for _, c := range key {
		h *= 1099511628211
		h ^= uint64(c)
	}
	return int(h % uint64(size))
}
