package page

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

// collide returns n remainders of three bytes with the same fingerprint, the i-th beginning with first[i]
// (found by trying all two-byte endings).
func collide(t *testing.T, first []byte) [][]byte {
	t.Helper()
	want := Fingerprint([]byte{first[0], 0, 0})
	var out [][]byte
	for _, f := range first {
		var hit []byte
		for x := 0; x < 65536 && hit == nil; x++ {
			if r := []byte{f, byte(x >> 8), byte(x)}; Fingerprint(r) == want {
				hit = r
			}
		}
		if hit == nil {
			t.Fatalf("no remainder beginning with %q has fingerprint %d", f, want)
		}
		out = append(out, hit)
	}
	return out
}

// pageWithCollisions returns the many-key page (key part "K") of 21 keys in which the remainders a, c and d
// (three bytes, one fingerprint) lie in the first, the second and the third group of eight fingerprints, among
// fillers of two bytes, and a remainder e of the same fingerprint and length that is not in the page. Each value
// is its key.
func pageWithCollisions(t *testing.T) (p *Str, in [][]byte, e []byte) {
	t.Helper()
	c := collide(t, []byte("bpyz"))
	in = [][]byte{c[0]}
	for i := range 10 {
		in = append(in, fmt.Appendf(nil, "m%d", i))
	}
	in = append(in, c[1])
	for i := range 8 {
		in = append(in, fmt.Appendf(nil, "q%d", i))
	}
	in = append(in, c[2])
	var rests, vals [][]byte
	for _, r := range in {
		rests, vals = append(rests, append([]byte("K"), r...)), append(vals, r)
	}
	p = BuildStrings(rests, vals)
	if p == nil || p.Len() != 21 || p.OneKey() || string(p.CP()) != "K" {
		t.Fatalf("setup: page %v", p)
	}
	return p, in, c[3]
}

// TestFindKeysWithTheSameFingerprint shows that a lookup is exact when keys of a page share a fingerprint.
//
// A user looks keys up in a map whose pages hold several keys; the page narrows the search down with one-byte
// fingerprints, which two keys of a page can share (about one in 256 pairs), and must still return exactly the
// value of the key asked for, and nothing for a key that is not there.
//
// Expected: in a page of 21 keys, three of which have the same fingerprint and length and lie in the first, the
// second and the third group of eight fingerprints, Get and EachValue find each of the three (and every other key)
// with its own value, a fourth remainder with that fingerprint and length is not found, and neither is a remainder
// of that fingerprint whose length is another's; for a key part that is the whole page of two keys at the end of a
// small object (where the word of fingerprints cannot be read whole), both keys are found.
func TestFindKeysWithTheSameFingerprint(t *testing.T) {
	p, in, e := pageWithCollisions(t)
	m, la := p.mem(), p.lay(true)
	if m[la.fp] != m[la.fp+11] || m[la.fp] != m[la.fp+20] || m[la.fp] != Fingerprint(e) {
		t.Fatalf("setup: the fingerprints of the values 0, 11 and 20 are %d, %d, %d; of e %d", m[la.fp], m[la.fp+11], m[la.fp+20], Fingerprint(e))
	}
	for i, r := range in {
		pos, off, ok := find(m, la.kl, la.currentValues, la.rem, r)
		if !ok || pos != i || !bytes.Equal(m[off:off+len(r)], r) {
			t.Errorf("find(%q) = %d, %d, %v; want idx %d", r, pos, off, ok, i)
		}
		if v, ok := p.Get(append([]byte("K"), r...)); !ok || !bytes.Equal(v, r) {
			t.Errorf("Get(K%q) = %q, %v", r, v, ok)
		}
		var vs []string
		if !p.EachValue(append([]byte("K"), r...), func(v []byte) bool { vs = append(vs, string(v)); return true }) || len(vs) != 1 || vs[0] != string(r) {
			t.Errorf("EachValue(K%q) = %q", r, vs)
		}
	}
	if _, ok := p.Get(append([]byte("K"), e...)); ok {
		t.Errorf("Get found %q, which is not in the page but has the fingerprint of three keys that are", e)
	}
	if _, _, ok := find(m, la.kl, la.currentValues, la.rem, e); ok {
		t.Errorf("find found %q", e)
	}
	for _, r := range []string{"", "b", "bx", "bxyz", strings.Repeat("b", 300)} {
		if _, ok := p.Get([]byte("K" + r)); ok {
			t.Errorf("Get(K%q) found a key that is not there", r)
		}
	}

	// the same for Fixed
	var rests [][]byte
	var vals []uint64
	for i, r := range in {
		rests, vals = append(rests, append([]byte("K"), r...)), append(vals, uint64(i))
	}
	f := BuildFixed(rests, vals)
	for i, r := range in {
		if v, ok := f.Get[uint64](append([]byte("K"), r...)); !ok || v != uint64(i) {
			t.Errorf("Fixed.Get(K%q) = %d, %v", r, v, ok)
		}
	}
	if _, ok := f.Get[uint64](append([]byte("K"), e...)); ok {
		t.Errorf("Fixed.Get found %q", e)
	}

	// a page of 64 bytes whose fingerprints end less than a word before the end of the object
	cp := strings.Repeat("P", 51)
	small := BuildStrings(bs(cp, cp+"x"), bs("1", "2"))
	if small == nil || small.Size() != 64 || small.Used() != 4+51+6+1 {
		t.Fatalf("setup: small page %v", small)
	}
	for i, k := range []string{cp, cp + "x"} {
		if v, ok := small.Get([]byte(k)); !ok || string(v) != fmt.Sprint(i+1) {
			t.Errorf("small page: Get(%q) = %q, %v", k, v, ok)
		}
	}
	if _, ok := small.Get([]byte(cp + "y")); ok {
		t.Error("small page found a key that is not there")
	}
}

// TestFindFurtherValues shows that a key with several values is found by its first value, and that the fingerprint
// of a further value cannot be taken for a key.
//
// A user adds several values to a key; the page keeps the key once and marks the further values, which have no
// remainder and the fingerprint 0, so a key whose own fingerprint is 0 must not find them.
//
// Expected: in a page whose second key has three values, the first and the second key are found with the right
// values in the right order; a remainder with the fingerprint 0 that is not in the page is not found, and one
// that is in the page (added with that fingerprint) is found next to the values with Further; after a Remove of
// the first value of the second key, the next value takes its place and the key is still found.
func TestFindFurtherValues(t *testing.T) {
	var zero []byte // a remainder of three bytes with the fingerprint 0
	for x := range 65536 {
		if r := []byte{'z', byte(x >> 8), byte(x)}; Fingerprint(r) == 0 {
			zero = r
			break
		}
	}
	if zero == nil {
		t.Fatal("no remainder with the fingerprint 0")
	}
	p := BuildStrings(bs("Ka", "Kb", "Kb", "Kb", "Kd"), bs("1", "2", "3", "4", "5"))
	if _, ok := p.Get(append([]byte("K"), zero...)); ok {
		t.Error("a remainder with the fingerprint 0 found a idx with Further")
	}
	var vs []string
	p.EachValue([]byte("Kb"), func(v []byte) bool { vs = append(vs, string(v)); return true })
	if fmt.Sprint(vs) != "[2 3 4]" {
		t.Errorf("values of Kb: %v", vs)
	}
	q, res := p.Add(append([]byte("K"), zero...), []byte("0"))
	if res != Added {
		t.Fatalf("Add of the remainder with the fingerprint 0: %v", res)
	}
	if v, ok := q.Get(append([]byte("K"), zero...)); !ok || string(v) != "0" {
		t.Errorf("Get of the remainder with the fingerprint 0: %q, %v", v, ok)
	}
	q, rm := q.Remove([]byte("Kb"), []byte("2"))
	if rm != Removed {
		t.Fatalf("Remove: %v", rm)
	}
	if v, ok := q.Get([]byte("Kb")); !ok || string(v) != "3" {
		t.Errorf("Get(Kb) after Remove of its first value: %q, %v", v, ok)
	}
}

// TestWidenRecomputesFingerprints shows that the keys of a page are still found after a key outside its
// common prefix has widened it.
//
// A user adds a key that shares less than the whole common prefix of a page with the keys in it; the page takes
// the key in and shortens its common prefix, so every remainder grows by the bytes the common prefix gives up,
// and a lookup of any key must still succeed.
//
// Expected: after Widen of a page of strings and a page of words, every old key and the new one are found with
// their values, and the fingerprint list is that of the new remainders.
func TestWidenRecomputesFingerprints(t *testing.T) {
	keys := bs("Bahnhofa", "Bahnhofb", "Bahnhofb", "Bahnhofcd")
	p := BuildStrings(keys, bs("1", "2", "3", "4"))
	q := p.Widen([]byte("Bahz"), []byte("5"))
	if q == nil {
		t.Fatal("Widen refused")
	}
	checkFingerprints(t, q.mem(), q.lay(true))
	for k, want := range map[string]string{"Bahnhofa": "1", "Bahnhofb": "2", "Bahnhofcd": "4", "Bahz": "5"} {
		if v, ok := q.Get([]byte(k)); !ok || string(v) != want {
			t.Errorf("Get(%q) = %q, %v; want %q", k, v, ok, want)
		}
	}
	f := BuildFixed(keys, []uint64{1, 2, 3, 4})
	g := f.Widen([]byte("Ba"), uint64(5), false) // the new key goes first
	if g == nil {
		t.Fatal("Widen of Fixed refused")
	}
	checkFingerprints(t, g.mem(), g.lay(false))
	for k, want := range map[string]uint64{"Bahnhofa": 1, "Bahnhofb": 2, "Bahnhofcd": 4, "Ba": 5} {
		if v, ok := g.Get[uint64]([]byte(k)); !ok || v != want {
			t.Errorf("Fixed.Get(%q) = %d, %v; want %d", k, v, ok, want)
		}
	}
}
