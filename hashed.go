package multimap

import (
	"iter"
	"unsafe"

	set3 "github.com/TomTonic/Set3"

	"github.com/TomTonic/multimap/internal/art"
	"github.com/TomTonic/multimap/internal/vset"
)

// Hashed is a multimap backed by a Go map. Point operations cost one hash
// lookup, which beats Ordered on long string keys. Every range query scans
// all keys, though, and iteration order is unspecified. Like every Go map it
// keeps its memory after deletions. Use it when range queries are rare.
//
// Hashed is not synchronized: concurrent reads are safe, but a write must not
// run concurrently with any other access. Wrap it with Synchronized for
// concurrent use.
type Hashed[T comparable] struct {
	// Values are held by pointer so that updating a set is a single lookup:
	// a Go map cannot update a value in place, and reassigning it would copy
	// the key again.
	m map[string]*vset.Set[T]
}

var _ MultiMap[int] = (*Hashed[int])(nil)

// NewHashed returns an empty, unsynchronized Hashed multimap. Use it directly
// when access is already serialized, or wrap it with Synchronized.
func NewHashed[T comparable]() *Hashed[T] { return &Hashed[T]{m: map[string]*vset.Set[T]{}} }

// AddValue adds value to the set at key, cloning key if it is new.
func (m *Hashed[T]) AddValue(key Key, value T) {
	s := m.m[string(key)] // no allocation: the compiler views key in place
	if s == nil {
		s = &vset.Set[T]{}
		m.m[string(key)] = s // copies the key
	}
	s.Add(value)
}

// RemoveValue removes value from the set at key, and the key once its set is
// empty.
func (m *Hashed[T]) RemoveValue(key Key, value T) {
	if s := m.m[string(key)]; s != nil && s.Remove(value) && s.Len() == 0 {
		delete(m.m, string(key))
	}
}

// RemoveKey removes key and all its values.
func (m *Hashed[T]) RemoveKey(key Key) { delete(m.m, string(key)) }

// Clear removes all keys and values.
func (m *Hashed[T]) Clear() { clear(m.m) }

// ContainsKey reports whether key is present.
func (m *Hashed[T]) ContainsKey(key Key) bool { return m.m[string(key)] != nil }

// NumberOfKeys returns the number of keys.
func (m *Hashed[T]) NumberOfKeys() uint64 { return uint64(len(m.m)) }

// ValuesFor returns a copy of the values of key; empty if key is absent.
func (m *Hashed[T]) ValuesFor(key Key) *set3.Set3[T] { return collect(m.ValuesForSeq(key)) }

// ValuesForSeq iterates over the values of key.
func (m *Hashed[T]) ValuesForSeq(key Key) iter.Seq[T] {
	return func(yield func(T) bool) {
		if s := m.m[string(key)]; s != nil {
			s.Each(yield)
		}
	}
}

// rangeSeq iterates over the values of all keys within b. It scans every key.
func (m *Hashed[T]) rangeSeq(b art.Bounds) iter.Seq[T] {
	return func(yield func(T) bool) {
		for k, s := range m.m {
			if b.Contains(view(k)) && !s.Each(yield) {
				return
			}
		}
	}
}

// ValuesBetweenInclusive returns a copy of the values of all keys in [from, to].
func (m *Hashed[T]) ValuesBetweenInclusive(from, to Key) *set3.Set3[T] {
	return collect(m.ValuesBetweenInclusiveSeq(from, to))
}

// ValuesBetweenInclusiveSeq iterates over the values of all keys in [from, to].
func (m *Hashed[T]) ValuesBetweenInclusiveSeq(from, to Key) iter.Seq[T] {
	return m.rangeSeq(between(from, to, true))
}

// ValuesBetweenExclusive returns a copy of the values of all keys in (from, to).
func (m *Hashed[T]) ValuesBetweenExclusive(from, to Key) *set3.Set3[T] {
	return collect(m.ValuesBetweenExclusiveSeq(from, to))
}

// ValuesBetweenExclusiveSeq iterates over the values of all keys in (from, to).
func (m *Hashed[T]) ValuesBetweenExclusiveSeq(from, to Key) iter.Seq[T] {
	return m.rangeSeq(between(from, to, false))
}

// ValuesFromInclusive returns a copy of the values of all keys >= from.
func (m *Hashed[T]) ValuesFromInclusive(from Key) *set3.Set3[T] {
	return collect(m.ValuesFromInclusiveSeq(from))
}

// ValuesFromInclusiveSeq iterates over the values of all keys >= from.
func (m *Hashed[T]) ValuesFromInclusiveSeq(from Key) iter.Seq[T] {
	return m.rangeSeq(fromBound(from, true))
}

// ValuesFromExclusive returns a copy of the values of all keys > from.
func (m *Hashed[T]) ValuesFromExclusive(from Key) *set3.Set3[T] {
	return collect(m.ValuesFromExclusiveSeq(from))
}

// ValuesFromExclusiveSeq iterates over the values of all keys > from.
func (m *Hashed[T]) ValuesFromExclusiveSeq(from Key) iter.Seq[T] {
	return m.rangeSeq(fromBound(from, false))
}

// ValuesToInclusive returns a copy of the values of all keys <= to.
func (m *Hashed[T]) ValuesToInclusive(to Key) *set3.Set3[T] {
	return collect(m.ValuesToInclusiveSeq(to))
}

// ValuesToInclusiveSeq iterates over the values of all keys <= to.
func (m *Hashed[T]) ValuesToInclusiveSeq(to Key) iter.Seq[T] {
	return m.rangeSeq(toBound(to, true))
}

// ValuesToExclusive returns a copy of the values of all keys < to.
func (m *Hashed[T]) ValuesToExclusive(to Key) *set3.Set3[T] {
	return collect(m.ValuesToExclusiveSeq(to))
}

// ValuesToExclusiveSeq iterates over the values of all keys < to.
func (m *Hashed[T]) ValuesToExclusiveSeq(to Key) iter.Seq[T] {
	return m.rangeSeq(toBound(to, false))
}

// AllValues returns a copy of all values.
func (m *Hashed[T]) AllValues() *set3.Set3[T] { return collect(m.AllValuesSeq()) }

// AllValuesSeq iterates over all values, key by key in unspecified order.
func (m *Hashed[T]) AllValuesSeq() iter.Seq[T] { return m.rangeSeq(art.Bounds{}) }

// AllKeys returns clones of all keys in unspecified order.
func (m *Hashed[T]) AllKeys() []Key {
	keys := make([]Key, 0, len(m.m))
	for k := range m.m {
		keys = append(keys, Key(k))
	}
	return keys
}

// AllKeysSeq iterates over all keys in unspecified order. The yielded Keys
// must not be retained after the loop step (clone them to keep them).
func (m *Hashed[T]) AllKeysSeq() iter.Seq[Key] {
	return func(yield func(Key) bool) {
		for k := range m.m {
			if !yield(view(k)) {
				return
			}
		}
	}
}

// view returns the bytes of s without copying. The result must not be
// modified: map keys are immutable strings.
func view(s string) []byte { return unsafe.Slice(unsafe.StringData(s), len(s)) }
