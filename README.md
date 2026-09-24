# multimap for Go

[![Go Report Card](https://goreportcard.com/badge/github.com/TomTonic/multimap)](https://goreportcard.com/report/github.com/TomTonic/multimap)
[![Go Reference](https://pkg.go.dev/badge/github.com/TomTonic/multimap.svg)](https://pkg.go.dev/github.com/TomTonic/multimap)
[![Linter](https://github.com/TomTonic/multimap/actions/workflows/linter.yml/badge.svg)](https://github.com/TomTonic/multimap/actions/workflows/linter.yml)
[![Tests](https://github.com/TomTonic/multimap/actions/workflows/coverage.yml/badge.svg?branch=main)](https://github.com/TomTonic/multimap/actions/workflows/coverage.yml)
![coverage](https://raw.githubusercontent.com/TomTonic/multimap/badges/.badges/main/coverage.svg)

`multimap` provides fast, compact multimaps for Go.
A multimap is a data structure that allows multiple values to be associated with a single key, unlike a regular map where each key has exactly one value.

## Implementations

| Constructor | Backed by | Range queries | Iteration order | Concurrency |
|-------------|-----------|---------------|-----------------|-------------|
| `New[T]()` | adaptive radix tree | visit only the keys in range | ascending key order | safe (read/write lock) |
| `NewOrdered[T]()` | adaptive radix tree | visit only the keys in range | ascending key order | not synchronized |
| `NewHashed[T]()` | Go map | scan all keys | unspecified | not synchronized |
| `Synchronized[T](m)` | wraps any of the above | as the wrapped map | as the wrapped map | safe (read/write lock) |

Choose `Ordered` (the default behind `New`) when range queries or ordered
iteration matter. Choose `Hashed` when they are rare: point operations on long
string keys are faster there, but every range query scans all keys, and like
every Go map it keeps its memory after deletions. The benchmarks behind these
statements are in [bench/README.md](bench/README.md).

Both keep the values of a key in a compact container: up to three values
inline, then a plain array, then a hash set.

## Concurrency

`Ordered` and `Hashed` are not synchronized. Any number of goroutines may read
at the same time, but a write must not run concurrently with any other
access; provide that exclusion yourself, or wrap the multimap with
`Synchronized`, which guards every method with a `sync.RWMutex` (reads run in
parallel, writes are exclusive). `New` returns a synchronized `Ordered`
multimap.

## Sets and iterators

Every read method exists twice:

- The plain form (`ValuesFor`, `ValuesBetweenInclusive`, `AllKeys`, ...)
  returns an independent copy (`*set3.Set3[T]` or `[]Key`). It allocates, but
  you may keep and modify the result, and on a synchronized multimap no lock is
  held once it returns.
- The `Seq` form (`ValuesForSeq`, `ValuesBetweenInclusiveSeq`, `AllKeysSeq`,
  ...) returns an `iter.Seq` over the stored data without copying. It yields
  the values key by key, so a value stored under several keys is yielded once
  per key. Keys yielded by `AllKeysSeq` must not be modified or retained, and
  the loop body must not modify the multimap it iterates over.

On a synchronized multimap, a `Seq` iterator holds the read lock for the whole
loop. The loop body must not call any method of the same multimap (this can
deadlock), and a long loop delays writers. Use the plain form when the loop
body needs the multimap or runs long.

## Key Characteristics

- One-to-many mapping: Each key can have zero, one, or multiple values
- Duplicate values: The same value can only be stored once per key but it can be stored multiple times for different keys
- Key uniqueness: Keys themselves are still unique - there's only one entry per key

## Common Use Cases

- Indexing: Group items by category (e.g., products by brand)
- Graph representations: Store adjacency lists where each node maps to multiple connected nodes
- HTTP headers: Multiple values for the same header name
- Database indexes: Multiple records sharing the same indexed value

## Indexing

The multimap uses `Key` objects for indexing, which are internally represented as `[]byte` arrays. This allows for efficient storage and comparison regardless of the original data type used to create the key.

### Key behavior

- **Mixed types**: A single multimap can contain keys created from different data types (strings, integers, custom byte arrays) without any issues.
- **Lexicographic ordering**: All keys are compared using byte-wise lexicographic ordering of their internal `[]byte` representation.
- **Custom keys**: You can create keys from any `[]byte` array using the generic constructor for your own implementations.

**Important**: When mixing different key types in range queries, the lexicographic ordering may be counter-intuitive. For example:

```go
// These keys will NOT be ordered numerically when mixed with strings:
key1 := FromString("100")     // UTF-8 bytes: [0x31, 0x30, 0x30]
key2 := FromInt64(50)         // Encoded as: [0x80, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x32]
key3 := FromString("25")      // UTF-8 bytes: [0x32, 0x35]

// Lexicographic order: key1 ("100") < key3 ("25") < key2 (50)
// This differs from intuitive numeric/alphabetic ordering!
```

For predictable range query behavior, use consistent key types within logical ranges. See the type-specific encoding details below.

### String keys

- Use `FromString(s)` to convert a string to a `Key`. `FromString` normalizes the
		input string using Unicode NFC so canonically equivalent strings map to the same key.
- `Key` ordering (`LessThan`) is a byte-wise lexicographic comparison of the UTF-8 bytes —
		it is neither locale-aware nor rune-aware.

### Numeric keys

- Use `FromInt*` / `FromUint*` helpers to build integer keys.
- All integer keys are encoded as fixed-width 8-byte big-endian sequences (MSB first).
- To ensure consistent ordering across signed/unsigned types and across widths, all
	integer constructors add an offset of `1<<63` before encoding. Signed values are
	converted to `int64` first; unsigned values are treated as `uint64` and the same
	offset is added. This maps signed and unsigned values into a single uint64 namespace
	so that lexicographic comparison matches numeric ordering.

Examples / consequences:
- `FromInt64(0)` equals `FromUint64(0)` because both are encoded as `0 + (1<<63)`.
- The smallest `int64` (`-2^63`) maps to `[00,00,00,00,00,00,00,00]` after encoding;
    negative signed values compare before zero/positive values as expected for numeric order.
- Values encoded from different widths are comparable — for example `FromInt32(x)` is
	identical to `FromInt64(x)` for the same numeric `x`.

## Behavior and semantics

- `AddValue(key, v)` clones `key` before inserting; mutating the caller's `Key` after
    calling `AddValue` will not affect the stored key.
- `ValuesFor(key)` and the other set-returning methods return copies;
    modifying a returned set does not affect the `MultiMap`'s contents.
- Removing the last value of a key removes the key.
- **Range queries**: Keys are ordered in lexicographic order, allowing efficient range
	queries between two key boundaries. Range operations return a set of all values of
	all keys where the key falls within the specified range (inclusive or exclusive based
	on the method). The ordering follows the byte-wise comparison rules described above
	for string and numeric keys.

## Examples

See the `example_test.go` in this package for runnable examples that also appear
in generated GoDoc.

