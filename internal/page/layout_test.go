package page

import (
	"bytes"
	"encoding/binary"
	"slices"
	"strconv"
	"strings"
	"testing"
	"unsafe"
)

func le(b []byte, v uint64) []byte { return binary.LittleEndian.AppendUint64(b, v) }

// exampleKeys are the keys of the layout example of step5-one-page.md: "Bahnhof" is the common
// prefix, "strasse" has three values.
var exampleKeys = bs("Bahnhof", "Bahnhofsallee", "Bahnhofstrasse", "Bahnhofstrasse", "Bahnhofstrasse", "Bahnhofweg")

// TestPageLayout pins the bytes of the pages of the example of the design note.
//
// The layout is what the tree and every later change of the page rely on: the head of four bytes
// (type, the length of the key part, the number of slots, rawWords), the key part right behind it,
// then, as far as the page has them, the key lengths with 255 for a further value of the key before
// it, the value lengths, the remainders, free bytes, and the values, which end with the object.
//
// Expected: the six slots of the example are the 63-byte page of strings (class 64) and the 128-byte
// page of `uint64` with the bytes of the design note; a key with three values is the 32-byte page of
// strings and the 64-byte page of `uint64` in the one-key form, without key lengths; a page of
// pointers has rawWords of its byte area.
func TestPageLayout(t *testing.T) {
	// many keys, strings
	p := BuildStrings(exampleKeys, bs("Mitte", "Nord", "Ost", "Sued", "West", "Ring"))
	want := []byte{TypeManyKeys + 2, 7, 6, 0} // class 64
	want = append(want, "Bahnhof"...)
	want = append(want, 0, 6, 7, Further, Further, 3)  // key lengths
	want = append(want, 5, 4, 3, 4, 4, 4)              // value lengths
	want = append(want, "sallee"+"strasse"+"weg"...)   // remainders
	want = append(want, 0)                             // free
	want = append(want, "MitteNordOstSuedWestRing"...) // values, ending with the object
	if got := p.mem(); !bytes.Equal(got, want) || p.Size() != 64 || p.Used() != 39 {
		t.Fatalf("many keys, strings:\n got %v\nwant %v (size %d, used %d)", got, want, p.Size(), p.Used())
	}
	if p.Len() != 6 || p.Keys() != 4 || p.PrefixLen() != 7 || p.OneKey() || p.Class() != 1 || p.RawWords() != 0 {
		t.Errorf("Len %d Keys %d PrefixLen %d OneKey %v Class %d RawWords %d", p.Len(), p.Keys(), p.PrefixLen(), p.OneKey(), p.Class(), p.RawWords())
	}

	// many keys, uint64
	f := BuildFixed(exampleKeys, []uint64{7, 1, 2, 5, 9, 4})
	want = []byte{TypeManyKeys + 4, 7, 6, 0} // class 128
	want = append(want, "Bahnhof"...)
	want = append(want, 0, 6, 7, Further, Further, 3)
	want = append(want, "sallee"+"strasse"+"weg"...)
	want = append(want, make([]byte, 128-6*8-len(want))...)
	for _, v := range []uint64{7, 1, 2, 5, 9, 4} {
		want = le(want, v)
	}
	if got := f.mem(); !bytes.Equal(got, want) || f.Size() != 128 || f.Used() != 33 {
		t.Fatalf("many keys, uint64:\n got %v\nwant %v (size %d, used %d)", got, want, f.Size(), f.Used())
	}

	// many keys, pointers: the same bytes, rawWords is the byte area in words
	pp := BuildFixed(exampleKeys, []*uint64{&ptrPool[7], &ptrPool[1], &ptrPool[2], &ptrPool[5], &ptrPool[9], &ptrPool[4]})
	if pp.RawWords() != 5 || pp.Used() != 33 { // 33 bytes are 5 words
		t.Errorf("pointer page: rawWords %d, used %d", pp.RawWords(), pp.Used())
	}
	if got, want := pp.mem()[:4], []byte{TypeManyKeys + 4, 7, 6, 5}; !bytes.Equal(got, want) {
		t.Errorf("pointer page head %v, want %v", got, want)
	}

	// one key, strings
	one := BuildStrings(bs("Bahnhofstrasse", "Bahnhofstrasse", "Bahnhofstrasse"), bs("Ost", "Sued", "West"))
	want = []byte{TypeOneKey, 14, 3, 0}
	want = append(want, "Bahnhofstrasse"...)
	want = append(want, 3, 4, 4)
	want = append(want, "OstSuedWest"...)
	if got := one.mem(); !bytes.Equal(got, want) || one.Size() != 32 || !one.OneKey() || one.Keys() != 1 || one.Used() != 21 {
		t.Fatalf("one key, strings:\n got %v\nwant %v (size %d, used %d)", got, want, one.Size(), one.Used())
	}

	// one key, uint64
	onef := BuildFixed(bs("Bahnhofstrasse", "Bahnhofstrasse", "Bahnhofstrasse"), []uint64{2, 5, 9})
	want = []byte{TypeOneKey + 2, 14, 3, 0}
	want = append(want, "Bahnhofstrasse"...)
	want = append(want, make([]byte, 64-3*8-len(want))...)
	want = le(le(le(want, 2), 5), 9)
	if got := onef.mem(); !bytes.Equal(got, want) || onef.Size() != 64 || onef.Used() != 18 {
		t.Fatalf("one key, uint64:\n got %v\nwant %v (size %d, used %d)", got, want, onef.Size(), onef.Used())
	}

	// the keys of a page are found in every form
	for _, k := range []string{"Bahnhof", "Bahnhofsallee", "Bahnhofstrasse", "Bahnhofweg"} {
		if _, ok := p.Get([]byte(k)); !ok {
			t.Errorf("Get(%q)", k)
		}
	}
	if v, ok := one.Get([]byte("Bahnhofstrasse")); !ok || string(v) != "Ost" {
		t.Errorf("one-key Get: %q %v", v, ok)
	}
	if _, ok := one.Get([]byte("Bahnhofstrasse2")); ok {
		t.Error("one-key page found another key")
	}
	if _, ok := p.Get([]byte("Bahnhofx")); ok {
		t.Error("string page found a key that is not there")
	}
	if _, ok := f.Get[uint64]([]byte("Bahnhofx")); ok {
		t.Error("found a key that is not there")
	}
	if _, ok := onef.Get[uint64]([]byte("Bahnhof")); ok {
		t.Error("one-key page of uint64 found a key that is not there")
	}
	if _, ok := f.Get[uint64]([]byte("Bahn")); ok {
		t.Error("found a key outside the common prefix")
	}
	_ = unsafe.Sizeof(0)
}

// TestPageLimits shows what the page refuses and where its limits are.
//
// The tree asks the page for entries of a subtree and must be told when they do not make a page: no
// entry, too many slots, a remainder, a value or a key part beyond the limits, content beyond 512 bytes.
//
// Expected: nil for those, a page at the borders (254 bytes of remainder and of value are taken), and the
// page of a type without a Fixed page is nil.
func TestPageLimits(t *testing.T) {
	big := make([][]byte, 256)
	vals := make([][]byte, 256)
	one := make([]uint8, 256)
	for i := range big {
		big[i], vals[i] = []byte{'k'}, []byte{'v'}
	}
	if BuildStrings(nil, nil) != nil || BuildStrings(bs("a"), nil) != nil || BuildStrings(big, vals) != nil || BuildFixed(big, one) != nil || BuildFixed[uint64](nil, nil) != nil {
		t.Error("no entry, another number of values, or 256 slots gave a page")
	}
	if BuildStrings(bs("a"), bs(strings.Repeat("v", MaxValue+1))) != nil || BuildStrings(bs("a"), bs(strings.Repeat("v", MaxValue))) == nil {
		t.Error("value of 255 or 254 bytes")
	}
	long := strings.Repeat("k", 600)
	if BuildStrings(bs(long), bs("v")) != nil || BuildFixed(bs(long), []uint64{1}) != nil {
		t.Error("one key of 600 bytes")
	}
	if BuildStrings(bs("a", "a"+strings.Repeat("x", MaxRemainder+1)), bs("v", "w")) != nil || BuildStrings(bs("a", "a"+strings.Repeat("x", MaxRemainder)), bs("v", "w")) == nil {
		t.Error("remainder of 255 or 254 bytes")
	}
	if BuildFixed(bs("a", "a"+strings.Repeat("x", MaxRemainder+1)), []uint64{1, 2}) != nil || BuildFixed(bs("a", "a"+strings.Repeat("x", MaxRemainder)), []uint8{1, 2}) == nil {
		t.Error("remainder of 255 or 254 bytes, fixed")
	}
	if BuildFixed(bs("a", "b"), []string{"x", "y"}) != nil || BuildFixed(bs("a"), [][3]uint64{{}}) != nil || BuildFixed(bs("a"), []*int{nil}) == nil {
		t.Error("types: a string, three words, a pointer")
	}
	if Supported[string]() || !Supported[*int]() || HoldsPointers[uint64]() || !HoldsPointers[*int]() || !Supported[uint64]() || !Supported[[2]uint64]() || Supported[[3]uint64]() {
		t.Error("Supported and HoldsPointers")
	}
	if BuildStrings(bs("a"), bs(strings.Repeat("v", 254), strings.Repeat("w", 254))) != nil {
		t.Error("value count does not match")
	}
	// content beyond 512 bytes
	many := make([][]byte, 40)
	mv := make([][]byte, 40)
	for i := range many {
		many[i] = []byte{'a', byte('A' + i)}
		mv[i] = []byte(strings.Repeat("v", 20))
	}
	if BuildStrings(many, mv) != nil {
		t.Error("40 values of 20 bytes make a page of more than 512 bytes")
	}
	// 100 values of one key, one byte each: the one-key page of strings, class 512 refused at 255
	hund := make([][]byte, 100)
	hv := make([][]byte, 100)
	for i := range hund {
		hund[i], hv[i] = []byte("key"), []byte{byte(i)}
	}
	if p := BuildStrings(hund, hv); p == nil || p.Keys() != 1 || p.Len() != 100 || !p.OneKey() {
		t.Error("100 values of one key")
	}
}

// TestPageOneKeyBecomesMany shows the transitions between the two forms.
//
// A user who adds a second key near a key that has a few values gets one page for both; one who
// removes the second key again gets the cheaper page of one key back, in place.
//
// Expected: Add of another key to a one-key page returns a new page of the many-key form with all the values
// (for 12 values of the first key too, the slices of the pairing are not limited), and Remove of the second
// key turns the page back into the one-key form in place, with the key part of the whole key and no key
// lengths; a key that does not fit is refused.
func TestPageOneKeyBecomesMany(t *testing.T) {
	for _, n := range []int{1, 3, 12} {
		var vs [][]byte
		var ks [][]byte
		for i := range n {
			ks, vs = append(ks, []byte("street-1")), append(vs, []byte{byte('a' + i)})
		}
		p := BuildStrings(ks, vs)
		q, res := p.Add([]byte("street-2"), []byte("z"))
		if res != Added || q == p || q.OneKey() || q.Len() != n+1 || q.Keys() != 2 || q.PrefixLen() != 7 {
			t.Fatalf("%d values: Add of another key: %v, %v keys %d prefix %d", n, res, q.OneKey(), q.Keys(), q.PrefixLen())
		}
		r, rm := q.Remove([]byte("street-2"), []byte("z"))
		if rm != Gone || r != q || !r.OneKey() || r.Len() != n || r.Keys() != 1 || string(r.CP()) != "street-1" {
			t.Fatalf("%d values: Remove of the second key: %v, one key %v, %d values, key part %q", n, rm, r.OneKey(), r.Len(), r.CP())
		}
		// the same for uint64: the one-key page of fixed values
		fs := make([]uint64, n)
		for i := range fs {
			fs[i] = uint64(i)
		}
		f := BuildFixed(ks, fs)
		g, res := f.Add([]byte("street-0"), uint64(99), false)
		if res != Added || g == f || g.OneKey() || g.Len() != n+1 {
			t.Fatalf("fixed, %d values: Add of another key: %v", n, res)
		}
		h, rm := g.Remove([]byte("street-0"), uint64(99), false)
		if rm != Gone || !h.OneKey() || h.Len() != n {
			t.Fatalf("fixed, %d values: Remove of the other key: %v", n, rm)
		}
	}
	// a one-key page that gets a key that does not fit: Full
	big := BuildStrings(bs("k", "k"), bs(strings.Repeat("v", 250), strings.Repeat("w", 250)))
	if q, res := big.Add([]byte("j"), []byte(strings.Repeat("x", 100))); res != Full || q != big {
		t.Errorf("another key that does not fit: %v", res)
	}
	bigf := BuildFixed(bs("k"), []uint64{1})
	if q, res := bigf.Add([]byte("j"+strings.Repeat("y", 520)), uint64(2), false); res != Full || q != bigf {
		t.Errorf("fixed, another key that does not fit: %v", res)
	}
}

// TestPageSkipAndPrepend shows the key part changing in place when the path of the page changes.
//
// A byte node above a page takes the first bytes of its key part (Skip); when the node goes away the
// page gets them back (Prepend), in place if its object holds them, else in a larger one.
//
// Expected: after Skip the keys are the old ones without their first bytes, with all their values; Prepend
// restores them; Prepend into a class that is full makes a larger page; Prepend beyond 512 bytes of content is nil; for a page of pointers the byte area decides whether it is in place.
func TestPageSkipAndPrepend(t *testing.T) {
	p := BuildStrings(exampleKeys, bs("Mitte", "Nord", "Ost", "Sued", "West", "Ring"))
	p.Skip(3)
	if string(p.CP()) != "nhof" || p.Len() != 6 {
		t.Fatalf("after Skip: %q", p.CP())
	}
	if v, ok := p.Get([]byte("nhofstrasse")); !ok || string(v) != "Ost" {
		t.Fatalf("Get after Skip: %q %v", v, ok)
	}
	q := p.Prepend([]byte("Bah"))
	if q != p || string(q.CP()) != "Bahnhof" {
		t.Fatalf("Prepend in place: %q", q.CP())
	}
	big := q.Prepend([]byte(strings.Repeat("P", 40)))
	if big == nil || big == q || big.Size() <= q.Size() || len(big.CP()) != 47 {
		t.Fatalf("Prepend of 40 bytes: %v", big)
	}
	if big.Prepend([]byte(strings.Repeat("P", 600))) != nil || big.Prepend([]byte(strings.Repeat("P", 470))) != nil {
		t.Error("Prepend beyond the largest class")
	}
	// the one-key form
	one := BuildStrings(bs("Bahnhofstrasse"), bs("Ost"))
	one.Skip(8)
	if string(one.CP()) != "trasse" {
		t.Fatalf("one-key Skip: %q", one.CP())
	}
	if o := one.Prepend([]byte("Bahnhofs")); o != one || string(o.CP()) != "Bahnhofstrasse" {
		t.Fatalf("one-key Prepend: %q", o.CP())
	}
	// uint64: a larger class and the limits
	f := BuildFixed(exampleKeys, []uint64{7, 1, 2, 5, 9, 4})
	f.Skip(7)
	if g := f.Prepend[uint64]([]byte("Bahnhof"), false); g != f || string(g.CP()) != "Bahnhof" {
		t.Fatalf("fixed Prepend in place: %q", g.CP())
	}
	g := f.Prepend[uint64]([]byte(strings.Repeat("P", 60)), false)
	if g == nil || g == f || g.Size() <= f.Size() {
		t.Fatalf("fixed Prepend of 60 bytes: %v", g)
	}
	if g.Prepend[uint64]([]byte(strings.Repeat("P", 600)), false) != nil || g.Prepend[uint64]([]byte(strings.Repeat("P", 470)), false) != nil {
		t.Error("fixed Prepend beyond the limits")
	}
	// pointers: the byte area of the typed object is fixed: 33 bytes are 5 words (40 bytes), so 7 more bytes are in place
	pp := BuildFixed(exampleKeys, []*uint64{&ptrPool[7], &ptrPool[1], &ptrPool[2], &ptrPool[5], &ptrPool[9], &ptrPool[4]})
	if a := pp.Prepend[*uint64]([]byte("1234567"), true); a != pp || a.RawWords() != 5 {
		t.Errorf("pointer page: Prepend of 7 bytes in place: %v", a == pp)
	}
	if b := pp.Prepend[*uint64]([]byte("1234567"), true); b == pp || b.RawWords() <= 5 {
		t.Errorf("pointer page: Prepend beyond the byte area is a new object: %v, rawWords %d", b == pp, b.RawWords())
	}
}

// TestPageKeysUpTo shows that the tree can ask whether a page has at least a few keys without counting
// them all.
//
// The tree decides on a removal whether one key is left or at most two; for a page of one value each the scan
// stops after the third slot.
//
// Expected: for pages of 1, 2, 3 and 5 keys, some with several values, KeysUpTo(limit) is the smaller of the
// number of keys and limit, and Keys is the number of keys.
func TestPageKeysUpTo(t *testing.T) {
	for _, keys := range []int{1, 2, 3, 5} {
		for _, per := range []int{1, 3} {
			var rests, vals [][]byte
			for k := range keys {
				for v := range per {
					rests = append(rests, []byte{byte('a' + k)})
					vals = append(vals, []byte{byte('0' + v)})
				}
			}
			p := BuildStrings(rests, vals)
			f := BuildFixed(rests, make([]uint8, len(rests)))
			if p.Keys() != keys || f.Keys() != keys || p.OneKey() != (keys == 1) {
				t.Fatalf("%d keys of %d values: Keys %d and %d", keys, per, p.Keys(), f.Keys())
			}
			for limit := 1; limit <= 6; limit++ {
				if got, want := p.KeysUpTo(limit), min(keys, limit); got != want || f.KeysUpTo(limit) != want {
					t.Errorf("%d keys of %d values: KeysUpTo(%d) = %d and %d, want %d", keys, per, limit, got, f.KeysUpTo(limit), want)
				}
			}
		}
	}
}

// TestPageWidenRefusals shows when a key outside the common prefix does not join the page.
//
// The tree calls Widen where it would put a byte node above the page, so that a key that differs early
// joins its neighbours when they fit one page.
//
// Expected: Widen takes a key that leaves the prefix after some bytes, as the first key or the last, with
// the old prefix tail in front of the old remainders, and refuses (nil) a remainder beyond 254 bytes, a
// longest old remainder that no longer fits after the prefix tail, content beyond 512 bytes, a value beyond
// 254 bytes.
func TestPageWidenRefusals(t *testing.T) {
	p := BuildStrings(bs("abc1", "abc2"), bs("x", "y"))
	for _, tt := range []struct {
		key string
		pos int // where the new key goes
	}{{"abd", 2}, {"aaa", 0}, {"a", 0}} {
		q := p.Widen([]byte(tt.key), []byte("n"))
		if q == nil || q.Keys() != 3 {
			t.Fatalf("Widen(%q) refused", tt.key)
		}
		var got []string
		cp := string(q.CP())
		q.Each(func(rem, _ []byte, _ bool) bool { got = append(got, cp+string(rem)); return true })
		if got[tt.pos] != tt.key {
			t.Errorf("Widen(%q): keys %v", tt.key, got)
		}
	}
	if p.Widen([]byte("a"+strings.Repeat("z", 255)), []byte("n")) != nil {
		t.Error("remainder of 256 bytes")
	}
	if p.Widen([]byte("abd"), []byte(strings.Repeat("n", 255))) != nil {
		t.Error("value of 255 bytes")
	}
	deep := BuildStrings(bs("abcdef", "abcdef"+strings.Repeat("c", 250)), bs("x", "y"))
	if deep.Widen([]byte("a"), []byte("n")) != nil {
		t.Error("a remainder that grows beyond 254 bytes by the prefix tail")
	}
	full := BuildStrings(bs("a1", "a2"), bs(strings.Repeat("x", 250), strings.Repeat("y", 250)))
	if full == nil || full.Widen([]byte("b"), []byte("n")) != nil {
		t.Error("content beyond 512 bytes")
	}
	f := BuildFixed(bs("abc1", "abc2"), []uint64{1, 2})
	if q := f.Widen([]byte("abd"), uint64(3), false); q == nil || q.Keys() != 3 {
		t.Error("fixed Widen")
	}
	if f.Widen([]byte("a"+strings.Repeat("z", 255)), uint64(3), false) != nil {
		t.Error("fixed Widen of a remainder of 256 bytes")
	}
	fdeep := BuildFixed(bs("abcdef", "abcdef"+strings.Repeat("c", 250)), []uint64{1, 2})
	if fdeep.Widen([]byte("a"), uint64(3), false) != nil {
		t.Error("fixed Widen: a remainder that grows beyond 254 bytes")
	}
	many := make([][]byte, 60)
	mv := make([]uint64, 60)
	for i := range many {
		many[i], mv[i] = []byte{'a', 'b', byte('A' + i)}, uint64(i)
	}
	ff := BuildFixed(many[:40], mv[:40])
	if ff.Widen([]byte("x"+strings.Repeat("y", 100)), uint64(3), false) != nil {
		t.Error("fixed Widen beyond 512 bytes")
	}
}

// TestPageLookupsAndRemovals covers the answers of the lookups and removals for what is not there, and the stops of
// the iterations.
//
// A user asks the index for a key that is not there, for a value the key does not have, stops a scan of the
// values early, and removes the last value of a key that is alone in its page.
//
// Expected: EachValue of an absent key is false and calls nothing, a stop after the first value is a stop, Each
// stops where it is told; Remove of an absent value or key is Absent and changes nothing; Remove of the only
// value of a page of one key makes the page nil; an Add that does not fit is Full and changes nothing.
func TestPageLookupsAndRemovals(t *testing.T) {
	calls := 0
	p := BuildStrings(exampleKeys, bs("Mitte", "Nord", "Ost", "Sued", "West", "Ring"))
	f := BuildFixed(exampleKeys, []uint64{7, 1, 2, 5, 9, 4})
	one := BuildStrings(bs("k", "k"), bs("a", "b"))
	onef := BuildFixed(bs("k", "k"), []uint64{1, 2})
	if p.EachValue([]byte("Bahn"), func([]byte) bool { calls++; return true }) || p.EachValue([]byte("Bahnhofx"), func([]byte) bool { calls++; return true }) ||
		f.EachValue([]byte("Bahn"), func(uint64) bool { calls++; return true }) || f.EachValue([]byte("Bahnhofx"), func(uint64) bool { calls++; return true }) ||
		one.EachValue([]byte("j"), func([]byte) bool { calls++; return true }) || onef.EachValue([]byte("j"), func(uint64) bool { calls++; return true }) || calls != 0 {
		t.Fatalf("a key that is not there: %d calls", calls)
	}
	if _, ok := p.Get([]byte("Bahn")); ok {
		t.Error("Get outside the common prefix")
	}
	if !p.EachValue([]byte("Bahnhofstrasse"), func([]byte) bool { calls++; return false }) || !f.EachValue([]byte("Bahnhofstrasse"), func(uint64) bool { calls++; return false }) ||
		!one.EachValue([]byte("k"), func([]byte) bool { calls++; return false }) || !onef.EachValue([]byte("k"), func(uint64) bool { calls++; return false }) || calls != 4 {
		t.Fatalf("a stop after the first value: %d calls", calls)
	}
	calls = 0
	if p.Each(func(_, _ []byte, _ bool) bool { calls++; return calls < 3 }) || f.Each(func(_ []byte, _ uint64, _ bool) bool { calls++; return calls < 6 }) || calls != 6 {
		t.Fatalf("Each stops where told: %d calls", calls)
	}
	calls = 0
	if one.Each(func(_, _ []byte, _ bool) bool { calls++; return false }) || onef.Each(func(_ []byte, _ uint64, _ bool) bool { calls++; return false }) || calls != 2 {
		t.Fatalf("Each of the one-key form stops where told: %d calls", calls)
	}
	// removals of what is not there
	for _, c := range []struct {
		key, val string
	}{{"Bahnhofstrasse", "none"}, {"Bahnhofx", "Ost"}, {"Bahn", "Ost"}} {
		if q, rm := p.Remove([]byte(c.key), []byte(c.val)); rm != Absent || q != p {
			t.Errorf("Remove(%q, %q) took something it must not", c.key, c.val)
		}
	}
	if q, rm := f.Remove([]byte("Bahnhofstrasse"), uint64(99), false); rm != Absent || q != f {
		t.Error("fixed Remove of a value the key does not have")
	}
	if q, rm := f.Remove([]byte("Bahn"), uint64(1), false); rm != Absent || q != f {
		t.Error("fixed Remove outside the prefix")
	}
	for _, c := range []struct {
		key, val string
	}{{"j", "a"}, {"k", "z"}} {
		if q, rm := one.Remove([]byte(c.key), []byte(c.val)); rm != Absent || q != one {
			t.Errorf("one-key Remove(%q, %q) took something it must not", c.key, c.val)
		}
		if q, rm := onef.Remove([]byte(c.key), fixedVal(c.val), false); rm != Absent || q != onef {
			t.Errorf("one-key fixed Remove(%q, %q) took something it must not", c.key, c.val)
		}
	}
	// the last value of a page of one key
	last := BuildStrings(bs("k"), bs("a"))
	lastf := BuildFixed(bs("k"), []uint64{1})
	if q, rm := last.Remove([]byte("k"), []byte("a")); rm != Gone || q != nil {
		t.Error("Remove of the only value of a one-key page")
	}
	if q, rm := lastf.Remove([]byte("k"), uint64(1), false); rm != Gone || q != nil {
		t.Error("fixed Remove of the only value of a one-key page")
	}
	// an add that does not fit
	if q, res := one.Add([]byte("k"), []byte(strings.Repeat("v", 255))); res != Full || q != one {
		t.Errorf("a value of 255 bytes for the key of a one-key page: %v", res)
	}
	if q, res := p.Add([]byte("Bahnhofstrasse"), []byte(strings.Repeat("v", 255))); res != Full || q != p {
		t.Errorf("a value of 255 bytes for a key of a many-key page: %v", res)
	}
	if q, res := f.Add([]byte("Bahnhof"+strings.Repeat("x", 255)), uint64(3), false); res != Full || q != f {
		t.Errorf("a remainder of 255 bytes: %v", res)
	}
}

// TestPageNeed shows that the sizes the tree computes before it builds a page are the sizes of the pages.
//
// The tree decides whether the entries of a subtree make a page from NeedStrings and NeedFixed, without
// building anything; a difference to the real page would make it build pages that do not fit, or refuse
// pages that do.
//
// Expected: for the six slots of the example, NeedStrings and NeedFixed are the bytes the pages use (the key area
// and the values), 63 and 81.
func TestPageNeed(t *testing.T) {
	p := BuildStrings(exampleKeys, bs("Mitte", "Nord", "Ost", "Sued", "West", "Ring"))
	f := BuildFixed(exampleKeys, []uint64{7, 1, 2, 5, 9, 4})
	if got := NeedStrings(6, 7, 16, 24); got != p.Used()+24 || got != 63 {
		t.Errorf("NeedStrings = %d, page uses %d", got, p.Used()+24)
	}
	if got := NeedFixed[uint64](6, 7, 16); got != f.Used()+48 || got != 81 {
		t.Errorf("NeedFixed = %d, page uses %d", got, f.Used()+48)
	}
}

// TestPageNew shows the pages for a key that no page holds yet.
//
// A user who adds a key to the index gets a page of that key with its first value at once; a key or a value
// that does not fit a page is refused (the tree then makes a value overflow), and the room that the values of a
// key may take back in a page is half of what the largest page has.
//
// Expected: NewStr and NewFixed are the one-key pages BuildStrings and BuildFixed make for one entry, byte for
// byte; they are nil for a value of 255 bytes, a key part that leaves no room, and (strings) content beyond 512
// bytes; Room is 512 less the head and the key part and BackFits is half of it.
func TestPageNew(t *testing.T) {
	for _, k := range []string{"", "k", "Bahnhofstrasse", strings.Repeat("x", 300)} {
		a, b := NewStr([]byte(k), []byte("Ost")), BuildStrings(bs(k), bs("Ost"))
		if !bytes.Equal(a.mem(), b.mem()) || !a.OneKey() || a.Len() != 1 {
			t.Errorf("NewStr(%q) differs from BuildStrings", k)
		}
		f, g := NewFixed([]byte(k), uint64(7), false), BuildFixed(bs(k), []uint64{7})
		if !bytes.Equal(f.mem(), g.mem()) || !f.OneKey() {
			t.Errorf("NewFixed(%q) differs from BuildFixed", k)
		}
		pp, qq := NewFixed([]byte(k), &ptrPool[7], true), BuildFixed(bs(k), []*uint64{&ptrPool[7]})
		if pp.RawWords() != qq.RawWords() || !bytes.Equal(pp.mem()[:Header+len(k)], qq.mem()[:Header+len(k)]) {
			t.Errorf("NewFixed of a pointer (%q) differs from BuildFixed", k)
		}
	}
	if NewStr([]byte("k"), []byte(strings.Repeat("v", 255))) != nil || NewStr([]byte(strings.Repeat("k", 507)), []byte("v")) != nil || NewStr([]byte(strings.Repeat("k", 506)), []byte("v")) == nil {
		t.Error("NewStr limits")
	}
	if NewFixed([]byte(strings.Repeat("k", 501)), uint64(1), false) != nil || NewFixed([]byte(strings.Repeat("k", 500)), uint64(1), false) == nil {
		t.Error("NewFixed limits")
	}
	if Room(0) != 508 || Room(100) != 408 || !BackFits(100, 204) || BackFits(100, 205) {
		t.Error("Room and BackFits")
	}
}

// TestPageSlotLimit shows that a page of tiny values stops at 255 slots.
//
// A user who gives one key many values of one byte each fills a page with 255 slots (the head holds n
// in one byte) before the 512 bytes are used up, and has to be told that the page is full, not get
// a corrupt one; a value that is already there is still found.
//
// Expected: the 256th value of a key in a Fixed page of bytes is refused with Full, the page keeps its 255
// values, and Add of one of them answers Present.
func TestPageSlotLimit(t *testing.T) {
	p := NewFixed[uint8](bs("key")[0], 0, false)
	if p == nil {
		t.Fatal("no page for the first value")
	}
	for v := 1; v < MaxEntries; v++ {
		q, res := p.Add(bs("key")[0], uint8(v), false)
		if res != AddedValue {
			t.Fatalf("value %d: %v", v, res)
		}
		p = q
	}
	if q, res := p.Add(bs("key")[0], uint8(255), false); res != Full || q != p || p.Len() != MaxEntries {
		t.Errorf("the 256th value: %v, %d values", res, p.Len())
	}
	if _, res := p.Add(bs("key")[0], uint8(7), false); res != Present {
		t.Errorf("a value that is there: %v", res)
	}
}

// TestPageEachSingle shows the plain scan of a page of one key.
//
// A user who lists the values of a key that has its own page gets them in the order they came in, and
// can stop the listing after any value; the tree does this for every key of a scan.
//
// Expected: EachSingle of a Str and of a Fixed page of one key calls fn with each value in slot order,
// stops when fn says so (and says so), and runs to completion otherwise.
func TestPageEachSingle(t *testing.T) {
	sp := BuildStrings(bs("key", "key", "key"), bs("a", "bcd", "ef"))
	fp := BuildFixed(bs("key", "key", "key"), []uint64{7, 8, 9})
	var gs []string
	if !sp.EachSingle(func(v []byte) bool { gs = append(gs, string(v)); return true }) || strings.Join(gs, ",") != "a,bcd,ef" {
		t.Errorf("strings: %v", gs)
	}
	var gf []uint64
	if !fp.EachSingle(func(v uint64) bool { gf = append(gf, v); return true }) || len(gf) != 3 || gf[0] != 7 || gf[2] != 9 {
		t.Errorf("words: %v", gf)
	}
	n := 0
	if sp.EachSingle(func([]byte) bool { n++; return n < 2 }) || n != 2 {
		t.Errorf("strings stopped after %d values", n)
	}
	n = 0
	if fp.EachSingle(func(uint64) bool { n++; return n < 2 }) || n != 2 {
		t.Errorf("words stopped after %d values", n)
	}
}

// TestPageValuesIn shows what a range scan gets from a page.
//
// A user who asks for the values or the keys of a key range gets, from each page the range touches, those of
// the keys inside the range in key order, a key with several values all of them; the tree passes the bounds
// behind the path to the page and ends the scan when a page reports a key above the upper bound.
//
// Expected: for pages of uint64 and of strings with the keys pa, pb (two values), pc and pd, ValuesIn and
// AppendStrings give every value without bounds, the values from or after a lower bound and up to or before an
// upper bound (inclusive and exclusive), nothing for an upper bound below the page and everything for bounds
// shorter or longer than the key part, and report a key above the upper bound exactly when one follows;
// AppendKeys gives the keys with the prefix in front, each once; a page of one key gives all its values.
func TestPageValuesIn(t *testing.T) {
	rests := bs("pa", "pb", "pb", "pc", "pd")
	fp := BuildFixed(rests, []uint64{1, 2, 3, 4, 5})
	sp := BuildStrings(rests, bs("1", "2", "3", "4", "5"))
	for _, tc := range []struct {
		name string
		b    Bounds
		want string
		keys string
		over bool
	}{
		{"no bounds", Bounds{}, "12345", "pa pb pc pd", false},
		{"from pb inclusive", Bounds{Lo: []byte("pb"), HasLo: true, LoIncl: true}, "2345", "pb pc pd", false},
		{"from pb exclusive", Bounds{Lo: []byte("pb"), HasLo: true}, "45", "pc pd", false},
		{"to pc inclusive", Bounds{Hi: []byte("pc"), HasHi: true, HiIncl: true}, "1234", "pa pb pc", true},
		{"to pc exclusive", Bounds{Hi: []byte("pc"), HasHi: true}, "123", "pa pb", true},
		{"pb to pb", Bounds{Lo: []byte("pb"), Hi: []byte("pb"), HasLo: true, HasHi: true, LoIncl: true, HiIncl: true}, "23", "pb", true},
		{"from a bound shorter than the key part", Bounds{Lo: []byte(""), HasLo: true}, "12345", "pa pb pc pd", false},
		{"to a bound below the page", Bounds{Hi: []byte("o"), HasHi: true, HiIncl: true}, "", "", true},
		{"to a bound above the page", Bounds{Hi: []byte("q"), HasHi: true}, "12345", "pa pb pc pd", false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			vs, over := fp.ValuesIn[uint64](&tc.b)
			got := ""
			for _, v := range vs {
				got += strconv.FormatUint(v, 10)
			}
			if got != tc.want || over != tc.over {
				t.Errorf("uint64: %q, over %v; want %q, %v", got, over, tc.want, tc.over)
			}
			ss, over := sp.AppendStrings(nil, &tc.b)
			if strings.Join(ss, "") != tc.want || over != tc.over {
				t.Errorf("strings: %q, over %v; want %q, %v", ss, over, tc.want, tc.over)
			}
			for _, p := range []*head{&fp.head, &sp.head} {
				buf, ends, over := p.AppendKeys(nil, nil, []byte("x"), &tc.b, p == &sp.head)
				var keys []string
				start := 0
				for _, e := range ends {
					keys = append(keys, string(buf[start+1:e]))
					if buf[start] != 'x' {
						t.Errorf("a key without the prefix: %q", buf[start:e])
					}
					start = e
				}
				if strings.Join(keys, " ") != tc.keys || over != tc.over {
					t.Errorf("keys: %q, over %v; want %q, %v", keys, over, tc.keys, tc.over)
				}
			}
		})
	}
	one := BuildFixed(bs("k", "k"), []uint64{7, 8})
	ones := BuildStrings(bs("k", "k"), bs("a", "b"))
	b := Bounds{Hi: []byte("a"), HasHi: true} // ignored: the caller checks the one key
	if vs, over := one.ValuesIn[uint64](&b); len(vs) != 2 || over {
		t.Errorf("one key, uint64: %v, %v", vs, over)
	}
	if ss, over := ones.AppendStrings(nil, &b); strings.Join(ss, "") != "ab" || over {
		t.Errorf("one key, strings: %q, %v", ss, over)
	}
}

// TestPageCompareOrdersAsBytes shows that the compare of remainders orders keys as bytes.Compare does.
//
// A user's keys are found and kept in order whatever their length: the search in a page compares the remainders
// with its own compare (a word-wide variant was measured and dropped, docs/redesign/page-search-design.md), which
// must agree with bytes.Compare at every length and every position of the first difference.
//
// Expected: for every pair of lengths from 0 to 24 and every position of a first difference (up or down), compare
// has the sign of bytes.Compare; locate finds each key of a page of keys of lengths 0 to 24, also at the very end of
// the page, and the place of a key that is not there.
func TestPageCompareOrdersAsBytes(t *testing.T) {
	sign := func(c int) int { return min(max(c, -1), 1) }
	base := []byte("abcdefghijklmnopqrstuvwxyz")
	for la := 0; la <= 24; la++ {
		for lb := 0; lb <= 24; lb++ {
			for d := -1; d < min(la, lb); d++ {
				for _, up := range []bool{false, true} {
					a := slices.Clone(base[:la])
					b := slices.Clone(base[:lb])
					if d >= 0 {
						if up {
							b[d]++
						} else {
							b[d]--
						}
					}
					if got, want := sign(compare(a, b)), bytes.Compare(a, b); got != want {
						t.Fatalf("compare(%q, %q) = %d, want %d", a, b, got, want)
					}
				}
			}
		}
	}
	var rests [][]byte
	for l := range 25 {
		rests = append(rests, append([]byte("k"), base[:l]...))
	}
	vals := make([]uint8, len(rests)) // one byte a value: 25 keys of up to 25 bytes fit one page
	p := BuildFixed(rests, vals)
	m, la := p.mem(), p.lay(false)
	for i, r := range rests {
		if pos, _, found := locate(m, la.kl, la.n, la.rem, r[la.l:]); !found || pos != i {
			t.Errorf("key of %d bytes: slot %d, found %v", len(r), pos, found)
		}
	}
	if pos, _, found := locate(m, la.kl, la.n, la.rem, []byte("kz")[la.l:]); found || pos != len(rests) {
		t.Errorf("a key after all: slot %d, found %v", pos, found)
	}
}
