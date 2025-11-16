# multi_map

[![Go Report Card](https://goreportcard.com/badge/github.com/TomTonic/multi_map)](https://goreportcard.com/report/github.com/TomTonic/multi_map)
[![Go Reference](https://pkg.go.dev/badge/github.com/TomTonic/multi_map.svg)](https://pkg.go.dev/github.com/TomTonic/multi_map)
[![Linter](https://github.com/TomTonic/multi_map/actions/workflows/linter.yml/badge.svg)](https://github.com/TomTonic/multi_map/actions/workflows/linter.yml)
[![Tests](https://github.com/TomTonic/multi_map/actions/workflows/coverage.yml/badge.svg?branch=main)](https://github.com/TomTonic/multi_map/actions/workflows/coverage.yml)
![coverage](https://raw.githubusercontent.com/TomTonic/multi_map/badges/.badges/main/coverage.svg)

`multi_map` provides a compact, thread-safe multi-map keyed by `Key`.

Highlights
- Thread-safe operations for concurrent use by multiple goroutines.
- Keys are represented by the `Key` type (alias for `[]byte`).
- Values for each key are stored in a `Set3` (from `github.com/TomTonic/Set3`).

Getting started

Install dependencies and run tests:

```bash
go mod tidy
go test ./... -v
```

String keys

- Use `FromString(s)` to convert a string to a `Key`. `FromString` normalizes the
		input string using Unicode NFC so canonically equivalent strings map to the same key.
- `Key` ordering (`LessThan`) is a byte-wise lexicographic comparison of the UTF-8 bytes —
		it is neither locale-aware nor rune-aware.

**Numeric keys**

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

Behavior and semantics

- `PutValue(key, v)` clones `key` before inserting; mutating the caller's `Key` after
	calling `PutValue` will not affect the stored key.
- `GetValuesFor(key)` and other getters return clones of stored `Set3` instances;
	modifying the returned set does not affect the `MultiMap`'s contents.
- `Size()` returns the number of keys stored (not the total number of values).

Examples

See the `example_test.go` in this package for runnable examples that also appear
in generated GoDoc.

