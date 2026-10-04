package art

import (
	"unsafe"

	"github.com/TomTonic/multimap/internal/skpage"
)

// In a map of strings (Map.flat == 3) a key lives in a single-key page
// (internal/skpage, docs/redesign/step3-skmv-design.md): the remainder and all
// the values as bytes, in one object without pointers of 32 to 512 bytes. The
// page starts like a leaf (kind, remainder length, number of values, length of
// the whole key), so everything the tree does with the key of a leaf works on it.
// Its kinds are kSet+1 to kSet+6, the kinds that flat and typed leaves have in
// other maps; kSet is still the set leaf, which takes a key whose values do not
// fit a page (the value overflow) or whose remainder is longer than a page holds.

func init() { skpage.KindBase = uint8(kSet) + 1 }

// stringType reports whether T is string, the value type that takes single-key
// pages.
func stringType[T comparable]() bool {
	var z T
	_, ok := any(z).(string)
	return ok
}

func asSK(l *leafHead) *skpage.Page   { return (*skpage.Page)(unsafe.Pointer(l)) }
func skLeaf(p *skpage.Page) *leafHead { return (*leafHead)(unsafe.Pointer(p)) }

// view returns the bytes of s without copying them. The page copies what it
// stores.
func view(s string) []byte { return unsafe.Slice(unsafe.StringData(s), len(s)) }

// newSK allocates an empty single-key page that holds key from base on, or a
// set leaf if the remainder or the key is too long for a page. It is a
// newLeafFunc; addSK adds the first value.
func newSK(key []byte, base int) *leafHead {
	if len(key) > maxKeyLen || len(key)-base > skpage.MaxRemainder {
		return newSetLeaf[string](key, base)
	}
	return skLeaf(skpage.Empty(key[base:], len(key)))
}

// addSK adds v to the values of the page l of key, which sits in slot loc. A
// value that the page cannot hold makes the key a set leaf (spillSK).
func addSK(loc **header, l *leafHead, key []byte, v string) {
	p := asSK(l)
	switch q, res := p.Add(view(v)); {
	case res == skpage.Full:
		*loc = leafHdr(spillSK(p, key, v))
	case q != p:
		*loc = leafHdr(skLeaf(q))
	}
}

// spillSK returns a set leaf with the values of page p of key and v, the value
// overflow of a key that outgrew its page.
func spillSK(p *skpage.Page, key []byte, v string) *leafHead {
	nl := newSetLeaf[string](key, skLeaf(p).base())
	s := vals[string](nl)
	p.Strings(func(x string) bool { s.Add(x); return true })
	s.Add(v)
	return nl
}

// removeSK removes v from the page l of key, which sits in the tree; rk is the
// map's rekeyFunc.
func (t *Tree) removeSK(l *leafHead, key []byte, v string, rk rekeyFunc) {
	switch q, ok := asSK(l).Remove(view(v)); {
	case !ok:
	case q == nil:
		t.remove(key, rk)
	case q != asSK(l):
		*t.findSlot(key) = leafHdr(skLeaf(q))
	}
}

// unspillSK returns the page for the set leaf l, whose values have shrunk to
// fit a page of skpage.BackLimit bytes, or nil if they do not.
func unspillSK(l *leafHead) *skpage.Page {
	if l.klen == longKey {
		return nil
	}
	s := vals[string](l)
	if s.Len() > skpage.BackLimit/2 { // every value takes two bytes at least
		return nil
	}
	need, ok := skpage.Header+int(l.klen), true
	vs := make([][]byte, 0, s.Len())
	s.Each(func(x string) bool {
		need += 1 + len(x)
		if ok = len(x) <= skpage.MaxValue && need <= skpage.BackLimit; ok {
			vs = append(vs, view(x))
		}
		return ok
	})
	if !ok {
		return nil
	}
	return skpage.Build(l.stored(), int(l.kl), vs)
}

// rekeySK is the map's rekeyFunc for strings (see rekeyFunc): a page takes the
// bytes in front of its remainder, or becomes a set leaf if they do not fit; a
// set leaf is handled as in rekey.
func rekeySK(l *leafHead, pre []byte, b, pathLen int) *leafHead {
	if l.kind == kSet {
		return rekey[string](l, pre, b, pathLen)
	}
	front := make([]byte, l.base()-pathLen)
	fillHead(front, pre, b, pathLen)
	p := asSK(l)
	if q := p.Prepend(front); q != nil {
		return skLeaf(q)
	}
	nl := newSetLeaf[string](wholeKey(l, pre, b), pathLen)
	s := vals[string](nl)
	p.Strings(func(x string) bool { s.Add(x); return true })
	return nl
}
