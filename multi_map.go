package multimap

import (
	"sync"

	set3 "github.com/TomTonic/Set3"
)

// MultiMap is a simple multi-map from string keys to a Set3 of values.
//
// NOTE: For now methods are API-only (placeholders). We'll fill implementations
// after the API is reviewed.
type MultiMap[T comparable] struct {
	mu   sync.RWMutex
	data []kvp[T]
}

type kvp[T comparable] struct {
	key Key
	val *set3.Set3[T]
}

// New creates and returns an initialized MultiMap.
func New[T comparable]() *MultiMap[T] {
	return &MultiMap[T]{
		// start with zero length, reserve capacity
		data: make([]kvp[T], 0, 20),
	}
}

// PutValue adds value v to the set of values at key. If the key does not exist it will be created.
func (m *MultiMap[T]) PutValue(key Key, v T) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i := range m.data {
		if m.data[i].key.Equal(key) {
			if m.data[i].val == nil {
				m.data[i].val = set3.Empty[T]()
			}
			m.data[i].val.Add(v)
			return
		}
	}
	newTuple := kvp[T]{
		key: key.Clone(), // store independent copy
		val: set3.Empty[T](),
	}
	newTuple.val.Add(v)
	m.data = append(m.data, newTuple)
}

// RemoveValue removes value v from the set of values at key. If the set becomes empty the key may be removed.
func (m *MultiMap[T]) RemoveValue(key Key, v T) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i := range m.data {
		if m.data[i].key.Equal(key) {
			if m.data[i].val != nil {
				m.data[i].val.Remove(v)
				// optional: if set becomes empty, remove the kvp
				// if m.data[i].val.IsEmpty() {
				//     m.data[i] = m.data[len(m.data)-1]
				//     m.data = m.data[:len(m.data)-1]
				// }
			}
			return
		}
	}
}

// ContainsKey checks whether the multi-map contains the specified key.
func (m *MultiMap[T]) ContainsKey(key Key) bool {
	m.mu.RLock()
	defer m.mu.RUnlock()
	for i := range m.data {
		if m.data[i].key.Equal(key) {
			return true
		}
	}
	return false
}

// RemoveKey removes the key and its associated set of values from the multi-map.
func (m *MultiMap[T]) RemoveKey(key Key) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i := range m.data {
		if m.data[i].key.Equal(key) {
			// Remove the key-value pair by swapping with the last and truncating
			m.data[i] = m.data[len(m.data)-1]
			m.data = m.data[:len(m.data)-1]
			return
		}
	}
}

// Get returns the set of values associated with key. If the key does not exist
// or if it exists but no values are stored for this key, it returns an empty set.
func (m *MultiMap[T]) GetValuesFor(key Key) *set3.Set3[T] {
	m.mu.RLock()
	defer m.mu.RUnlock()
	for i := range m.data {
		if m.data[i].key.Equal(key) {
			if m.data[i].val != nil {
				return m.data[i].val.Clone()
			}
			return set3.EmptyWithCapacity[T](0)
		}
	}
	return set3.EmptyWithCapacity[T](0)
}

// GetAllValues returns a set with all values currently stored in the multi-map.
func (m *MultiMap[T]) GetAllValues() *set3.Set3[T] {
	m.mu.RLock()
	defer m.mu.RUnlock()
	result := set3.Empty[T]()
	for i := range m.data {
		if m.data[i].val != nil {
			result.AddAll(m.data[i].val)
		}
	}
	return result
}

// GetValuesBetweenInclusive returns a set with all values whose keys are between from and to, including values stored for from and to.
func (m *MultiMap[T]) GetValuesBetweenInclusive(from, to Key) *set3.Set3[T] {
	m.mu.RLock()
	defer m.mu.RUnlock()
	result := set3.Empty[T]()
	for _, kv := range m.data {
		if (kv.key.LessThan(to) || kv.key.Equal(to)) && (from.LessThan(kv.key) || from.Equal(kv.key)) {
			if kv.val != nil {
				result.AddAll(kv.val)
			}
		}
	}
	return result
}

// GetValuesBetweenExclusive returns a set with all values whose keys are between from and to, excluding values stored for from and to.
func (m *MultiMap[T]) GetValuesBetweenExclusive(from, to Key) *set3.Set3[T] {
	m.mu.RLock()
	defer m.mu.RUnlock()
	result := set3.Empty[T]()
	for _, kv := range m.data {
		if kv.key.LessThan(to) && from.LessThan(kv.key) {
			if kv.val != nil {
				result.AddAll(kv.val)
			}
		}
	}
	return result
}

// GetValuesFromInclusive returns a set with all values whose keys are greater than or equal to from.
func (m *MultiMap[T]) GetValuesFromInclusive(from Key) *set3.Set3[T] {
	m.mu.RLock()
	defer m.mu.RUnlock()
	result := set3.Empty[T]()
	for _, kv := range m.data {
		if from.LessThan(kv.key) || from.Equal(kv.key) {
			if kv.val != nil {
				result.AddAll(kv.val)
			}
		}
	}
	return result
}

// GetValuesToInclusive returns a set with all values whose keys are less than or equal to to.
func (m *MultiMap[T]) GetValuesToInclusive(to Key) *set3.Set3[T] {
	m.mu.RLock()
	defer m.mu.RUnlock()
	result := set3.Empty[T]()
	for _, kv := range m.data {
		if kv.key.LessThan(to) || kv.key.Equal(to) {
			if kv.val != nil {
				result.AddAll(kv.val)
			}
		}
	}
	return result
}

// GetValuesFromExclusive returns a set with all values whose keys are actually greater than from.
func (m *MultiMap[T]) GetValuesFromExclusive(from Key) *set3.Set3[T] {
	m.mu.RLock()
	defer m.mu.RUnlock()
	result := set3.Empty[T]()
	for _, kv := range m.data {
		if from.LessThan(kv.key) {
			if kv.val != nil {
				result.AddAll(kv.val)
			}
		}
	}
	return result
}

// GetValuesToExclusive returns a set with all values whose keys are actually less than to.
func (m *MultiMap[T]) GetValuesToExclusive(to Key) *set3.Set3[T] {
	m.mu.RLock()
	defer m.mu.RUnlock()
	result := set3.Empty[T]()
	for _, kv := range m.data {
		if kv.key.LessThan(to) {
			if kv.val != nil {
				result.AddAll(kv.val)
			}
		}
	}
	return result
}

// Size returns the number of keys currently stored in the map.
func (m *MultiMap[T]) Size() uint64 {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return uint64(len(m.data))
}

// Keys returns a slice with all keys currently stored in the map.
func (m *MultiMap[T]) Keys() []Key {
	m.mu.RLock()
	defer m.mu.RUnlock()
	result := make([]Key, 0, len(m.data))
	for i := range m.data {
		result = append(result, m.data[i].key.Clone())
	}
	return result
}

// Clear removes all keys and values from the multi-map.
func (m *MultiMap[T]) Clear() {
	m.mu.Lock()
	defer m.mu.Unlock()
	// reset to zero length, reserve capacity
	m.data = make([]kvp[T], 0, 20)
}

// Example usage (for documentation):
//
//  mm := multimap.New()
//  mm.Put("a", 1)
//  s := mm.Get("a") // returns *set3.Set3
//
