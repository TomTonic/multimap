# Typed leaves for values with a pointer (2026-10-01)

Values that hold a pointer, strings above all, lived in set leaves: a
64-byte `vset.Set` after a key area of a fixed size. Typed leaves (`40ca08f`)
keep the key and 1 to 16 values in one object, allocated with its real type so
that the garbage collector sees the pointers; see `internal/art/typed.go`. The
growth was tuned afterwards (`373b33e`): a key with a second value gets room
for four at once.

This is `node-layout` against `main` (`818e861`) with string values
(`-tags strvals`), on native Windows, Ryzen 9 7900, with rtcompare v0.8.0.
The first runs are against `40ca08f`, the follow-up against `373b33e`. The
earlier figures of the same setup, before typed leaves, are in
[`../consolidation-vs-main`](../consolidation-vs-main/README.md).

## Results, serial (`serial/`)

The dev suite, 4K, 16K and 256K keys, a first stage of 6 processes and at
most 12, memory at 1M keys, 1 h 19 min (06:12-07:30). In brackets the figure
before typed leaves. Each cell gives the median over the key kinds; `*n`
marks imprecise comparisons (31 of 182 in all).

```sh
bench -suite dev -vs baseline -sizes 4096,16384,262144 -minprocs 6 -maxprocs 12 -memn 1048576 -memrounds 3
```

| values | operation | 4K | 16K | 256K |
|---|---|---|---|---|
| multi-str | `valuesFor` | 1.00 (0.98) | 1.01 (0.99) | 1.11 (1.06) *3 |
| multi-str | `valuesBetween` | 0.97 (0.92) *4 | 0.99 (0.93) | 1.11 (0.98) *2 |
| multi-str | `prefix` | 0.96 (0.93) | 0.98 (0.94) | 1.11 (0.98) *2 |
| multi-str | `churn` | 0.92 (0.99) | 0.95 (0.98) | 0.94 (0.95) *5 |
| multi-str | `build` | 0.90 (0.97) | 0.91 (0.98) | |
| unique-str | `valuesFor` | 1.01 (0.96) | 1.03 (0.96) | 1.17 (1.04) *2 |
| unique-str | `valuesBetween` | 0.92 (0.82) | 0.97 (0.85) | 1.19 (0.97) |
| unique-str | `prefix` | 0.94 (0.90) | 0.96 (0.90) | 1.12 (0.99) *2 |
| unique-str | `churn` | 0.99 (0.96) | 1.03 (0.96) | 1.07 (1.01) *3 |
| unique-str | `build` | 0.94 (0.92) | 0.98 (0.94) | |

**Memory** at 1M keys, heap bytes per key, against `main`:

| values | u64 | str | uuid | email | url | path | street |
|---|---|---|---|---|---|---|---|
| multi-str | −26% | −28% | −22% | −24% | −27% | −33% | −38% |
| unique-str | −61% | −58% | −44% | −49% | −47% | −58% | −54% |

(before typed leaves: unique-str ±0 to −28%, multi-str 0 to −16%)

## Results, parallel (`parallel/`)

256K and 1M keys, kinds `u64`, `uuid`, `url`, `path` (no `path` at 1M), 8
processes at a time, 50 minutes (07:31-08:21; a first stage of one wave of
8, at most 160 processes); all 52 comparisons precise. Not comparable with
the serial figures.

| values | operation | 256K | 1M |
|---|---|---|---|
| multi-str | `valuesFor` | 1.08 (1.04) | 1.05 (1.01) |
| multi-str | `valuesBetween` | 1.10 (0.97) | 1.06 (0.96) |
| multi-str | `prefix` | 1.10 (0.98) | 1.10 (0.97) |
| multi-str | `churn` | 0.96 (1.00) | 0.96 (0.98) |
| unique-str | `valuesFor` | 1.11 (1.06) | 1.08 (1.01) |
| unique-str | `valuesBetween` | 1.15 (0.96) | 1.12 (0.95) |
| unique-str | `prefix` | 1.12 (0.96) | 1.12 (0.94) |
| unique-str | `churn` | 1.08 (1.02) | 1.05 (1.00) |

## Follow-up after the growth change (`growth-fix/`)

`churn` and `build` for 4K and 16K keys, the kinds `u64`, `uuid`, `url`,
`path` and `street`, serial, against `373b33e`:

| values | operation | 4K | 16K |
|---|---|---|---|
| multi-str | `churn` | 0.94 (0.92) | 0.97 (0.95) |
| multi-str | `build` | 0.92 (0.89) | 0.93 (0.91) |
| unique-str | `churn` | 1.00 (0.99) | 1.03 (1.03) |
| unique-str | `build` | 0.94 (0.94) | 0.98 (0.99) |

## Findings

1. **Typed leaves halve the memory of string values** against `main`:
   unique-str 44-61% less, multi-str 22-38% less.
2. **They close the range gap.** `valuesBetween` and `prefix` were 0.82-0.98
   at every size before, and are 0.92-0.99 at 4K and 16K and 1.10-1.19 from 256K
   up, as fast as with `uint64` values or faster. Lookups are at 1.00-1.03 from
   4K and 1.05-1.17 from 256K.
3. **The price is the several-valued key.** A key that gets a second and a
   third value moves from a leaf for one value into a leaf for four, an
   allocation the garbage collector has to scan, where `main` held three
   values inline in its set leaf. `build` and `churn` with several values per
   key are 3-8% slower than `main` at 4K and 16K (build 0.92-0.93, churn
   0.94-0.97 after the growth change; 0.89-0.92 before). A microbenchmark of
   building 4K keys with three strings each takes 31% longer than with set
   leaves after the change (60% before) and allocates 2.4 objects per key
   instead of 1.4.
4. **Single-valued keys gain from it**: building with one value per key is
   16% faster than with set leaves, with 40% less memory.
