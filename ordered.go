package multimap

import (
	"iter"

	set3 "github.com/TomTonic/Set3"

	"github.com/TomTonic/multimap/internal/art"
)

// Ordered is a multimap backed by an adaptive radix tree. Point operations
// descend one node per distinguishing key byte; range queries visit only the
// keys in range; all iteration is in ascending key order. The zero value is
// an empty multimap.
//
// Ordered is not synchronized: concurrent reads are safe, but a write must
// not run concurrently with any other access. Wrap it with Synchronized for
// concurrent use.
type Ordered[T comparable] struct {
	m art.Map[T]
}

var _ MultiMap[int] = (*Ordered[int])(nil)

// NewOrdered returns an empty, unsynchronized Ordered multimap. Use it
// directly when access is already serialized, or wrap it with Synchronized.
func NewOrdered[T comparable]() *Ordered[T] { return &Ordered[T]{} }

// AddValue adds value to the set at key, cloning key if it is new.
func (m *Ordered[T]) AddValue(key Key, value T) { m.m.Add(key, value) }

// RemoveValue removes value from the set at key, and the key once its set is
// empty.
func (m *Ordered[T]) RemoveValue(key Key, value T) { m.m.Remove(key, value) }

// RemoveKey removes key and all its values.
func (m *Ordered[T]) RemoveKey(key Key) { m.m.RemoveKey(key) }

// Clear removes all keys and values.
func (m *Ordered[T]) Clear() { m.m.Clear() }

// ContainsKey reports whether key is present.
func (m *Ordered[T]) ContainsKey(key Key) bool { return m.m.Values(key).Found() }

// NumberOfKeys returns the number of keys.
func (m *Ordered[T]) NumberOfKeys() uint64 { return uint64(m.m.Len()) }

// ValuesFor returns a copy of the values of key; empty if key is absent.
func (m *Ordered[T]) ValuesFor(key Key) *set3.Set3[T] { return collect(m.ValuesForSeq(key)) }

// ValuesForSeq iterates over the values of key.
func (m *Ordered[T]) ValuesForSeq(key Key) iter.Seq[T] {
	return func(yield func(T) bool) { m.m.Values(key).Each(yield) }
}

// rangeSeq iterates over the values of all keys within b, key by key in
// ascending key order.
func (m *Ordered[T]) rangeSeq(b art.Bounds) iter.Seq[T] {
	return func(yield func(T) bool) { m.m.RangeValues(&b, yield) }
}

// ValuesBetweenInclusive returns a copy of the values of all keys in [from, to].
func (m *Ordered[T]) ValuesBetweenInclusive(from, to Key) *set3.Set3[T] {
	return collect(m.ValuesBetweenInclusiveSeq(from, to))
}

// ValuesBetweenInclusiveSeq iterates over the values of all keys in [from, to].
func (m *Ordered[T]) ValuesBetweenInclusiveSeq(from, to Key) iter.Seq[T] {
	return m.rangeSeq(between(from, to, true))
}

// ValuesBetweenExclusive returns a copy of the values of all keys in (from, to).
func (m *Ordered[T]) ValuesBetweenExclusive(from, to Key) *set3.Set3[T] {
	return collect(m.ValuesBetweenExclusiveSeq(from, to))
}

// ValuesBetweenExclusiveSeq iterates over the values of all keys in (from, to).
func (m *Ordered[T]) ValuesBetweenExclusiveSeq(from, to Key) iter.Seq[T] {
	return m.rangeSeq(between(from, to, false))
}

// ValuesFromInclusive returns a copy of the values of all keys >= from.
func (m *Ordered[T]) ValuesFromInclusive(from Key) *set3.Set3[T] {
	return collect(m.ValuesFromInclusiveSeq(from))
}

// ValuesFromInclusiveSeq iterates over the values of all keys >= from.
func (m *Ordered[T]) ValuesFromInclusiveSeq(from Key) iter.Seq[T] {
	return m.rangeSeq(fromBound(from, true))
}

// ValuesFromExclusive returns a copy of the values of all keys > from.
func (m *Ordered[T]) ValuesFromExclusive(from Key) *set3.Set3[T] {
	return collect(m.ValuesFromExclusiveSeq(from))
}

// ValuesFromExclusiveSeq iterates over the values of all keys > from.
func (m *Ordered[T]) ValuesFromExclusiveSeq(from Key) iter.Seq[T] {
	return m.rangeSeq(fromBound(from, false))
}

// ValuesToInclusive returns a copy of the values of all keys <= to.
func (m *Ordered[T]) ValuesToInclusive(to Key) *set3.Set3[T] {
	return collect(m.ValuesToInclusiveSeq(to))
}

// ValuesToInclusiveSeq iterates over the values of all keys <= to.
func (m *Ordered[T]) ValuesToInclusiveSeq(to Key) iter.Seq[T] {
	return m.rangeSeq(toBound(to, true))
}

// ValuesToExclusive returns a copy of the values of all keys < to.
func (m *Ordered[T]) ValuesToExclusive(to Key) *set3.Set3[T] {
	return collect(m.ValuesToExclusiveSeq(to))
}

// ValuesToExclusiveSeq iterates over the values of all keys < to.
func (m *Ordered[T]) ValuesToExclusiveSeq(to Key) iter.Seq[T] {
	return m.rangeSeq(toBound(to, false))
}

// AllValues returns a copy of all values.
func (m *Ordered[T]) AllValues() *set3.Set3[T] { return collect(m.AllValuesSeq()) }

// AllValuesSeq iterates over all values, key by key in ascending key order.
func (m *Ordered[T]) AllValuesSeq() iter.Seq[T] { return m.rangeSeq(art.Bounds{}) }

// AllKeys returns clones of all keys in ascending order.
func (m *Ordered[T]) AllKeys() []Key {
	keys := make([]Key, 0, m.m.Len())
	for k := range m.AllKeysSeq() {
		keys = append(keys, k.Clone())
	}
	return keys
}

// AllKeysSeq iterates over all keys in ascending order. The yielded Keys are
// views into the multimap and must not be modified or retained.
func (m *Ordered[T]) AllKeysSeq() iter.Seq[Key] {
	return func(yield func(Key) bool) {
		m.m.Range(&art.Bounds{}, func(k []byte, _ art.View[T]) bool { return yield(k) })
	}
}

func between(from, to Key, incl bool) art.Bounds {
	return art.Bounds{From: from, To: to, HasFrom: true, HasTo: true, FromIncl: incl, ToIncl: incl}
}

func fromBound(from Key, incl bool) art.Bounds {
	return art.Bounds{From: from, HasFrom: true, FromIncl: incl}
}

func toBound(to Key, incl bool) art.Bounds {
	return art.Bounds{To: to, HasTo: true, ToIncl: incl}
}
