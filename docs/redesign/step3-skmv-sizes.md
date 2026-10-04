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

Summary (bytes per key, the page only):

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
