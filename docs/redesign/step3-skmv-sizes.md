# Step 3: what a page per key holds (the size class question)

Question of 2026-10-04: if every key is a single-key page (SKMV) of its own, in which size classes?
The plan's grid is 128, 256, 384 and 512 bytes (R1). This is the model that shows what that costs on the
real data. Tool: `bench/cmd/skmodel` (`go run ./cmd/skmodel -detail`, a second, no tree is built).

**Model.** The keys of the corpus in key order. A tree of byte nodes with lazy expansion hangs the page of
a key at the first byte that tells it from both neighbours; the page holds the key from the next byte on
(the remainder). A page is: 4 bytes of header (kind, number of values, length of the remainder as two
bytes), the remainder, then each value as a length byte and its bytes. "Content" is that, without padding.
The values are the real ones (`street`: localities, `dirs`: file names). Nodes are not counted. A page whose
content is above 512 bytes is a full page of 512 with a value set behind it (value overflow); the bytes of the
value set are not counted here.

Summary (bytes per key, the page only; header of 4 bytes in this table, the design note uses 3, `-header` sets it):

| data | values | keys | remainder | values a key | value bytes | content | page in the grid 128..512 | page in Go's classes | content over 512 B | content up to 64 B |
|---|---|--:|--:|--:|--:|--:|--:|--:|--:|--:|
| street | multi | 212449 | 5.2 | 2.73 | 27.6 | 39.5 | 132.9 | 38.7 | 0.51 % | 95 % |
| street | unique | 212449 | 5.2 | 1.00 | 9.9 | 20.1 | 128.0 | 27.9 | 0.00 % | 100 % |
| dirs | multi | 86215 | 8.9 | 3.44 | 67.9 | 84.3 | 144.4 | 64.9 | 1.61 % | 81 % |
| dirs | unique | 86215 | 8.9 | 1.00 | 15.4 | 29.3 | 128.0 | 37.0 | 0.00 % | 98 % |

For comparison, measured: the whole tree of today (nodes, pages and leaves, strings not counted) takes 119 B/key on `street` and 149 on `dirs` in the natural mix, 67 and 78 with one value per key (step3-tree-pages.md).

## Distribution and examples

How the content spreads over the size classes, and keys at quantiles of the content size.

street, multi: content of the page of a key

| content bytes | keys | share | content B/key | grid 128..512 B/key | Go classes B/key |
|---|--:|--:|--:|--:|--:|
| 1 to 16 | 42948 | 20.2 % | 14.3 | 128 | 16.0 |
| 17 to 32 | 134702 | 63.4 % | 22.4 | 128 | 32.0 |
| 33 to 48 | 19284 | 9.1 % | 38.2 | 128 | 48.0 |
| 49 to 64 | 5286 | 2.5 % | 55.4 | 128 | 64.0 |
| 65 to 96 | 3952 | 1.9 % | 77.7 | 128 | 85.9 |
| 97 to 128 | 1706 | 0.8 % | 110.8 | 128 | 118.7 |
| 129 to 192 | 1624 | 0.8 % | 156.3 | 256 | 164.1 |
| 193 to 256 | 771 | 0.4 % | 221.3 | 256 | 229.0 |
| 257 to 384 | 709 | 0.3 % | 314.5 | 384 | 330.7 |
| 385 to 512 | 383 | 0.2 % | 444.2 | 512 | 460.4 |
| 513 to more | 1084 | 0.5 % | 2221.0 | 512 | 512.0 |

examples (key; remainder bytes; values; content bytes):

- p5: "Schöllentrup"; remainder 3 B; ["Lemgo"]; content 13 B
- p25: "Rostergeweg"; remainder 4 B; ["Dortmund"]; content 17 B
- p50: "Schlosspark Wendorf"; remainder 4 B; ["Möllenhagen"]; content 21 B
- p75: "Exerzierhügel"; remainder 10 B; ["Kümmersbruck"]; content 28 B
- p90: "Rottalblick"; remainder 4 B; ["Achstetten" "Großerlach" "Schwendi"]; content 40 B
- p99: "Flandernstr."; remainder 3 B; ["Albstadt" "Augsburg" "Erkelenz" "Essen" "Esslingen am Neckar" "Geilenkirchen" "... (22 values)"]; content 263 B

street, unique: content of the page of a key

| content bytes | keys | share | content B/key | grid 128..512 B/key | Go classes B/key |
|---|--:|--:|--:|--:|--:|
| 1 to 16 | 60943 | 28.7 % | 14.2 | 128 | 16.0 |
| 17 to 32 | 144463 | 68.0 % | 21.8 | 128 | 32.0 |
| 33 to 48 | 6941 | 3.3 % | 35.9 | 128 | 48.0 |
| 49 to 64 | 91 | 0.0 % | 52.8 | 128 | 64.0 |
| 65 to 96 | 11 | 0.0 % | 74.8 | 128 | 84.4 |

examples (key; remainder bytes; values; content bytes):

- p5: "Aureliusplatz"; remainder 4 B; ["Calw"]; content 13 B
- p25: "Kellinghusener Weg"; remainder 2 B; ["Lockstedt"]; content 16 B
- p50: "Naundorfer Mühle"; remainder 5 B; ["Delitzsch"]; content 19 B
- p75: "Memmelers Wiese"; remainder 8 B; ["Pfullingen"]; content 23 B
- p90: "Brethausstr."; remainder 7 B; ["Lauter-Bernsbach"]; content 28 B
- p99: "Eugenie-von-Soden-Str."; remainder 13 B; ["Esslingen am Neckar"]; content 37 B

dirs, multi: content of the page of a key

| content bytes | keys | share | content B/key | grid 128..512 B/key | Go classes B/key |
|---|--:|--:|--:|--:|--:|
| 1 to 16 | 5492 | 6.4 % | 14.1 | 128 | 16.0 |
| 17 to 32 | 34876 | 40.5 % | 24.6 | 128 | 32.0 |
| 33 to 48 | 20668 | 24.0 % | 39.3 | 128 | 48.0 |
| 49 to 64 | 9033 | 10.5 % | 55.4 | 128 | 64.0 |
| 65 to 96 | 7119 | 8.3 % | 77.6 | 128 | 85.8 |
| 97 to 128 | 2765 | 3.2 % | 110.5 | 128 | 118.5 |
| 129 to 192 | 2356 | 2.7 % | 155.2 | 256 | 163.1 |
| 193 to 256 | 1020 | 1.2 % | 220.5 | 256 | 228.3 |
| 257 to 384 | 988 | 1.1 % | 311.8 | 384 | 328.2 |
| 385 to 512 | 512 | 0.6 % | 443.8 | 512 | 460.4 |
| 513 to more | 1386 | 1.6 % | 2188.4 | 512 | 512.0 |

examples (key; remainder bytes; values; content bytes):

- p5: "/usr/lib/python3/dist-packages/past/"; remainder 0 B; ["__init__.py"]; content 16 B
- p25: "/usr/share/doc/openvswitch-doc/html/_sources/faq/"; remainder 3 B; ["releases.rst.txt"]; content 24 B
- p50: "/usr/share/doc/librust-autocfg-dev/"; remainder 10 B; ["changelog.Debian.gz"]; content 34 B
- p75: "/usr/share/php/League/CommonMark/Extension/DescriptionList/Event/"; remainder 21 B; ["LooseDescriptionHandler.php"]; content 53 B
- p90: "/usr/share/doc/liblgooddatepicker-java/api/com/github/lgooddatepicker/components/"; remainder 10 B; ["CalendarPanelBeanInfo.html" "TimePickerSettings.TimeIncrement.html" "package-summary.html"]; content 100 B
- p99: "/usr/share/kicad/footprints/Connector_PCBEdge.pretty/"; remainder 13 B; ["BUS_AT.kicad_mod" "BUS_PCIexpress_x1.kicad_mod" "Samtec_MECF-05-01-L-DV-WT_2x05_P1.27mm_Polarized_Socket_Horizontal.kicad_mod" "Samtec_MECF-05-02-L-DV-WT_2x05_P1.27mm_Polarized_Socket_Horizontal.kicad_mod" "Samtec_MECF-05-02-NP-L-DV-WT_2x05_P1.27mm_Socket_Horizontal.kicad_mod" "Samtec_MECF-05-02-NP-L-DV_2x05_P1.27mm_Socket_Horizontal.kicad_mod" "... (12 values)"]; content 789 B

dirs, unique: content of the page of a key

| content bytes | keys | share | content B/key | grid 128..512 B/key | Go classes B/key |
|---|--:|--:|--:|--:|--:|
| 1 to 16 | 8510 | 9.9 % | 14.1 | 128 | 16.0 |
| 17 to 32 | 50723 | 58.8 % | 24.3 | 128 | 32.0 |
| 33 to 48 | 20759 | 24.1 % | 38.7 | 128 | 48.0 |
| 49 to 64 | 4653 | 5.4 % | 54.4 | 128 | 64.0 |
| 65 to 96 | 1473 | 1.7 % | 74.5 | 128 | 83.5 |
| 97 to 128 | 86 | 0.1 % | 105.8 | 128 | 115.5 |
| 129 to 192 | 11 | 0.0 % | 144.9 | 256 | 151.3 |

examples (key; remainder bytes; values; content bytes):

- p5: "/usr/include/iceoryx/v2.0.3/iceoryx_hoofs/internal/concurrent/"; remainder 0 B; ["loffli.hpp"]; content 15 B
- p25: "/usr/lib/haskell-packages/ghc/lib/x86_64-linux-ghc-9.0.2/bloomfilter-2.0.1.0-6Wb4Gh0qlqF68sl55KSOyv/include/"; remainder 7 B; ["lookup3.h"]; content 21 B
- p50: "/usr/share/doc/libbiojava-java/apidocs/src-html/org/biojava/utils/cache/"; remainder 5 B; ["WeakCacheMap.html"]; content 27 B
- p75: "/usr/share/doc/php-cache-integration-tests/"; remainder 21 B; ["README.md"]; content 35 B
- p90: "/usr/share/doc/libsimple-http-java/api/jquery/images/"; remainder 13 B; ["ui-icons_222222_256x240.png"]; content 45 B
- p99: "/usr/share/gnuradio/modtool/templates/gr-newmod/python/howto/bindings/"; remainder 49 B; ["python_bindings.cc"]; content 72 B

## Sets of size classes (`go run ./cmd/skmodel -sets`)

Same model, what a page per key takes with other size classes. "Content" is capped at 512 bytes here (the
rest is a value set), "fill" is content over page bytes.


street, multi: sets of size classes

| size classes | page B/key | content B/key | fill | keys per class |
|---|--:|--:|--:|---|
| 128, 256, 384, 512 (plan) | 132.9 | 30.7 | 23 % | 128: 97.8 %, 256: 1.1 %, 384: 0.3 %, 512: 0.7 % |
| 64 + plan | 72.0 | 30.7 | 43 % | 64: 95.2 %, 128: 2.7 %, 256: 1.1 %, 384: 0.3 %, 512: 0.7 % |
| 32, 64 + plan | 45.3 | 30.7 | 68 % | 32: 83.6 %, 64: 11.6 %, 128: 2.7 %, 256: 1.1 %, 384: 0.3 %, 512: 0.7 % |
| 32, 64, 96 + plan | 44.7 | 30.7 | 69 % | 32: 83.6 %, 64: 11.6 %, 96: 1.9 %, 128: 0.8 %, 256: 1.1 %, 384: 0.3 %, 512: 0.7 % |
| 32, 64, 96, 128, 192, 256, 384, 512 | 44.2 | 30.7 | 70 % | 32: 83.6 %, 64: 11.6 %, 96: 1.9 %, 128: 0.8 %, 192: 0.8 %, 256: 0.4 %, 384: 0.3 %, 512: 0.7 % |
| 16, 32, 48, 64 + plan | 40.6 | 30.7 | 76 % | 16: 20.2 %, 32: 63.4 %, 48: 9.1 %, 64: 2.5 %, 128: 2.7 %, 256: 1.1 %, 384: 0.3 %, 512: 0.7 % |
| Go's classes up to 512 | 38.7 | 30.7 | 80 % | 16: 20.2 %, 32: 63.4 %, 48: 9.1 %, 64: 2.5 %, 80: 1.2 %, 96: 0.7 %, 112: 0.5 %, 128: 0.3 %, 144: 0.3 %, 160: 0.2 %, 176: 0.2 %, 192: 0.1 %, 208: 0.1 %, 224: 0.1 %, 240: 0.1 %, 256: 0.1 %, 288: 0.1 %, 320: 0.1 %, 352: 0.1 %, 384: 0.1 %, 416: 0.1 %, 448: 0.1 %, 480: 0.0 %, 512: 0.6 % |
| street | multi | 212449 | 5.2 | 2.73 | 27.6 | 39.5 | 132.9 | 38.7 | 0.51 % | 95 % |

street, unique: sets of size classes

| size classes | page B/key | content B/key | fill | keys per class |
|---|--:|--:|--:|---|
| 128, 256, 384, 512 (plan) | 128.0 | 20.1 | 16 % | 128: 100.0 % |
| 64 + plan | 64.0 | 20.1 | 31 % | 64: 100.0 %, 128: 0.0 % |
| 32, 64 + plan | 33.1 | 20.1 | 61 % | 32: 96.7 %, 64: 3.3 %, 128: 0.0 % |
| 32, 64, 96 + plan | 33.1 | 20.1 | 61 % | 32: 96.7 %, 64: 3.3 %, 96: 0.0 % |
| 32, 64, 96, 128, 192, 256, 384, 512 | 33.1 | 20.1 | 61 % | 32: 96.7 %, 64: 3.3 %, 96: 0.0 % |
| 16, 32, 48, 64 + plan | 28.0 | 20.1 | 72 % | 16: 28.7 %, 32: 68.0 %, 48: 3.3 %, 64: 0.0 %, 128: 0.0 % |
| Go's classes up to 512 | 27.9 | 20.1 | 72 % | 16: 28.7 %, 32: 68.0 %, 48: 3.3 %, 64: 0.0 %, 80: 0.0 %, 96: 0.0 % |
| street | unique | 212449 | 5.2 | 1.00 | 9.9 | 20.1 | 128.0 | 27.9 | 0.00 % | 100 % |

dirs, multi: sets of size classes

| size classes | page B/key | content B/key | fill | keys per class |
|---|--:|--:|--:|---|
| 128, 256, 384, 512 (plan) | 144.4 | 57.3 | 40 % | 128: 92.7 %, 256: 3.9 %, 384: 1.1 %, 512: 2.2 % |
| 64 + plan | 92.4 | 57.3 | 62 % | 64: 81.3 %, 128: 11.5 %, 256: 3.9 %, 384: 1.1 %, 512: 2.2 % |
| 32, 64 + plan | 77.4 | 57.3 | 74 % | 32: 46.8 %, 64: 34.4 %, 128: 11.5 %, 256: 3.9 %, 384: 1.1 %, 512: 2.2 % |
| 32, 64, 96 + plan | 74.8 | 57.3 | 77 % | 32: 46.8 %, 64: 34.4 %, 96: 8.3 %, 128: 3.2 %, 256: 3.9 %, 384: 1.1 %, 512: 2.2 % |
| 32, 64, 96, 128, 192, 256, 384, 512 | 73.0 | 57.3 | 79 % | 32: 46.8 %, 64: 34.4 %, 96: 8.3 %, 128: 3.2 %, 192: 2.7 %, 256: 1.2 %, 384: 1.1 %, 512: 2.2 % |
| 16, 32, 48, 64 + plan | 72.5 | 57.3 | 79 % | 16: 6.4 %, 32: 40.5 %, 48: 24.0 %, 64: 10.5 %, 128: 11.5 %, 256: 3.9 %, 384: 1.1 %, 512: 2.2 % |
| Go's classes up to 512 | 64.9 | 57.3 | 88 % | 16: 6.4 %, 32: 40.5 %, 48: 24.0 %, 64: 10.5 %, 80: 5.3 %, 96: 3.0 %, 112: 1.9 %, 128: 1.3 %, 144: 1.0 %, 160: 0.8 %, 176: 0.5 %, 192: 0.5 %, 208: 0.4 %, 224: 0.3 %, 240: 0.3 %, 256: 0.2 %, 288: 0.4 %, 320: 0.3 %, 352: 0.2 %, 384: 0.2 %, 416: 0.2 %, 448: 0.1 %, 480: 0.1 %, 512: 1.7 % |
| dirs | multi | 86215 | 8.9 | 3.44 | 67.9 | 84.3 | 144.4 | 64.9 | 1.61 % | 81 % |

dirs, unique: sets of size classes

| size classes | page B/key | content B/key | fill | keys per class |
|---|--:|--:|--:|---|
| 128, 256, 384, 512 (plan) | 128.0 | 29.3 | 23 % | 128: 100.0 %, 256: 0.0 % |
| 64 + plan | 65.2 | 29.3 | 45 % | 64: 98.2 %, 128: 1.8 %, 256: 0.0 % |
| 32, 64 + plan | 43.2 | 29.3 | 68 % | 32: 68.7 %, 64: 29.5 %, 128: 1.8 %, 256: 0.0 % |
| 32, 64, 96 + plan | 42.6 | 29.3 | 69 % | 32: 68.7 %, 64: 29.5 %, 96: 1.7 %, 128: 0.1 %, 256: 0.0 % |
| 32, 64, 96, 128, 192, 256, 384, 512 | 42.6 | 29.3 | 69 % | 32: 68.7 %, 64: 29.5 %, 96: 1.7 %, 128: 0.1 %, 192: 0.0 % |
| 16, 32, 48, 64 + plan | 37.8 | 29.3 | 78 % | 16: 9.9 %, 32: 58.8 %, 48: 24.1 %, 64: 5.4 %, 128: 1.8 %, 256: 0.0 % |
| Go's classes up to 512 | 37.0 | 29.3 | 79 % | 16: 9.9 %, 32: 58.8 %, 48: 24.1 %, 64: 5.4 %, 80: 1.3 %, 96: 0.4 %, 112: 0.1 %, 128: 0.0 %, 144: 0.0 %, 160: 0.0 %, 176: 0.0 % |
| dirs | unique | 86215 | 8.9 | 1.00 | 15.4 | 29.3 | 128.0 | 37.0 | 0.00 % | 98 % |
