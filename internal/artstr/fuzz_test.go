package artstr

import (
	"slices"
	"strings"
	"testing"
)

// FuzzOperations drives the page tree with operations decoded from bytes and
// compares it with a model after every one.
//
// A user adds and removes values of keys of every shape, in any order; whatever
// the sequence, the tree must hold exactly the keys and values of the model, in
// order, in every mode (immutable pages or not, several values per key in the
// pages or not). The first two bytes choose the mode, then three bytes make one
// operation: its kind, a key out of a small set with shared prefixes and one
// that is too long for a page, and a value of several lengths.
func FuzzOperations(f *testing.F) {
	f.Add([]byte{0, 0, 1, 2, 3, 1, 2, 4, 0, 5, 6, 2, 2, 3, 3, 2, 3})
	f.Add([]byte{1, 1, 0, 0, 0, 0, 0, 1, 0, 0, 2, 0, 0, 3, 0, 0, 4, 2, 0, 0})
	f.Fuzz(func(t *testing.T, data []byte) {
		if len(data) < 2 {
			return
		}
		m := Map[string]{Pairs: data[0]&1 == 1, ZeroCopy: data[0]&2 == 2}
		want := model{}
		keyOf := func(b byte) string {
			switch b % 8 {
			case 0:
				return ""
			case 1:
				return "a"
			case 2:
				return "ab"
			case 3:
				return "abc"
			case 4:
				return "abd" + strings.Repeat("x", int(b)/8)
			case 5:
				return "b"
			case 6:
				return strings.Repeat("q", 250+int(b)/8)
			}
			return string([]byte{'a', b, b})
		}
		valOf := func(b byte) string { return strings.Repeat("v", int(b)%5*int(b)%7) + string('0'+b%10) }
		for ops := data[2:]; len(ops) >= 3; ops = ops[3:] {
			k, v := keyOf(ops[1]), valOf(ops[2])
			switch ops[0] % 4 {
			case 0, 1:
				m.Add([]byte(k), v)
				want.add(k, v)
			case 2:
				m.Remove([]byte(k), v)
				if want[k][v] {
					want.remove(k, v)
				}
			case 3:
				m.RemoveKey([]byte(k))
				delete(want, k)
			}
			if m.Len() != len(want) {
				t.Fatalf("Len = %d, want %d", m.Len(), len(want))
			}
		}
		for k, vs := range want {
			got := map[string]bool{}
			m.Each([]byte(k), func(v string) bool { got[v] = true; return true })
			if !maps(got, vs) {
				t.Fatalf("key %q has %v, want %v", k, got, vs)
			}
		}
		var keys []string
		m.Range(&Bounds{}, func(k []byte) bool { keys = append(keys, string(k)); return true })
		if !slices.Equal(keys, want.sorted()) {
			t.Fatalf("Range = %q, want %q", keys, want.sorted())
		}
	})
}
