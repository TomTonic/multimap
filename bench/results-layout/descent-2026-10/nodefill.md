| case | node | a 1,000 keys | slots used 1 / 2 / 3 / 4 / 5 (N12: 6 / 7 / 8 / 9-12) | with end page | with prefix tail |
|---|---|--:|---|--:|--:|
| str single-value 4096 | N5 | 0.5 | 0 / 50 / 0 / 50 / 0 % | 0 % | 0 % |
| str single-value 4096 | N12 | 2.7 | 0 / 9 / 0 / 91 % | 0 % | 0 % |
| str single-value 65536 | N5 | 11.2 | 0 / 85 / 7 / 0 / 9 % | 0 % | 0 % |
| str single-value 65536 | N12 | 5.4 | 9 / 23 / 23 / 46 % | 0 % | 0 % |
| email single-value 4096 | N5 | 0.0 | 0 / 0 / 0 / 0 / 0 % | 0 % | 0 % |
| email single-value 4096 | N12 | 0.0 | 0 / 0 / 0 / 0 % | 0 % | 0 % |
| email single-value 65536 | N5 | 0.0 | 0 / 0 / 0 / 0 / 0 % | 0 % | 0 % |
| email single-value 65536 | N12 | 0.0 | 0 / 0 / 0 / 0 % | 0 % | 0 % |
| url single-value 4096 | N5 | 43.9 | 0 / 34 / 22 / 21 / 23 % | 0 % | 1 % |
| url single-value 4096 | N12 | 25.1 | 32 / 21 / 17 / 29 % | 0 % | 1 % |
| url single-value 65536 | N5 | 45.0 | 0 / 40 / 25 / 19 / 16 % | 0 % | 2 % |
| url single-value 65536 | N12 | 23.6 | 27 / 21 / 17 / 35 % | 0 % | 3 % |
| path single-value 4096 | N5 | 53.5 | 0 / 52 / 21 / 16 / 12 % | 0 % | 7 % |
| path single-value 4096 | N12 | 18.3 | 32 / 25 / 16 / 27 % | 0 % | 1 % |
| path single-value 65536 | N5 | 50.8 | 0 / 52 / 21 / 14 / 12 % | 0 % | 6 % |
| path single-value 65536 | N12 | 15.8 | 29 / 22 / 14 / 35 % | 0 % | 11 % |
| street single-value 4096 | N5 | 5.4 | 0 / 14 / 41 / 18 / 27 % | 0 % | 0 % |
| street single-value 4096 | N12 | 14.6 | 27 / 12 / 15 / 47 % | 0 % | 0 % |
| street single-value 65536 | N5 | 6.0 | 0 / 25 / 27 / 25 / 23 % | 1 % | 0 % |
| street single-value 65536 | N12 | 9.8 | 19 / 16 / 15 / 50 % | 1 % | 0 % |
| dirs single-value 4096 | N5 | 34.4 | 0 / 49 / 23 / 16 / 11 % | 0 % | 8 % |
| dirs single-value 4096 | N12 | 15.6 | 22 / 22 / 16 / 41 % | 0 % | 2 % |
| dirs single-value 65536 | N5 | 35.2 | 0 / 52 / 23 / 14 / 11 % | 4 % | 6 % |
| dirs single-value 65536 | N12 | 12.8 | 26 / 21 / 18 / 36 % | 7 % | 5 % |
| links single-value 4096 | N5 | 2.9 | 0 / 25 / 25 / 25 / 25 % | 0 % | 0 % |
| links single-value 4096 | N12 | 10.7 | 7 / 11 / 14 / 68 % | 0 % | 0 % |
| links single-value 65536 | N5 | 10.2 | 0 / 33 / 27 / 20 / 20 % | 5 % | 0 % |
| links single-value 65536 | N12 | 10.6 | 19 / 15 / 16 / 50 % | 7 % | 0 % |
