## vpage (step 2 page)
chunk 256 keys, largest class 2, split fill 85%, split by bytes false, min prefix 1, min gain 48, slack 0

| kind | keys | suffix B | pages | keys/page | fill | uniform | page B/key | router B/key | total B/key | tail B/key | prefix B/page | too long |
|---|--:|--:|--:|--:|--:|--:|--:|--:|--:|--:|--:|--:|
| u64 | 262144 | 7.2 | 13055 | 20.1 | 70 % | 100 % | 25.5 | 1.0 | 26.5 | 0.0 | 0.0 | 0 |
| str | 262144 | 17.6 | 19763 | 13.3 | 69 % | 0 % | 37.9 | 1.0 | 38.9 | 4.7 | 4.6 | 0 |
| uuid | 262144 | 34.3 | 36906 | 7.1 | 67 % | 0 % | 72.1 | 2.0 | 74.1 | 26.3 | 0.0 | 0 |
| email | 262144 | 23.9 | 27451 | 9.5 | 70 % | 0 % | 53.5 | 2.0 | 55.5 | 15.9 | 0.0 | 0 |
| url | 262144 | 51.1 | 46303 | 5.7 | 70 % | 0 % | 86.4 | 2.0 | 88.4 | 36.6 | 5.4 | 0 |
| path | 262142 | 43.4 | 34926 | 7.5 | 70 % | 1 % | 64.8 | 2.0 | 66.8 | 21.2 | 14.4 | 2 |
| street | 212449 | 11.3 | 14424 | 14.7 | 70 % | 0 % | 34.7 | 1.0 | 35.7 | 3.1 | 0.6 | 0 |

## lens -lenheader 8
chunk 256 keys, largest class 2, split fill 85%, split by bytes false, min prefix 1, min gain 48, slack 0

lens layout: fixed values false, value 8 B, largest header 8 B (3 entries), classes [128 256 512]

| kind | keys | suffix B | pages | keys/page | fill | page B/key | router B/key | total B/key | header B/page | prefix B/page | keys end at B | keys end within 128 B | too long |
|---|--:|--:|--:|--:|--:|--:|--:|--:|--:|--:|--:|--:|--:|
| u64 | 262144 | 7.2 | 112337 | 2.3 | 33 % | 54.9 | 8.2 | 63.1 | 8.0 | 0.9 | 24 | 100 % | 0 |
| str | 262144 | 17.6 | 112304 | 2.3 | 46 % | 54.8 | 8.2 | 63.1 | 8.0 | 6.7 | 40 | 100 % | 0 |
| uuid | 262144 | 34.3 | 112387 | 2.3 | 61 % | 73.0 | 8.2 | 81.2 | 8.0 | 2.0 | 85 | 100 % | 0 |
| email | 262144 | 23.9 | 112389 | 2.3 | 63 % | 54.9 | 8.2 | 63.1 | 8.0 | 1.8 | 61 | 100 % | 0 |
| url | 262144 | 51.1 | 112334 | 2.3 | 69 % | 81.4 | 8.2 | 89.7 | 8.0 | 12.2 | 112 | 68 % | 0 |
| path | 262142 | 43.4 | 112281 | 2.3 | 66 % | 65.0 | 8.3 | 73.3 | 8.0 | 21.6 | 81 | 91 % | 2 |
| street | 212449 | 11.3 | 91145 | 2.3 | 37 % | 54.9 | 8.3 | 63.2 | 8.0 | 4.1 | 29 | 100 % | 0 |

## lens -lenheader 16
chunk 256 keys, largest class 2, split fill 85%, split by bytes false, min prefix 1, min gain 48, slack 0

lens layout: fixed values false, value 8 B, largest header 16 B (7 entries), classes [128 256 512]

| kind | keys | suffix B | pages | keys/page | fill | page B/key | router B/key | total B/key | header B/page | prefix B/page | keys end at B | keys end within 128 B | too long |
|---|--:|--:|--:|--:|--:|--:|--:|--:|--:|--:|--:|--:|--:|
| u64 | 262144 | 7.2 | 51653 | 5.1 | 72 % | 25.2 | 2.0 | 27.2 | 16.0 | 0.5 | 51 | 100 % | 0 |
| str | 262144 | 17.6 | 51707 | 5.1 | 71 % | 34.1 | 2.0 | 36.1 | 16.0 | 5.9 | 82 | 99 % | 0 |
| uuid | 262144 | 34.3 | 51712 | 5.1 | 66 % | 67.1 | 2.0 | 69.1 | 16.0 | 1.4 | 184 | 0 % | 0 |
| email | 262144 | 23.9 | 51660 | 5.1 | 68 % | 49.7 | 2.0 | 51.7 | 16.0 | 1.4 | 132 | 52 % | 0 |
| url | 262144 | 51.1 | 52504 | 5.0 | 70 % | 78.6 | 2.0 | 80.6 | 15.8 | 9.0 | 236 | 10 % | 0 |
| path | 262142 | 43.4 | 51904 | 5.1 | 70 % | 59.3 | 2.1 | 61.4 | 15.9 | 16.8 | 168 | 32 % | 2 |
| street | 212449 | 11.3 | 41908 | 5.1 | 70 % | 28.8 | 2.0 | 30.8 | 16.0 | 2.9 | 62 | 100 % | 0 |

## lens -lenheader 24
chunk 256 keys, largest class 2, split fill 85%, split by bytes false, min prefix 1, min gain 48, slack 0

lens layout: fixed values false, value 8 B, largest header 24 B (11 entries), classes [128 256 512]

| kind | keys | suffix B | pages | keys/page | fill | page B/key | router B/key | total B/key | header B/page | prefix B/page | keys end at B | keys end within 128 B | too long |
|---|--:|--:|--:|--:|--:|--:|--:|--:|--:|--:|--:|--:|--:|
| u64 | 262144 | 7.2 | 33509 | 7.8 | 72 % | 24.5 | 2.0 | 26.5 | 20.0 | 0.3 | 75 | 100 % | 0 |
| str | 262144 | 17.6 | 33459 | 7.8 | 69 % | 33.8 | 2.0 | 35.8 | 20.0 | 5.6 | 120 | 64 % | 0 |
| uuid | 262144 | 34.3 | 33407 | 7.8 | 67 % | 65.2 | 2.0 | 67.2 | 20.0 | 1.2 | 281 | 0 % | 0 |
| email | 262144 | 23.9 | 33467 | 7.8 | 69 % | 48.4 | 2.0 | 50.4 | 20.0 | 1.2 | 199 | 0 % | 0 |
| url | 262144 | 51.1 | 43688 | 6.0 | 72 % | 76.4 | 2.0 | 78.4 | 17.3 | 8.1 | 282 | 2 % | 0 |
| path | 262142 | 43.4 | 36366 | 7.2 | 71 % | 57.7 | 2.0 | 59.7 | 19.1 | 15.9 | 238 | 9 % | 2 |
| street | 212449 | 11.3 | 27085 | 7.8 | 71 % | 27.9 | 2.0 | 29.9 | 20.0 | 2.5 | 92 | 90 % | 0 |

## lens -lenheader 8 -lenfixed
chunk 256 keys, largest class 2, split fill 85%, split by bytes false, min prefix 1, min gain 48, slack 0

lens layout: fixed values true, value 8 B, largest header 8 B (6 entries), classes [128 256 512]

| kind | keys | suffix B | pages | keys/page | fill | page B/key | router B/key | total B/key | header B/page | prefix B/page | keys end at B | keys end within 128 B | too long |
|---|--:|--:|--:|--:|--:|--:|--:|--:|--:|--:|--:|--:|--:|
| u64 | 262144 | 7.2 | 60641 | 4.3 | 56 % | 29.6 | 7.8 | 37.4 | 8.0 | 0.6 | 38 | 100 % | 0 |
| str | 262144 | 17.6 | 60711 | 4.3 | 68 % | 33.5 | 7.9 | 41.5 | 8.0 | 6.0 | 64 | 100 % | 0 |
| uuid | 262144 | 34.3 | 60688 | 4.3 | 71 % | 60.1 | 7.9 | 68.0 | 8.0 | 1.5 | 151 | 25 % | 0 |
| email | 262144 | 23.9 | 60655 | 4.3 | 68 % | 48.1 | 7.8 | 55.9 | 8.0 | 1.5 | 106 | 79 % | 0 |
| url | 262144 | 51.1 | 60714 | 4.3 | 69 % | 77.9 | 7.9 | 85.8 | 8.0 | 9.5 | 198 | 20 % | 0 |
| path | 262142 | 43.4 | 60901 | 4.3 | 70 % | 57.5 | 7.9 | 65.4 | 8.0 | 17.7 | 138 | 52 % | 2 |
| street | 212449 | 11.3 | 49122 | 4.3 | 62 % | 30.2 | 7.9 | 38.2 | 8.0 | 3.1 | 47 | 100 % | 0 |

## lens -lenheader 16 -lenfixed
chunk 256 keys, largest class 2, split fill 85%, split by bytes false, min prefix 1, min gain 48, slack 0

lens layout: fixed values true, value 8 B, largest header 16 B (14 entries), classes [128 256 512]

| kind | keys | suffix B | pages | keys/page | fill | page B/key | router B/key | total B/key | header B/page | prefix B/page | keys end at B | keys end within 128 B | too long |
|---|--:|--:|--:|--:|--:|--:|--:|--:|--:|--:|--:|--:|--:|
| u64 | 262144 | 7.2 | 26527 | 9.9 | 68 % | 24.3 | 1.9 | 26.2 | 16.0 | 0.2 | 85 | 100 % | 0 |
| str | 262144 | 17.6 | 26568 | 9.9 | 70 % | 32.1 | 1.9 | 34.0 | 16.0 | 5.4 | 143 | 42 % | 0 |
| uuid | 262144 | 34.3 | 31327 | 8.4 | 76 % | 56.7 | 2.0 | 58.7 | 14.6 | 1.2 | 293 | 0 % | 0 |
| email | 262144 | 23.9 | 26500 | 9.9 | 69 % | 47.0 | 1.9 | 48.9 | 16.0 | 1.1 | 243 | 0 % | 0 |
| url | 262144 | 51.1 | 42146 | 6.2 | 72 % | 74.5 | 2.0 | 76.5 | 11.0 | 8.0 | 285 | 1 % | 0 |
| path | 262142 | 43.4 | 32790 | 8.0 | 72 % | 55.8 | 2.0 | 57.8 | 13.7 | 15.8 | 258 | 5 % | 2 |
| street | 212449 | 11.3 | 21493 | 9.9 | 70 % | 27.1 | 1.9 | 29.0 | 16.0 | 2.3 | 108 | 77 % | 0 |

## lens -lenheader 24 -lenfixed
chunk 256 keys, largest class 2, split fill 85%, split by bytes false, min prefix 1, min gain 48, slack 0

lens layout: fixed values true, value 8 B, largest header 24 B (22 entries), classes [128 256 512]

| kind | keys | suffix B | pages | keys/page | fill | page B/key | router B/key | total B/key | header B/page | prefix B/page | keys end at B | keys end within 128 B | too long |
|---|--:|--:|--:|--:|--:|--:|--:|--:|--:|--:|--:|--:|--:|
| u64 | 262144 | 7.2 | 16951 | 15.5 | 68 % | 23.9 | 1.0 | 24.9 | 20.3 | 0.2 | 129 | 46 % | 0 |
| str | 262144 | 17.6 | 17802 | 14.7 | 70 % | 31.6 | 1.0 | 32.6 | 19.8 | 5.0 | 210 | 1 % | 0 |
| uuid | 262144 | 34.3 | 31327 | 8.4 | 76 % | 56.7 | 2.0 | 58.7 | 14.6 | 1.2 | 293 | 0 % | 0 |
| email | 262144 | 23.9 | 24523 | 10.7 | 70 % | 46.5 | 1.3 | 47.9 | 16.6 | 1.0 | 262 | 0 % | 0 |
| url | 262144 | 51.1 | 41767 | 6.3 | 72 % | 74.4 | 2.0 | 76.3 | 11.0 | 7.9 | 288 | 1 % | 0 |
| path | 262142 | 43.4 | 31322 | 8.4 | 72 % | 55.4 | 1.9 | 57.3 | 13.9 | 15.8 | 269 | 2 % | 2 |
| street | 212449 | 11.3 | 13747 | 15.5 | 70 % | 26.9 | 1.0 | 27.9 | 20.3 | 2.0 | 167 | 23 % | 0 |

## lens -lenheader 32 -lenfixed -lenclasses 128,256,512,1024
chunk 256 keys, largest class 2, split fill 85%, split by bytes false, min prefix 1, min gain 48, slack 0

lens layout: fixed values true, value 8 B, largest header 32 B (30 entries), classes [128 256 512 1024]

| kind | keys | suffix B | pages | keys/page | fill | page B/key | router B/key | total B/key | header B/page | prefix B/page | keys end at B | keys end within 128 B | too long |
|---|--:|--:|--:|--:|--:|--:|--:|--:|--:|--:|--:|--:|--:|
| u64 | 262144 | 7.2 | 12634 | 20.7 | 68 % | 23.8 | 1.0 | 24.8 | 26.6 | 0.2 | 172 | 0 % | 0 |
| str | 262144 | 17.6 | 12576 | 20.8 | 70 % | 32.8 | 1.0 | 33.8 | 26.7 | 4.4 | 308 | 0 % | 0 |
| uuid | 262144 | 34.3 | 15878 | 16.5 | 72 % | 59.6 | 1.0 | 60.6 | 21.7 | 1.0 | 572 | 0 % | 0 |
| email | 262144 | 23.9 | 12628 | 20.8 | 69 % | 47.0 | 1.0 | 48.0 | 26.6 | 0.7 | 510 | 0 % | 0 |
| url | 262144 | 51.1 | 20960 | 12.5 | 72 % | 76.3 | 1.0 | 77.3 | 18.2 | 6.0 | 584 | 0 % | 0 |
| path | 262142 | 43.4 | 16712 | 15.7 | 72 % | 58.2 | 1.0 | 59.2 | 21.2 | 12.4 | 530 | 0 % | 2 |
| street | 212449 | 11.3 | 10189 | 20.9 | 69 % | 27.7 | 1.0 | 28.7 | 26.7 | 1.8 | 229 | 2 % | 0 |
