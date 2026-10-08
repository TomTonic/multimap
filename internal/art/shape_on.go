//go:build mkstats

package art

import (
	"encoding/binary"
	"fmt"
	"math/bits"
	"slices"
	"strings"
)

// Shape describes the routing part of the tree for the questions of step 6 (docs/redesign/PLAN.md): how many
// byte nodes of each type, how full, how many carry a tail for their common prefix, how many have only pages
// below them, and how many nodes a lookup passes on its way to a key. Diagnosis only (build tag mkstats).
func (m *Map[T]) Shape() string {
	type typeStats struct{ count, children, tails, endPages, onlyPages, bytes int }
	var st [64]typeStats
	n5kids := [6]int{}
	depthKeys := map[int]int{}
	keys, plens := 0, [3]int{} // common prefix 0, 1..12, 13+
	var walk func(n *header, depth int)
	walk = func(n *header, depth int) {
		if isPage(n.objType) {
			k := 1
			if isMultiKey(n.objType) {
				k = m.multiKeyObject(n).Keys
			}
			depthKeys[depth] += k
			keys += k
			return
		}
		t := n.objType & objTypeMask
		s := &st[t]
		s.count++
		s.bytes += m.object(n).Size
		pl := n.prefixLen()
		switch {
		case pl == 0:
			plens[0]++
		case pl <= 12:
			plens[1]++
		default:
			plens[2]++
		}
		if tailClass(pl) != tailNone {
			s.tails++
		}
		kids, pages := 0, 0
		if e := endPageOf(n); e != nil {
			s.endPages++
			depthKeys[depth+1]++
			keys++
		}
		eachByteNode(n, func(_ byte, c *header) {
			kids++
			if isPage(c.objType) {
				pages++
			}
			walk(c, depth+1)
		})
		s.children += kids
		if pages == kids {
			s.onlyPages++
		}
		if t == kN5 {
			n5kids[min(kids, 5)]++
		}
	}
	if m.t.root == nil {
		return "empty tree\n"
	}
	walk(m.t.root, 0)
	sel := m.selectivity()
	var b strings.Builder
	fmt.Fprintf(&b, "| node | count | bytes | children (mean) | with tail | with end page | only pages below |\n|---|--:|--:|--:|--:|--:|--:|\n")
	for _, t := range []objType{kN5, kN12, kN26, kN58, kN256} {
		s := st[t]
		if s.count == 0 {
			continue
		}
		fmt.Fprintf(&b, "| %s | %d | %d | %.1f | %d | %d | %d |\n", objTypeLabels[t], s.count, s.bytes, float64(s.children)/float64(s.count), s.tails, s.endPages, s.onlyPages)
	}
	fmt.Fprintf(&b, "\nN5 by number of byte children (0..5): %v; common prefix of the nodes: none %d, 1-12 bytes %d, 13+ bytes %d\n", n5kids, plens[0], plens[1], plens[2])
	sum, maxD := 0, 0
	for d, k := range depthKeys {
		sum += d * k
		maxD = max(maxD, d)
	}
	fmt.Fprintf(&b, "nodes on the way to a key: mean %.2f, by depth:", float64(sum)/float64(max(1, keys)))
	for d := 0; d <= maxD; d++ {
		fmt.Fprintf(&b, " %d: %.1f %%", d, 100*float64(depthKeys[d])/float64(max(1, keys)))
	}
	b.WriteString("\n")
	b.WriteString(sel)
	b.WriteString(m.fingerprints())
	return b.String()
}

// wyMix, wh64Det and fingerprint are the fingerprint under study (review, page-search-design.md): wyhash's mixing of the
// Set3 project (hashing.WH64Det), applied to a block of 16 bytes: the length of the input (2 bytes) and its last up to
// 14 bytes, right-aligned with zeros in front, as two words; the hash of the first word is the seed of the second.
func wyMix(a, b uint64) uint64 {
	hi, lo := bits.Mul64(a, b)
	return hi ^ lo
}

func wh64Det(val, seed uint64) uint64 {
	const m5, p1 = 0x1d8e4e27c47d124f, 0xf20a3e5e0b7b9731
	return wyMix(m5^8, wyMix(val^p1, bits.RotateLeft64(val, 32)^seed))
}

// fingerprintEnds is fingerprint with the first 6 and the last 8 bytes of an input longer than 14 bytes, instead of
// its last 14: the keys of a page differ near the start of their remainders as often as at the end.
func fingerprintEnds(in []byte) uint64 {
	if len(in) <= 14 {
		return fingerprint(in)
	}
	var block [16]byte
	binary.BigEndian.PutUint16(block[:], uint16(len(in)))
	copy(block[2:8], in[:6])
	copy(block[8:], in[len(in)-8:])
	return wh64Det(binary.LittleEndian.Uint64(block[8:]), wh64Det(binary.LittleEndian.Uint64(block[:8]), 0))
}

func fingerprint(in []byte) uint64 {
	var block [16]byte
	binary.BigEndian.PutUint16(block[:], uint16(len(in)))
	tail := in[max(0, len(in)-14):]
	copy(block[16-len(tail):], tail)
	w1, w2 := binary.LittleEndian.Uint64(block[:8]), binary.LittleEndian.Uint64(block[8:])
	return wh64Det(w2, wh64Det(w1, 0))
}

// fingerprints counts, over the keys of every multi-key page, how many other keys of the same page share a key's
// fingerprint in its lowest 7 or 8 bits: the compares a lookup would make in vain. Two inputs: (A) the remainder of
// the key in its page, (B) the whole key. Chance would give (keys of the page - 1) / 128 or / 256.
func (m *Map[T]) fingerprints() string {
	type tally struct{ others, hit, chance float64 }
	var res [3][2]tally // input A/B/A', bits 7/8
	keys, worst := 0, [3][2]int{}
	var example []string // the remainders that share the fingerprint (A, 8 bits) in the worst page
	var visit func(n *header, path []byte)
	visit = func(n *header, path []byte) {
		if isPage(n.objType) {
			if !isMultiKey(n.objType) {
				return
			}
			var cp []byte
			var rems [][]byte
			each := func(rem []byte, first bool) {
				if first {
					rems = append(rems, rem)
				}
			}
			if m.flat == 3 {
				p := asMKStr(n)
				cp = p.CP()
				p.Each(func(rem, _ []byte, first bool) bool { each(rem, first); return true })
			} else {
				p := asMKFix(n)
				cp = p.CP()
				p.Each(func(rem []byte, _ T, first bool) bool { each(rem, first); return true })
			}
			var fp [3][]uint64
			for _, r := range rems {
				fp[0] = append(fp[0], fingerprint(r))
				fp[1] = append(fp[1], fingerprint(append(append(append([]byte{}, path...), cp...), r...)))
				fp[2] = append(fp[2], fingerprintEnds(r))
			}
			k := len(rems)
			keys += k
			for in := range 3 {
				for bi, mask := range []uint64{0x7f, 0xff} {
					for i := range k {
						same := 0
						for j := range k {
							if j != i && fp[in][i]&mask == fp[in][j]&mask {
								same++
							}
						}
						res[in][bi].others += float64(same)
						if same > 0 {
							res[in][bi].hit++
						}
						if in == 0 && bi == 1 && same > worst[0][1] {
							example = example[:0]
							for j := range k {
								if fp[0][i]&0xff == fp[0][j]&0xff {
									example = append(example, fmt.Sprintf("%q", append(append([]byte{}, cp...), rems[j]...)))
								}
							}
						}
						worst[in][bi] = max(worst[in][bi], same)
						res[in][bi].chance += float64(k-1) / float64(mask+1)
					}
				}
			}
			return
		}
		path = appendPrefix(path, n)
		if e := endPageOf(n); e != nil {
			visit(singleKeyHdr(e), path)
		}
		eachByteNode(n, func(b byte, c *header) { visit(c, append(slices.Clip(path), b)) })
	}
	visit(m.t.root, nil)
	if keys == 0 {
		return ""
	}
	var b strings.Builder
	k := float64(keys)
	for in, name := range []string{"(A) remainder, last 14", "(B) whole key, last 14", "(A') remainder, first 6 and last 8"} {
		for bi, bitsN := range []int{7, 8} {
			r := res[in][bi]
			fmt.Fprintf(&b, "fingerprint %s, %d bits: other keys of the page with the same fingerprint %.4f a key (chance %.4f), keys with one or more %.2f %%, most %d\n",
				name, bitsN, r.others/k, r.chance/k, 100*r.hit/k, worst[in][bi])
		}
	}
	fmt.Fprintf(&b, "the keys (key part and remainder) that share a fingerprint (A, 8 bits) in the worst page: %s\n", strings.Join(example, " "))
	return b.String()
}

// selectivity describes how well the first bytes of the remainders tell the keys of a multi-key page apart, for the
// question of a list of first bytes searched with SWAR (docs/redesign/review-2026-10.md, analysis b): for a key
// of a page, the number of keys of its page that share its first byte (or its first two bytes) — the keys a
// lookup would still compare in full — averaged over all keys of multi-key pages; and the share of the keys that
// are alone with their first byte(s).
func (m *Map[T]) selectivity() string {
	var keys, remBytes int
	var share, alone [3]int // by prefix length 1, 2 (index 1, 2)
	var pageKeys []int
	var visit func(n *header)
	visit = func(n *header) {
		if isPage(n.objType) {
			if !isMultiKey(n.objType) {
				return
			}
			var rems [][]byte
			each := func(rem []byte, first bool) {
				if first {
					rems = append(rems, rem)
				}
			}
			if m.flat == 3 {
				asMKStr(n).Each(func(rem, _ []byte, first bool) bool { each(rem, first); return true })
			} else {
				asMKFix(n).Each(func(rem []byte, _ T, first bool) bool { each(rem, first); return true })
			}
			pageKeys = append(pageKeys, len(rems))
			for _, r := range rems {
				keys++
				remBytes += len(r)
				for w := 1; w <= 2; w++ {
					same := 0
					for _, o := range rems {
						if string(o[:min(w, len(o))]) == string(r[:min(w, len(r))]) {
							same++
						}
					}
					share[w] += same
					if same == 1 {
						alone[w]++
					}
				}
			}
			return
		}
		if e := endPageOf(n); e != nil {
			visit(singleKeyHdr(e))
		}
		eachByteNode(n, func(_ byte, c *header) { visit(c) })
	}
	visit(m.t.root)
	if keys == 0 {
		return "no multi-key pages\n"
	}
	k := float64(keys)
	return fmt.Sprintf("first bytes of the remainders in multi-key pages (%d keys in %d pages, %.1f keys a page, remainder %.1f bytes): keys sharing the first byte with a key %.2f (alone %.0f %%), the first two bytes %.2f (alone %.0f %%)\n",
		keys, len(pageKeys), k/float64(len(pageKeys)), float64(remBytes)/k, float64(share[1])/k, 100*float64(alone[1])/k, float64(share[2])/k, 100*float64(alone[2])/k)
}
