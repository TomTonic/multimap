| case | packing | objects a lookup | lines a lookup | routing objects | routing B a key | all B a key | objects a key |
|---|---|--:|--:|--:|--:|--:|--:|
| street single-value 4096, today's pages | no packing | 4.01 | 7.70 | 100 (0 packed) | 3.4 | 30.4 | 0.220 |
| street single-value 4096, today's pages | packed into 128 B | 3.77 | 7.47 | 80 (19 packed) | 3.1 | 30.1 | 0.215 |
| street single-value 4096, today's pages | packed into 256 B | 3.41 | 7.14 | 55 (25 packed) | 2.8 | 29.8 | 0.209 |
| street single-value 4096, today's pages | packed into 512 B | 3.09 | 6.86 | 33 (24 packed) | 2.7 | 29.7 | 0.204 |
| street single-value 4096, range pages | no packing | 4.01 | 7.75 | 100 (0 packed) | 2.1 | 27.0 | 0.095 |
| street single-value 4096, range pages | packed into 128 B | 3.44 | 7.31 | 56 (28 packed) | 1.6 | 26.5 | 0.084 |
| street single-value 4096, range pages | packed into 256 B | 3.11 | 7.18 | 30 (24 packed) | 1.2 | 26.1 | 0.078 |
| street single-value 4096, range pages | packed into 512 B | 3.01 | 7.10 | 26 (21 packed) | 1.2 | 26.1 | 0.077 |
| street single-value 65536, today's pages | no packing | 5.87 | 10.82 | 1499 (0 packed) | 3.5 | 28.3 | 0.227 |
| street single-value 65536, today's pages | packed into 128 B | 5.59 | 10.55 | 1266 (195 packed) | 3.3 | 28.1 | 0.224 |
| street single-value 65536, today's pages | packed into 256 B | 4.99 | 10.00 | 907 (382 packed) | 2.9 | 27.7 | 0.218 |
| street single-value 65536, today's pages | packed into 512 B | 4.32 | 9.30 | 662 (312 packed) | 2.8 | 27.6 | 0.215 |
| street single-value 65536, range pages | no packing | 5.87 | 11.00 | 1500 (0 packed) | 1.9 | 24.6 | 0.092 |
| street single-value 65536, range pages | packed into 128 B | 5.17 | 10.41 | 866 (406 packed) | 1.5 | 24.2 | 0.082 |
| street single-value 65536, range pages | packed into 256 B | 4.55 | 9.71 | 601 (351 packed) | 1.3 | 24.0 | 0.078 |
| street single-value 65536, range pages | packed into 512 B | 4.02 | 9.44 | 481 (283 packed) | 1.2 | 23.9 | 0.077 |
| street natural 4096, today's pages | no packing | 4.23 | 7.81 | 169 (0 packed) | 4.9 | 40.9 | 0.307 |
| street natural 4096, today's pages | packed into 128 B | 3.94 | 7.54 | 122 (38 packed) | 4.2 | 40.2 | 0.296 |
| street natural 4096, today's pages | packed into 256 B | 3.53 | 7.18 | 82 (42 packed) | 3.7 | 39.7 | 0.286 |
| street natural 4096, today's pages | packed into 512 B | 3.15 | 6.93 | 52 (33 packed) | 3.4 | 39.4 | 0.279 |
| street natural 4096, range pages | no packing | 4.23 | 7.99 | 170 (0 packed) | 3.3 | 36.5 | 0.147 |
| street natural 4096, range pages | packed into 128 B | 3.66 | 7.40 | 90 (48 packed) | 2.5 | 35.7 | 0.128 |
| street natural 4096, range pages | packed into 256 B | 3.24 | 7.26 | 54 (35 packed) | 2.1 | 35.2 | 0.119 |
| street natural 4096, range pages | packed into 512 B | 3.05 | 7.18 | 33 (23 packed) | 1.8 | 35.0 | 0.114 |
| street natural 65536, today's pages | no packing | 6.14 | 10.99 | 2819 (0 packed) | 5.3 | 39.5 | 0.323 |
| street natural 65536, today's pages | packed into 128 B | 5.77 | 10.65 | 1969 (627 packed) | 4.6 | 38.8 | 0.310 |
| street natural 65536, today's pages | packed into 256 B | 5.11 | 10.05 | 1339 (747 packed) | 3.9 | 38.2 | 0.300 |
| street natural 65536, today's pages | packed into 512 B | 4.41 | 9.40 | 982 (579 packed) | 3.8 | 38.1 | 0.295 |
| street natural 65536, range pages | no packing | 6.14 | 11.16 | 2821 (0 packed) | 3.4 | 35.1 | 0.157 |
| street natural 65536, range pages | packed into 128 B | 5.39 | 10.50 | 1411 (832 packed) | 2.5 | 34.2 | 0.135 |
| street natural 65536, range pages | packed into 256 B | 4.73 | 9.90 | 971 (660 packed) | 2.2 | 33.8 | 0.129 |
| street natural 65536, range pages | packed into 512 B | 4.11 | 9.42 | 728 (486 packed) | 2.1 | 33.7 | 0.125 |
| dirs single-value 4096, today's pages | no packing | 8.64 | 13.40 | 237 (0 packed) | 6.5 | 54.0 | 0.377 |
| dirs single-value 4096, today's pages | packed into 128 B | 6.88 | 13.22 | 156 (56 packed) | 5.6 | 53.2 | 0.357 |
| dirs single-value 4096, today's pages | packed into 256 B | 5.13 | 10.80 | 110 (54 packed) | 5.1 | 52.6 | 0.346 |
| dirs single-value 4096, today's pages | packed into 512 B | 4.54 | 10.25 | 90 (43 packed) | 5.1 | 52.6 | 0.341 |
| dirs single-value 4096, range pages | no packing | 8.64 | 14.11 | 237 (0 packed) | 4.6 | 49.6 | 0.204 |
| dirs single-value 4096, range pages | packed into 128 B | 7.03 | 13.68 | 124 (66 packed) | 3.4 | 48.5 | 0.176 |
| dirs single-value 4096, range pages | packed into 256 B | 4.80 | 10.84 | 95 (53 packed) | 3.1 | 48.2 | 0.169 |
| dirs single-value 4096, range pages | packed into 512 B | 4.04 | 10.89 | 78 (47 packed) | 2.9 | 48.0 | 0.165 |
| dirs single-value 65536, today's pages | no packing | 12.37 | 18.62 | 3487 (0 packed) | 5.3 | 40.9 | 0.297 |
| dirs single-value 65536, today's pages | packed into 128 B | 9.79 | 17.71 | 2045 (837 packed) | 4.3 | 40.0 | 0.275 |
| dirs single-value 65536, today's pages | packed into 256 B | 6.94 | 14.17 | 1447 (784 packed) | 3.8 | 39.5 | 0.266 |
| dirs single-value 65536, today's pages | packed into 512 B | 5.94 | 13.30 | 1151 (651 packed) | 3.8 | 39.4 | 0.262 |
| dirs single-value 65536, range pages | no packing | 12.37 | 20.34 | 3491 (0 packed) | 4.1 | 38.7 | 0.173 |
| dirs single-value 65536, range pages | packed into 128 B | 9.61 | 17.23 | 1733 (1021 packed) | 3.0 | 37.7 | 0.146 |
| dirs single-value 65536, range pages | packed into 256 B | 7.10 | 15.53 | 1258 (797 packed) | 2.7 | 37.3 | 0.139 |
| dirs single-value 65536, range pages | packed into 512 B | 5.84 | 13.60 | 1009 (613 packed) | 2.6 | 37.2 | 0.135 |
| dirs natural 4096, today's pages | no packing | 8.85 | 13.57 | 346 (0 packed) | 8.4 | 72.2 | 0.465 |
| dirs natural 4096, today's pages | packed into 128 B | 7.01 | 13.33 | 215 (81 packed) | 7.0 | 70.8 | 0.433 |
| dirs natural 4096, today's pages | packed into 256 B | 5.21 | 11.00 | 145 (79 packed) | 6.2 | 70.0 | 0.416 |
| dirs natural 4096, today's pages | packed into 512 B | 4.60 | 10.32 | 114 (59 packed) | 6.3 | 70.0 | 0.408 |
| dirs natural 4096, range pages | no packing | 8.85 | 15.20 | 346 (0 packed) | 6.7 | 67.6 | 0.283 |
| dirs natural 4096, range pages | packed into 128 B | 7.25 | 12.96 | 186 (102 packed) | 5.2 | 66.2 | 0.244 |
| dirs natural 4096, range pages | packed into 256 B | 5.00 | 10.95 | 134 (79 packed) | 4.6 | 65.5 | 0.231 |
| dirs natural 4096, range pages | packed into 512 B | 4.48 | 11.24 | 104 (62 packed) | 4.5 | 65.5 | 0.224 |
| dirs natural 65536, today's pages | no packing | 12.64 | 18.81 | 5280 (0 packed) | 7.4 | 58.8 | 0.397 |
| dirs natural 65536, today's pages | packed into 128 B | 9.93 | 17.83 | 2890 (1341 packed) | 5.7 | 57.2 | 0.360 |
| dirs natural 65536, today's pages | packed into 256 B | 7.05 | 14.28 | 2079 (1177 packed) | 5.2 | 56.6 | 0.348 |
| dirs natural 65536, today's pages | packed into 512 B | 6.03 | 13.32 | 1634 (945 packed) | 5.1 | 56.5 | 0.341 |
| dirs natural 65536, range pages | no packing | 12.64 | 20.44 | 5294 (0 packed) | 6.1 | 56.0 | 0.256 |
| dirs natural 65536, range pages | packed into 128 B | 9.83 | 17.52 | 2606 (1547 packed) | 4.5 | 54.4 | 0.215 |
| dirs natural 65536, range pages | packed into 256 B | 7.46 | 16.62 | 1852 (1193 packed) | 4.0 | 53.8 | 0.203 |
| dirs natural 65536, range pages | packed into 512 B | 5.95 | 13.90 | 1461 (907 packed) | 3.8 | 53.7 | 0.197 |
| links single-value 4096, today's pages | no packing | 3.74 | 7.55 | 101 (0 packed) | 4.6 | 36.3 | 0.302 |
| links single-value 4096, today's pages | packed into 128 B | 3.68 | 7.52 | 95 (6 packed) | 4.5 | 36.2 | 0.301 |
| links single-value 4096, today's pages | packed into 256 B | 3.46 | 7.21 | 73 (21 packed) | 4.0 | 35.7 | 0.296 |
| links single-value 4096, today's pages | packed into 512 B | 3.16 | 6.88 | 46 (22 packed) | 3.6 | 35.4 | 0.289 |
| links single-value 4096, range pages | no packing | 3.74 | 7.95 | 101 (0 packed) | 2.2 | 30.5 | 0.105 |
| links single-value 4096, range pages | packed into 128 B | 3.48 | 7.64 | 72 (23 packed) | 1.9 | 30.1 | 0.098 |
| links single-value 4096, range pages | packed into 256 B | 3.06 | 7.45 | 37 (23 packed) | 1.4 | 29.7 | 0.089 |
| links single-value 4096, range pages | packed into 512 B | 2.98 | 7.40 | 26 (23 packed) | 1.4 | 29.6 | 0.086 |
| links single-value 65536, today's pages | no packing | 5.32 | 10.17 | 1832 (0 packed) | 4.1 | 33.1 | 0.267 |
| links single-value 65536, today's pages | packed into 128 B | 5.02 | 9.96 | 1429 (278 packed) | 3.8 | 32.7 | 0.261 |
| links single-value 65536, today's pages | packed into 256 B | 4.71 | 9.68 | 1067 (424 packed) | 3.4 | 32.4 | 0.256 |
| links single-value 65536, today's pages | packed into 512 B | 4.41 | 9.36 | 853 (353 packed) | 3.3 | 32.3 | 0.252 |
| links single-value 65536, range pages | no packing | 5.32 | 10.56 | 1836 (0 packed) | 2.3 | 28.8 | 0.112 |
| links single-value 65536, range pages | packed into 128 B | 4.74 | 10.08 | 1051 (477 packed) | 1.8 | 28.3 | 0.100 |
| links single-value 65536, range pages | packed into 256 B | 4.39 | 9.83 | 804 (420 packed) | 1.6 | 28.0 | 0.096 |
| links single-value 65536, range pages | packed into 512 B | 3.96 | 9.47 | 608 (324 packed) | 1.5 | 27.9 | 0.093 |
| links natural 4096, today's pages | no packing | 4.93 | 9.04 | 1048 (0 packed) | 21.0 | 157.8 | 1.038 |
| links natural 4096, today's pages | packed into 128 B | 4.42 | 8.63 | 584 (278 packed) | 16.1 | 153.0 | 0.924 |
| links natural 4096, today's pages | packed into 256 B | 3.99 | 8.22 | 418 (235 packed) | 13.6 | 150.5 | 0.884 |
| links natural 4096, today's pages | packed into 512 B | 3.68 | 7.93 | 398 (220 packed) | 14.2 | 151.0 | 0.879 |
| links natural 4096, range pages | no packing | 4.93 | 9.24 | 1050 (0 packed) | 21.1 | 155.5 | 0.874 |
| links natural 4096, range pages | packed into 128 B | 4.48 | 8.77 | 608 (292 packed) | 16.7 | 151.1 | 0.766 |
| links natural 4096, range pages | packed into 256 B | 4.04 | 8.40 | 423 (247 packed) | 13.7 | 148.1 | 0.720 |
| links natural 4096, range pages | packed into 512 B | 3.69 | 8.27 | 370 (213 packed) | 13.3 | 147.7 | 0.708 |
| links natural 65536, today's pages | no packing | 6.72 | 11.77 | 18222 (0 packed) | 21.8 | 154.3 | 1.052 |
| links natural 65536, today's pages | packed into 128 B | 5.81 | 11.11 | 9545 (4720 packed) | 16.2 | 148.7 | 0.920 |
| links natural 65536, today's pages | packed into 256 B | 5.30 | 10.72 | 7313 (4013 packed) | 14.6 | 147.1 | 0.886 |
| links natural 65536, today's pages | packed into 512 B | 4.88 | 10.44 | 5677 (3278 packed) | 14.1 | 146.6 | 0.861 |
| links natural 65536, range pages | no packing | 6.72 | 12.21 | 18235 (0 packed) | 22.2 | 152.9 | 0.902 |
| links natural 65536, range pages | packed into 128 B | 5.93 | 11.46 | 10084 (5128 packed) | 17.2 | 147.9 | 0.777 |
| links natural 65536, range pages | packed into 256 B | 5.39 | 11.10 | 7400 (4235 packed) | 14.8 | 145.5 | 0.736 |
| links natural 65536, range pages | packed into 512 B | 4.96 | 10.87 | 5795 (3337 packed) | 13.9 | 144.6 | 0.712 |
| url single-value 4096, today's pages | no packing | 6.86 | 11.89 | 317 (0 packed) | 8.3 | 85.8 | 0.497 |
| url single-value 4096, today's pages | packed into 128 B | 6.00 | 11.19 | 223 (64 packed) | 7.3 | 84.8 | 0.474 |
| url single-value 4096, today's pages | packed into 256 B | 5.21 | 10.38 | 163 (70 packed) | 6.3 | 83.8 | 0.459 |
| url single-value 4096, today's pages | packed into 512 B | 4.92 | 11.62 | 95 (51 packed) | 6.2 | 83.7 | 0.443 |
| url single-value 4096, range pages | no packing | 6.86 | 12.60 | 318 (0 packed) | 6.6 | 79.6 | 0.296 |
| url single-value 4096, range pages | packed into 128 B | 5.89 | 11.80 | 195 (76 packed) | 5.5 | 78.5 | 0.266 |
| url single-value 4096, range pages | packed into 256 B | 5.08 | 11.02 | 129 (67 packed) | 4.4 | 77.5 | 0.250 |
| url single-value 4096, range pages | packed into 512 B | 4.82 | 11.96 | 66 (44 packed) | 3.8 | 76.9 | 0.234 |
| url single-value 65536, today's pages | no packing | 9.19 | 15.83 | 4934 (0 packed) | 7.8 | 76.7 | 0.449 |
| url single-value 65536, today's pages | packed into 128 B | 8.07 | 15.05 | 3166 (1128 packed) | 6.5 | 75.4 | 0.422 |
| url single-value 65536, today's pages | packed into 256 B | 6.99 | 13.89 | 2294 (1068 packed) | 5.8 | 74.7 | 0.408 |
| url single-value 65536, today's pages | packed into 512 B | 6.57 | 13.56 | 1825 (822 packed) | 5.7 | 74.7 | 0.401 |
| url single-value 65536, range pages | no packing | 9.19 | 15.98 | 4942 (0 packed) | 6.1 | 70.5 | 0.268 |
| url single-value 65536, range pages | packed into 128 B | 7.54 | 14.58 | 2727 (1332 packed) | 4.7 | 69.1 | 0.235 |
| url single-value 65536, range pages | packed into 256 B | 6.85 | 14.06 | 1954 (1060 packed) | 4.0 | 68.5 | 0.223 |
| url single-value 65536, range pages | packed into 512 B | 6.40 | 13.72 | 1531 (815 packed) | 3.8 | 68.3 | 0.216 |
| url natural 4096, today's pages | no packing | 7.13 | 12.01 | 540 (0 packed) | 12.0 | 112.9 | 0.659 |
| url natural 4096, today's pages | packed into 128 B | 6.13 | 11.22 | 313 (136 packed) | 9.6 | 110.6 | 0.604 |
| url natural 4096, today's pages | packed into 256 B | 5.33 | 10.40 | 245 (132 packed) | 8.7 | 109.6 | 0.587 |
| url natural 4096, today's pages | packed into 512 B | 5.00 | 11.23 | 149 (88 packed) | 8.0 | 108.9 | 0.563 |
| url natural 4096, range pages | no packing | 7.13 | 12.58 | 540 (0 packed) | 10.6 | 106.2 | 0.462 |
| url natural 4096, range pages | packed into 128 B | 6.08 | 11.72 | 298 (148 packed) | 8.4 | 104.1 | 0.403 |
| url natural 4096, range pages | packed into 256 B | 5.25 | 10.97 | 219 (128 packed) | 7.1 | 102.7 | 0.384 |
| url natural 4096, range pages | packed into 512 B | 4.92 | 11.91 | 117 (78 packed) | 6.2 | 101.8 | 0.359 |
| url natural 65536, today's pages | no packing | 9.52 | 15.93 | 8589 (0 packed) | 11.7 | 103.8 | 0.629 |
| url natural 65536, today's pages | packed into 128 B | 8.24 | 15.06 | 4823 (2233 packed) | 9.1 | 101.2 | 0.571 |
| url natural 65536, today's pages | packed into 256 B | 7.12 | 13.89 | 3585 (1955 packed) | 8.3 | 100.4 | 0.552 |
| url natural 65536, today's pages | packed into 512 B | 6.68 | 13.56 | 2917 (1596 packed) | 8.2 | 100.3 | 0.542 |
| url natural 65536, range pages | no packing | 9.52 | 16.06 | 8595 (0 packed) | 10.2 | 97.2 | 0.433 |
| url natural 65536, range pages | packed into 128 B | 7.80 | 14.57 | 4448 (2436 packed) | 7.7 | 94.7 | 0.370 |
| url natural 65536, range pages | packed into 256 B | 7.07 | 14.05 | 3299 (1973 packed) | 6.8 | 93.8 | 0.352 |
| url natural 65536, range pages | packed into 512 B | 6.60 | 13.75 | 2662 (1529 packed) | 6.6 | 93.6 | 0.343 |
