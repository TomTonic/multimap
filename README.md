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

Numeric keys

- You can build keys from integers using the `FromInt*` / `FromUint*` helpers.
- Integer keys are encoded as fixed-width big-endian byte sequences (most-significant
	byte first). This makes unsigned integer keys sort in the natural numeric order when
	compared lexicographically.
- Signed integers are stored using their raw two's-complement bit pattern. Because
	lexicographic byte-wise comparison treats the bytes as unsigned values, ordering of
	signed keys does not correspond to mathematical order (negative values will not
	necessarily compare "less than" positive values). If you need a lexicographic
	representation that preserves signed numeric ordering, convert/encode the value
	accordingly before creating a `Key`.

Behavior and semantics

- `PutValue(key, v)` clones `key` before inserting; mutating the caller's `Key` after
	calling `PutValue` will not affect the stored key.
- `GetValuesFor(key)` and other getters return clones of stored `Set3` instances;
	modifying the returned set does not affect the `MultiMap`'s contents.
- `Size()` returns the number of keys stored (not the total number of values).

Examples

See the `example_test.go` in this package for runnable examples that also appear
in generated GoDoc.

