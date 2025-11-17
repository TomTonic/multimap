package multimap

import (
	"sync"

	set3 "github.com/TomTonic/Set3"
)

// arrayBasedMultiMap is the default implementation using a simple slice of key/value pairs.
// It preserves previous behavior while satisfying the MultiMap interface.
type arrayBasedMultiMap[T comparable] struct {
	mu   sync.RWMutex
	data []kvp[T]
}

type kvp[T comparable] struct {
	key Key
	val *set3.Set3[T]
}

func newArrayBased[T comparable]() *arrayBasedMultiMap[T] {
	return &arrayBasedMultiMap[T]{
		data: make([]kvp[T], 0, 20),
	}
}

func (m *arrayBasedMultiMap[T]) AddValue(key Key, v T) {
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
		key: key.Clone(),
		val: set3.Empty[T](),
	}
	newTuple.val.Add(v)
	m.data = append(m.data, newTuple)
}

func (m *arrayBasedMultiMap[T]) RemoveValue(key Key, v T) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i := range m.data {
		if m.data[i].key.Equal(key) {
			if m.data[i].val != nil {
				m.data[i].val.Remove(v)
			}
			return
		}
	}
}

func (m *arrayBasedMultiMap[T]) ContainsKey(key Key) bool {
	m.mu.RLock()
	defer m.mu.RUnlock()
	for i := range m.data {
		if m.data[i].key.Equal(key) {
			return true
		}
	}
	return false
}

func (m *arrayBasedMultiMap[T]) RemoveKey(key Key) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i := range m.data {
		if m.data[i].key.Equal(key) {
			m.data[i] = m.data[len(m.data)-1]
			m.data = m.data[:len(m.data)-1]
			return
		}
	}
}

func (m *arrayBasedMultiMap[T]) ValuesFor(key Key) *set3.Set3[T] {
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

func (m *arrayBasedMultiMap[T]) AllValues() *set3.Set3[T] {
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

func (m *arrayBasedMultiMap[T]) ValuesBetweenInclusive(from, to Key) *set3.Set3[T] {
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

func (m *arrayBasedMultiMap[T]) ValuesBetweenExclusive(from, to Key) *set3.Set3[T] {
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

func (m *arrayBasedMultiMap[T]) ValuesFromInclusive(from Key) *set3.Set3[T] {
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

func (m *arrayBasedMultiMap[T]) ValuesToInclusive(to Key) *set3.Set3[T] {
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

func (m *arrayBasedMultiMap[T]) ValuesFromExclusive(from Key) *set3.Set3[T] {
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

func (m *arrayBasedMultiMap[T]) ValuesToExclusive(to Key) *set3.Set3[T] {
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

func (m *arrayBasedMultiMap[T]) NumberOfKeys() uint64 {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return uint64(len(m.data))
}

func (m *arrayBasedMultiMap[T]) AllKeys() []Key {
	m.mu.RLock()
	defer m.mu.RUnlock()
	result := make([]Key, 0, len(m.data))
	for i := range m.data {
		result = append(result, m.data[i].key.Clone())
	}
	return result
}

func (m *arrayBasedMultiMap[T]) Clear() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.data = make([]kvp[T], 0, 20)
}
