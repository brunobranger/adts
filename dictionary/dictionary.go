package dictionary

type Dictionary[K any, V any] interface {

	// Save stores the key-value pair in the Dictionary. If the key already exists, it updates the associated value
	Save(key K, value V)

	// Belongs determines whether a key is already in the dictionary or not
	Belongs(key K) bool

	// Get returns the value associated with a key. If the key does not exist, it must panic with the message
	// 'The key does not belong to the dictionary'
	Get(key K) V

	// Delete removes the specified key from the Dictionary, returning the value that was associated with it.
	// If the key does not belong to the dictionary, it must panic with the message 'The key does not belong to the dictionary'
	Delete(key K) V

	// Count returns the number of elements in the dictionary
	Count() int

	// Iterate internally iterates through the dictionary, applying the passed function to all elements
	Iterate(func(key K, value V) bool)

	// Iterator returns a DictionaryIterator for this Dictionary
	Iterator() DictionaryIterator[K, V]
}

type DictionaryIterator[K any, V any] interface {

	// HasNext returns whether there are more elements to see. That is, if the current position of the iterator
	// has an element.
	HasNext() bool

	// Current returns the key and value of the current element where the iterator is positioned.
	// If there is no next element (HasNext returns false), it must panic with the message 'The iterator has finished iterating'
	Current() (K, V)

	// Next advances the iterator to the next element in the dictionary if HasNext returns true.
	// If there is no next element, it must panic with the message 'The iterator has finished iterating'
	Next()
}
