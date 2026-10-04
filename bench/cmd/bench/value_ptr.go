//go:build ptrvals && !strvals

package main

// V is *rec in a bench built with the tag ptrvals: values that are pointers to
// records of the caller's, as an application indexes its objects by key. The
// candidates hold the pointers (8 bytes); the records are the caller's and
// cost every candidate the same, so memory figures do not count them.
type V = *rec

// rec is a record of 16 bytes, an object of its own in the Go heap.
type rec struct{ id, aux uint64 }

// valueTag marks results of this build: they are reported as natural-ptr and
// single-value-ptr.
const valueTag = "-ptr"

// toVs returns the values for the value numbers u: one record for each
// different number, allocated one by one in the order the numbers first
// appear, so that equal numbers are the same pointer and the records lie
// about as scattered in the heap as an application's would. The names that
// some corpora give the numbers are not used.
func toVs(u []uint64, _ []string) []V {
	out := make([]V, len(u))
	seen := make(map[uint64]V, len(u))
	for i, x := range u {
		r, ok := seen[x]
		if !ok {
			r = &rec{id: x, aux: x * 0x9e3779b97f4a7c15}
			seen[x] = r
		}
		out[i] = r
	}
	return out
}

// weigh returns what the timed loops add up per value: the record's id. It
// reads the record, as a caller that uses a value does, and so pays the
// cache miss of the pointer.
func weigh(v V) uint64 { return v.id }

// checkWeigh is weigh for the checks that candidates hold the same values:
// the id and the other field, so that a wrong pointer shows.
func checkWeigh(v V) uint64 { return v.id ^ v.aux }

// valueBytes returns the bytes of the string values vals; none for pointers.
func valueBytes(_ []V) int { return 0 }
