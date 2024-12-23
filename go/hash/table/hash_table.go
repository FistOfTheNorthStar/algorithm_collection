package hashtable

const SIZE = 121

// Entry represents a key-value pair in the hash table
type Entry[K comparable, V any] struct {
	key   K
	value V
	next  *Entry[K, V]
}

// HashTable implements a generic hash table with separate chaining
type HashTable[K comparable, V any] struct {
	entries [SIZE]*Entry[K, V]
}

// NewHashTable creates a new hash table instance
func NewHashTable[K comparable, V any]() *HashTable[K, V] {
	return &HashTable[K, V]{
		entries: [SIZE]*Entry[K, V]{},
	}
}

// hashTheKey implements the custom hashing function for strings
func (ht *HashTable[K, V]) hashTheKey(key K) int {
	// Type assert key to string
	str, ok := any(key).(string)
	if !ok {
		return 0
	}

	product := 1
	for _, c := range str {
		switch c {
		case 'A':
			product *= 1
		case 'B':
			product *= 2
		case 'C':
			product *= 3
		case 'D':
			product *= 4
		case 'E':
			product *= 5
		}
	}

	return product
}

// Put adds a key-value pair to the hash table
func (ht *HashTable[K, V]) Put(key K, value V) {
	hash := ht.hashTheKey(key)
	newEntry := &Entry[K, V]{
		key:   key,
		value: value,
	}

	if ht.entries[hash] == nil {
		ht.entries[hash] = newEntry
		return
	}

	// Handle collision with chaining
	currentEntry := ht.entries[hash]
	for currentEntry.next != nil {
		currentEntry = currentEntry.next
	}
	currentEntry.next = newEntry
}

// Get retrieves a value by key from the hash table
func (ht *HashTable[K, V]) Get(key K) (V, bool) {
	hash := ht.hashTheKey(key)
	currentEntry := ht.entries[hash]

	for currentEntry != nil {
		if currentEntry.key == key {
			return currentEntry.value, true
		}
		currentEntry = currentEntry.next
	}

	var zero V
	return zero, false
}
