# The page tree for string values on the Apple M1 Pro (jobs l1 to l4, 2026-10-04)

Jobs `l1` to `l4` of `bench/remote/queue.txt`, run by the user on 2026-10-04 between 08:06 and 10:51 on the
Apple M1 Pro (8 performance cores, 128-byte cache lines), at commit `c4dc4b6` (the experiment of
`internal/artstr`, now parked on the branch `mkmv-experiment`). Raw files on the branch `arm-results`;
`env.txt`, `args.txt` and the two summaries of each job are here. The profiles are named as at the time: `unique-str`
is the single-value profile, `multi-str` the natural mix. `l3` started with the machine at load 16 (the
runner warned); the other jobs at 4 to 8. The tables below are made from the summaries: speed of a candidate
as many times as fast as today's `ordered` (above 1: faster; `*`: interval wider than asked; `?`: not
resolved). Memory tables are in the summaries.

Reading, in one paragraph: the M1 repeats the PC result. The multi-key page with one value per entry gives, on
the natural mix, today's tree (0.9 to 1.0 in every cell; `l3`, `l4`): the fall back takes the pages away. With several values
per entry (`mv`) the memory falls to 88 B/key on `street` (the PC: 88), the bytes the garbage collector scans to 44
(PC 44), and GC per cycle from +42 to +9 ms; lookups at 4K to 16K keys are 0.51 to 0.62 of today's tree
(PC 0.49 to 0.65) and equal at the full corpus (0.87 to 0.89, PC 0.92 to 1.08), range scans 1.28 to 1.61 on `street`
(PC 1.05 to 1.56), `build` 0.43 to 0.54 (PC 0.44 to 0.57). With one value per entry (`l1`, `l2`) the pages are 30
B/key instead of 67 on `street`, ranges 1.3 to 2.3 times today's tree (PC 1.6 to 2.5), lookups 0.5 to 0.6 at
4K and 16K, equal from 86K keys on; today's tree is 0.6 to 0.75 of `btree-map` on ranges on the M1 (the PC:
0.28 to 0.39). Nothing here changes the course set on 2026-10-04.

## Memory (B/key, the string bytes inside only for the page tree)

`l1`:

| values | keys | n | candidate | rounds | heap B/key | scannable B/key | GC CPU per cycle | heap B/key after removing half the keys |
|---|---|---:|---|---:|---:|---:|---:|---:|
| unique-str | street | 212449 | ordered | 3 | 67 | 58 | +31 ms | 35 |
| unique-str | street | 212449 | btree-map | 3 | 66 | 49 | +18 ms | 34 |
| unique-str | street | 212449 | ordered-lpage | 3 | 30 | 4 | +1 ms | 20 |
| unique-str | street | 212449 | ordered-lpage-zc | 3 | 30 | 4 | +1 ms | 20 |

`l2`:

| values | keys | n | candidate | rounds | heap B/key | scannable B/key | GC CPU per cycle | heap B/key after removing half the keys |
|---|---|---:|---|---:|---:|---:|---:|---:|
| unique-str | dirs | 86215 | ordered | 3 | 78 | 66 | +13 ms | 41 |
| unique-str | dirs | 86215 | btree-map | 3 | 107 | 49 | +6 ms | 55 |
| unique-str | dirs | 86215 | ordered-lpage | 3 | 54 | 9 | +1 ms | 37 |
| unique-str | dirs | 86215 | ordered-lpage-zc | 3 | 54 | 9 | +1 ms | 36 |

`l3`:

| values | keys | n | candidate | rounds | heap B/key | scannable B/key | GC CPU per cycle | heap B/key after removing half the keys |
|---|---|---:|---|---:|---:|---:|---:|---:|
| multi-str | street | 212449 | ordered | 3 | 119 | 108 | +42 ms | 61 |
| multi-str | street | 212449 | btree-sets | 3 | 367 | 302 | +66 ms | 185 |
| multi-str | street | 212449 | ordered-lpage | 3 | 119 | 108 | +43 ms | 61 |
| multi-str | street | 212449 | ordered-lpage-mv | 3 | 88 | 44 | +9 ms | 50 |
| multi-str | street | 212449 | ordered-lpage-mvzc | 3 | 88 | 44 | +8 ms | 50 |

`l4`:

| values | keys | n | candidate | rounds | heap B/key | scannable B/key | GC CPU per cycle | heap B/key after removing half the keys |
|---|---|---:|---|---:|---:|---:|---:|---:|
| multi-str | dirs | 86215 | ordered | 3 | 149 | 135 | +20 ms | 75 |
| multi-str | dirs | 86215 | btree-sets | 3 | 421 | 315 | +25 ms | 209 |
| multi-str | dirs | 86215 | ordered-lpage | 3 | 149 | 135 | +20 ms | 75 |
| multi-str | dirs | 86215 | ordered-lpage-mv | 3 | 139 | 59 | +6 ms | 79 |
| multi-str | dirs | 86215 | ordered-lpage-mvzc | 3 | 139 | 59 | +6 ms | 79 |

## Speed

```

l1: unique-str street: ordered as fast as btree-map (above 1: ordered faster)
| op | 4096 | 16384 | 212449 |
| valuesFor | 1.87 | 1.98 | 1.50 |
| valuesBetween | 0.64 | 0.64 | 0.62 |
| prefix | 0.86 | 0.77 | 0.60 |
| churn | 1.17 | 1.17 | 1.08 |
| build | 1.26 | 1.31 | - |

l1: unique-str street: ordered-lpage as fast as ordered (above 1: ordered-lpage faster)
| op | 4096 | 16384 | 212449 |
| valuesFor | 0.53 | 0.60 | 1.02 |
| valuesBetween | 1.89 | 1.85 | 2.33 |
| prefix | 1.04 | 1.69 | 2.78 |
| churn | 0.66 | 0.75 | 1.18* |
| build | 0.55 | 0.61 | - |

l1: unique-str street: ordered-lpage-zc as fast as ordered (above 1: ordered-lpage-zc faster)
| op | 4096 | 16384 | 212449 |
| valuesFor | 0.62 | 0.68 | 1.10 |
| valuesBetween | 2.33 | 2.38 | 2.86 |
| prefix | 1.20 | 2.08 | 3.70 |
| churn | 0.52 | 0.65 | 1.04* |
| build | 0.44 | 0.48 | - |

l2: unique-str dirs: ordered as fast as btree-map (above 1: ordered faster)
| op | 4096 | 16384 | 86215 |
| valuesFor | 1.26 | 1.39 | 1.18 |
| valuesBetween | 0.72 | 0.73 | 0.75 |
| prefix | 0.75 | 0.71 | 0.68 |
| churn | 0.87 | 0.89 | 0.97 |
| build | 0.91 | 0.95 | - |

l2: unique-str dirs: ordered-lpage as fast as ordered (above 1: ordered-lpage faster)
| op | 4096 | 16384 | 86215 |
| valuesFor | 0.61 | 0.65 | 0.85* |
| valuesBetween | 1.33 | 1.37 | 1.59 |
| prefix | 1.37 | 1.56 | 1.75 |
| churn | 0.75 | 0.84 | 0.97 |
| build | 0.59 | 0.64 | - |

l2: unique-str dirs: ordered-lpage-zc as fast as ordered (above 1: ordered-lpage-zc faster)
| op | 4096 | 16384 | 86215 |
| valuesFor | 0.68 | 0.71 | 0.88* |
| valuesBetween | 1.75 | 1.79 | 1.92 |
| prefix | 1.85 | 2.22 | 2.44 |
| churn | 0.64 | 0.75 | 0.90 |
| build | 0.52 | 0.56 | - |


l3: multi-str street: ordered as fast as btree-sets (above 1: ordered faster)
| op | 4096 | 16384 | 212449 |
| valuesFor | 2.33 | 2.39 | 2.19 |
| valuesBetween | 2.73 | 2.70 | 2.86 |
| prefix | 2.89 | 3.02 | 3.10 |
| churn | 1.81 | 1.66 | 1.38 |
| build | 1.85 | 1.83 | - |

l3: multi-str street: ordered-lpage as fast as ordered (above 1: ordered-lpage faster)
| op | 4096 | 16384 | 212449 |
| valuesFor | 0.99 | 1.00? | 0.99? |
| valuesBetween | 1.00? | 1.01 | 1.00? |
| prefix | 0.98 | 1.00? | 0.98? |
| churn | 1.00 | 1.00? | 1.19 |
| build | 0.99? | 0.91 | - |

l3: multi-str street: ordered-lpage-mv as fast as ordered (above 1: ordered-lpage-mv faster)
| op | 4096 | 16384 | 212449 |
| valuesFor | 0.51 | 0.56 | 0.89 |
| valuesBetween | 1.28 | 1.30 | 1.61 |
| prefix | 0.91 | 1.23 | 1.75 |
| churn | 0.47 | 0.57 | 1.01? |
| build | 0.43 | 0.51 | - |

l3: multi-str street: ordered-lpage-mvzc as fast as ordered (above 1: ordered-lpage-mvzc faster)
| op | 4096 | 16384 | 212449 |
| valuesFor | 0.59 | 0.64 | 0.97 |
| valuesBetween | 1.61 | 1.59 | 1.92 |
| prefix | 1.08 | 1.52 | 2.17 |
| churn | 0.38 | 0.51 | 0.90 |
| build | 0.36 | 0.41 | - |

l4: multi-str dirs: ordered as fast as btree-sets (above 1: ordered faster)
| op | 4096 | 16384 | 86215 |
| valuesFor | 1.62 | 1.65 | 1.61 |
| valuesBetween | 3.11 | 3.16 | 3.13 |
| prefix | 3.73 | 3.49 | 3.92 |
| churn | 1.33 | 1.27 | 1.17 |
| build | 1.36 | 1.39 | - |

l4: multi-str dirs: ordered-lpage as fast as ordered (above 1: ordered-lpage faster)
| op | 4096 | 16384 | 86215 |
| valuesFor | 1.01 | 1.00? | 0.98 |
| valuesBetween | 1.00? | 1.01 | 1.00? |
| prefix | 0.99? | 1.00? | 0.99? |
| churn | 1.00? | 1.01? | 1.18 |
| build | 0.98 | 0.93 | - |

l4: multi-str dirs: ordered-lpage-mv as fast as ordered (above 1: ordered-lpage-mv faster)
| op | 4096 | 16384 | 86215 |
| valuesFor | 0.58 | 0.62 | 0.87 |
| valuesBetween | 0.98 | 1.05 | 1.22 |
| prefix | 0.90 | 1.11 | 1.32 |
| churn | 0.55 | 0.67 | 0.94 |
| build | 0.49 | 0.54 | - |

l4: multi-str dirs: ordered-lpage-mvzc as fast as ordered (above 1: ordered-lpage-mvzc faster)
| op | 4096 | 16384 | 86215 |
| valuesFor | 0.64 | 0.68 | 0.92 |
| valuesBetween | 1.28 | 1.35 | 1.52 |
| prefix | 1.25 | 1.49 | 1.82 |
| churn | 0.47 | 0.60 | 0.86 |
| build | 0.42 | 0.46 | - |

```
