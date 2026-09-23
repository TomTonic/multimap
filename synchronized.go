package multimap

import (
	"iter"
	"sync"

	set3 "github.com/TomTonic/Set3"
)

// synchronized guards a MultiMap with a sync.RWMutex: reads share the lock,
// writes hold it exclusively.
type synchronized[T comparable] struct {
	mu sync.RWMutex
	m  MultiMap[T]
}

// Synchronized returns a MultiMap that forwards every call to m under a
// sync.RWMutex, making all methods safe for concurrent use. Reads run in
// parallel; a write waits for all running reads and blocks new ones.
//
// The Seq methods hold the read lock for the whole loop over the returned
// iterator. The loop body must not call any method of the same multimap:
// a write deadlocks at once, and a read deadlocks if a writer is waiting.
// A long loop delays all writers. Use the set-returning methods when the
// loop body needs the multimap or runs long; they release the lock before
// returning.
//
// m must not be used directly while the wrapper is in use.
func Synchronized[T comparable](m MultiMap[T]) MultiMap[T] { return &synchronized[T]{m: m} }

// read runs f under the read lock.
func read[T comparable, R any](s *synchronized[T], f func(MultiMap[T]) R) R {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return f(s.m)
}

// readSeq wraps an iterator so that the whole loop runs under the read lock.
func readSeq[T comparable, V any](s *synchronized[T], seq func(MultiMap[T]) iter.Seq[V]) iter.Seq[V] {
	return func(yield func(V) bool) {
		s.mu.RLock()
		defer s.mu.RUnlock()
		for v := range seq(s.m) {
			if !yield(v) {
				return
			}
		}
	}
}

func (s *synchronized[T]) write(f func(MultiMap[T])) {
	s.mu.Lock()
	defer s.mu.Unlock()
	f(s.m)
}

func (s *synchronized[T]) AddValue(key Key, value T) {
	s.write(func(m MultiMap[T]) { m.AddValue(key, value) })
}

func (s *synchronized[T]) RemoveValue(key Key, value T) {
	s.write(func(m MultiMap[T]) { m.RemoveValue(key, value) })
}

func (s *synchronized[T]) RemoveKey(key Key) { s.write(func(m MultiMap[T]) { m.RemoveKey(key) }) }

func (s *synchronized[T]) Clear() { s.write(func(m MultiMap[T]) { m.Clear() }) }

func (s *synchronized[T]) ContainsKey(key Key) bool {
	return read(s, func(m MultiMap[T]) bool { return m.ContainsKey(key) })
}

func (s *synchronized[T]) NumberOfKeys() uint64 {
	return read(s, func(m MultiMap[T]) uint64 { return m.NumberOfKeys() })
}

func (s *synchronized[T]) ValuesFor(key Key) *set3.Set3[T] {
	return read(s, func(m MultiMap[T]) *set3.Set3[T] { return m.ValuesFor(key) })
}

func (s *synchronized[T]) ValuesForSeq(key Key) iter.Seq[T] {
	return readSeq(s, func(m MultiMap[T]) iter.Seq[T] { return m.ValuesForSeq(key) })
}

func (s *synchronized[T]) ValuesBetweenInclusive(from, to Key) *set3.Set3[T] {
	return read(s, func(m MultiMap[T]) *set3.Set3[T] { return m.ValuesBetweenInclusive(from, to) })
}

func (s *synchronized[T]) ValuesBetweenInclusiveSeq(from, to Key) iter.Seq[T] {
	return readSeq(s, func(m MultiMap[T]) iter.Seq[T] { return m.ValuesBetweenInclusiveSeq(from, to) })
}

func (s *synchronized[T]) ValuesBetweenExclusive(from, to Key) *set3.Set3[T] {
	return read(s, func(m MultiMap[T]) *set3.Set3[T] { return m.ValuesBetweenExclusive(from, to) })
}

func (s *synchronized[T]) ValuesBetweenExclusiveSeq(from, to Key) iter.Seq[T] {
	return readSeq(s, func(m MultiMap[T]) iter.Seq[T] { return m.ValuesBetweenExclusiveSeq(from, to) })
}

func (s *synchronized[T]) ValuesFromInclusive(from Key) *set3.Set3[T] {
	return read(s, func(m MultiMap[T]) *set3.Set3[T] { return m.ValuesFromInclusive(from) })
}

func (s *synchronized[T]) ValuesFromInclusiveSeq(from Key) iter.Seq[T] {
	return readSeq(s, func(m MultiMap[T]) iter.Seq[T] { return m.ValuesFromInclusiveSeq(from) })
}

func (s *synchronized[T]) ValuesFromExclusive(from Key) *set3.Set3[T] {
	return read(s, func(m MultiMap[T]) *set3.Set3[T] { return m.ValuesFromExclusive(from) })
}

func (s *synchronized[T]) ValuesFromExclusiveSeq(from Key) iter.Seq[T] {
	return readSeq(s, func(m MultiMap[T]) iter.Seq[T] { return m.ValuesFromExclusiveSeq(from) })
}

func (s *synchronized[T]) ValuesToInclusive(to Key) *set3.Set3[T] {
	return read(s, func(m MultiMap[T]) *set3.Set3[T] { return m.ValuesToInclusive(to) })
}

func (s *synchronized[T]) ValuesToInclusiveSeq(to Key) iter.Seq[T] {
	return readSeq(s, func(m MultiMap[T]) iter.Seq[T] { return m.ValuesToInclusiveSeq(to) })
}

func (s *synchronized[T]) ValuesToExclusive(to Key) *set3.Set3[T] {
	return read(s, func(m MultiMap[T]) *set3.Set3[T] { return m.ValuesToExclusive(to) })
}

func (s *synchronized[T]) ValuesToExclusiveSeq(to Key) iter.Seq[T] {
	return readSeq(s, func(m MultiMap[T]) iter.Seq[T] { return m.ValuesToExclusiveSeq(to) })
}

func (s *synchronized[T]) AllValues() *set3.Set3[T] {
	return read(s, func(m MultiMap[T]) *set3.Set3[T] { return m.AllValues() })
}

func (s *synchronized[T]) AllValuesSeq() iter.Seq[T] {
	return readSeq(s, func(m MultiMap[T]) iter.Seq[T] { return m.AllValuesSeq() })
}

func (s *synchronized[T]) AllKeys() []Key {
	return read(s, func(m MultiMap[T]) []Key { return m.AllKeys() })
}

func (s *synchronized[T]) AllKeysSeq() iter.Seq[Key] {
	return readSeq(s, func(m MultiMap[T]) iter.Seq[Key] { return m.AllKeysSeq() })
}
