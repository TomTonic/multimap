// Package keys generates the benchmark key corpora and probe sequences.
//
// All keys of a corpus live in one contiguous buffer, and the []byte and string
// views are stored in contiguous arrays, so that walking a probe sequence is a
// sequential, prefetch-friendly stream. Otherwise every probe would add a cache
// miss of its own to every candidate and dilute the differences between them.
package keys

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
	"slices"
	"strconv"
	"unsafe"

	"github.com/TomTonic/rtcompare"
)

// Set is a packed list of keys with []byte and string views.
type Set struct {
	B [][]byte
	S []string
}

// Pack copies keys into one buffer and builds both views.
func Pack(keys [][]byte) Set {
	total := 0
	for _, k := range keys {
		total += len(k)
	}
	buf := make([]byte, 0, total)
	offs := make([]int, len(keys)+1)
	for i, k := range keys {
		buf = append(buf, k...)
		offs[i+1] = len(buf)
	}
	s := Set{B: make([][]byte, len(keys)), S: make([]string, len(keys))}
	str := unsafe.String(unsafe.SliceData(buf), len(buf))
	for i := range keys {
		s.B[i] = buf[offs[i]:offs[i+1]:offs[i+1]]
		s.S[i] = str[offs[i]:offs[i+1]]
	}
	return s
}

// Kind names a key distribution.
type Kind string

const (
	// U64 are uniformly random 64-bit integers, encoded like FromUint64
	// (8 bytes, big endian).
	U64 Kind = "u64"
	// Str are path-like strings "tenant/category/word/NNNNN" with realistic
	// shared prefixes, about 20-35 bytes long (so mostly longer than 16).
	Str Kind = "str"
	// UUID are random UUIDs in lowercase hex with dashes (36 bytes): long keys
	// without shared prefixes.
	UUID Kind = "uuid"
	// Email are addresses "first.last@domain" over a few common domains,
	// about 25 bytes: the shared part comes last.
	Email Kind = "email"
	// URL are web addresses over three hosts with 2-4 path segments and a
	// query, about 66 bytes: long shared prefixes, differences at the end.
	URL Kind = "url"
	// Path are real file paths from the packages of Debian 12, about 65
	// bytes and up to 300 (testdata/paths.txt.gz, see cmd/mkcorpora).
	Path Kind = "path"
	// Street are German street names from OpenStreetMap, about 14 bytes, with
	// the localities that have a street of that name as their natural values
	// (testdata/streets.tsv.gz, see cmd/mkcorpora).
	Street Kind = "street"
)

// Kinds lists every key kind, the synthetic ones first.
var Kinds = []Kind{U64, Str, UUID, Email, URL, Path, Street}

// Text reports whether the keys of kind are UTF-8 text, which never contains
// the byte 0xFF, so that [p, p+0xFF...] holds exactly the keys that start
// with p (see PrefixEnd).
func Text(kind Kind) bool { return kind != U64 }

// Prefix returns what a user searches key by: for paths and URLs the
// directory it lies in (up to its last slash, as when listing a directory),
// for other text keys the first four characters, as typed into a search
// field.
func Prefix(kind Kind, key []byte) []byte {
	if kind == Path || kind == URL {
		if i := bytes.LastIndexByte(key, '/'); i >= 0 {
			return key[:i+1]
		}
	}
	return key[:min(4, len(key))]
}

// PrefixEnd returns the upper bound of the keys that start with p, for text
// keys (see Text): p followed by 0xFF bytes beyond any key's length.
func PrefixEnd(p []byte) []byte {
	return append(append([]byte(nil), p...), bytes.Repeat([]byte{0xFF}, maxKeyLen)...)
}

// maxKeyLen bounds the length of every key of every kind.
const maxKeyLen = 512

// Capacity returns the largest n that Generate accepts for kind: a corpus
// kind holds a fixed number of keys, of which Generate draws n keys and n
// misses. Synthetic kinds have no limit.
func Capacity(kind Kind) int {
	switch kind {
	case Path:
		return len(pathCorpus()) / 2
	case Street:
		return len(streetCorpus().names) / 2
	}
	return math.MaxInt
}

// Corpus holds the inserted keys and disjoint miss keys.
type Corpus struct {
	Keys   Set // in insertion order (random)
	Hits   Set // all keys again, in an independent random order
	Misses Set // same distribution, none of them present
	// Natural holds the values each key of Keys has in the real world, for
	// kinds that have them (Street: its localities, numbered from 1), else nil.
	Natural [][]uint64
}

// Generate builds a corpus of n keys (and n misses) deterministically from seed.
// For a corpus kind, n must not exceed Capacity(kind).
func Generate(kind Kind, n int, seed uint64) Corpus {
	rng := rtcompare.NewDPRNG(seed)
	switch kind {
	case Path:
		return fromList(pathCorpus(), nil, n, &rng)
	case Street:
		c := streetCorpus()
		return fromList(c.names, c.locs, n, &rng)
	}
	gen := generator(kind, &rng)
	seen := make(map[string]struct{}, 2*n)
	draw := func() [][]byte {
		out := make([][]byte, 0, n)
		for len(out) < n {
			k := gen()
			if _, dup := seen[string(k)]; dup {
				continue
			}
			seen[string(k)] = struct{}{}
			out = append(out, k)
		}
		return out
	}
	keys := draw()
	misses := draw()
	hits := append([][]byte(nil), keys...)
	shuffle(hits, &rng)
	return Corpus{Keys: Pack(keys), Hits: Pack(hits), Misses: Pack(misses)}
}

func generator(kind Kind, rng *rtcompare.DPRNG) func() []byte {
	switch kind {
	case U64:
		return func() []byte { return binary.BigEndian.AppendUint64(nil, rng.Uint64()) }
	case UUID:
		const hex = "0123456789abcdef"
		return func() []byte {
			k := make([]byte, 36)
			for i := range k {
				k[i] = hex[rng.Uint64()%16]
			}
			k[8], k[13], k[18], k[23] = '-', '-', '-', '-'
			return k
		}
	case Email:
		domains := []string{"gmail.com", "yahoo.com", "outlook.com", "example.org", "web.de"}
		return func() []byte {
			k := append(word(rng, 4+int(rng.Uint64()%8)), '.')
			k = append(append(k, word(rng, 3+int(rng.Uint64()%8))...), '@')
			return append(k, domains[rng.Uint64()%uint64(len(domains))]...)
		}
	case URL:
		hosts := []string{"https://www.example.com/", "https://shop.example.com/", "https://api.example.net/v2/"}
		segs := []string{"products", "category", "users", "orders", "items", "search", "images", "static"}
		return func() []byte {
			k := []byte(hosts[rng.Uint64()%3])
			for range 2 + rng.Uint64()%3 {
				k = append(append(k, segs[rng.Uint64()%uint64(len(segs))]...), '/')
			}
			k = strconv.AppendUint(k, rng.Uint64()%1_000_000, 10)
			k = append(k, "?ref="...)
			return strconv.AppendUint(k, rng.Uint64()%(1<<32), 16)
		}
	case Str:
		tenants := words(rng, 16)
		cats := words(rng, 256)
		items := words(rng, 1024)
		return func() []byte {
			k := make([]byte, 0, 40)
			k = append(k, tenants[rng.Uint64()%16]...)
			k = append(k, '/')
			k = append(k, cats[rng.Uint64()%256]...)
			k = append(k, '/')
			k = append(k, items[rng.Uint64()%1024]...)
			k = append(k, '/')
			return strconv.AppendUint(k, 10000+rng.Uint64()%90000, 10)
		}
	}
	panic(fmt.Sprintf("unknown key kind %q", kind))
}

// word returns n random lowercase letters.
func word(rng *rtcompare.DPRNG, n int) []byte {
	w := make([]byte, n)
	for j := range w {
		w[j] = 'a' + byte(rng.Uint64()%26)
	}
	return w
}

func words(rng *rtcompare.DPRNG, n int) []string {
	out := make([]string, n)
	for i := range out {
		w := make([]byte, 3+rng.Uint64()%6)
		for j := range w {
			w[j] = 'a' + byte(rng.Uint64()%26)
		}
		out[i] = string(w)
	}
	return out
}

func shuffle(k [][]byte, rng *rtcompare.DPRNG) {
	for i := len(k) - 1; i > 0; i-- {
		j := int(rng.Uint64() % uint64(i+1))
		k[i], k[j] = k[j], k[i]
	}
}

// Values assigns each of n keys a deterministic set of uint64 values with a
// skewed size distribution typical of multimap indexes: 50% of keys hold one
// value, 35% hold 2-4, 12% hold 5-16 and 3% hold 17-200 (about 6 on average).
// Values have the top bit clear, so a caller can make guaranteed-absent values
// by setting it. Key i's values are vals[offs[i]:offs[i+1]].
func Values(n int, seed uint64) (vals []uint64, offs []int) {
	rng := rtcompare.NewDPRNG(seed)
	offs = make([]int, n+1)
	for i := range n {
		var c uint64
		switch r := rng.Uint64() % 100; {
		case r < 50:
			c = 1
		case r < 85:
			c = 2 + rng.Uint64()%3
		case r < 97:
			c = 5 + rng.Uint64()%12
		default:
			c = 17 + rng.Uint64()%184
		}
		for range c {
			vals = append(vals, rng.Uint64()>>1)
		}
		offs[i+1] = len(vals)
	}
	return vals, offs
}

// Sorted returns the keys of s in ascending byte order, packed.
func Sorted(s Set) Set {
	k := append([][]byte(nil), s.B...)
	slices.SortFunc(k, bytes.Compare)
	return Pack(k)
}
