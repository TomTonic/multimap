## street: 212449 keys (0 too big for a page, 212449 absent keys)

### point lookups of present keys: tags against lengths, numbers
A = directory of tags (vpage), numbers; B = length header (lpage), numbers

A: 63.67 ns per operation (median)
B: 87.18 ns per operation (median)
A needs 27.4% less time than B: B takes 1.38 times as long as A.
With 95% confidence the difference lies between 26.9% and 28.2% less time (B takes between 1.37 and 1.39 times as long as A).
Noise floor: 1.1%. Identical code measured against itself on this setup can look that different, so smaller differences mean nothing.
Neighbouring measurements are correlated (autocorrelation +0.27), so they were resampled in blocks of 4.

Verdict: RESOLVED. A is faster than B: the difference is real, as its interval excludes zero and it exceeds the noise floor.

Warnings:
  - the program holds 115 MB of live data, more than the caches of many machines; for data that size, where it lies in memory can shift the result by several points in ways the interval does not cover, and differently in the next process; run the comparison in several processes with the multiproc package

### point lookups of absent keys: tags against lengths, numbers
A = directory of tags (vpage), numbers; B = length header (lpage), numbers

A: 36.46 ns per operation (median)
B: 57.79 ns per operation (median)
A needs 37.2% less time than B: B takes 1.59 times as long as A.
With 95% confidence the difference lies between 35.4% and 38.8% less time (B takes between 1.55 and 1.63 times as long as A).
Noise floor: 1.5%. Identical code measured against itself on this setup can look that different, so smaller differences mean nothing.
Neighbouring measurements are nearly independent (autocorrelation -0.03), so they were resampled one by one.

Verdict: RESOLVED. A is faster than B: the difference is real, as its interval excludes zero and it exceeds the noise floor.

Warnings:
  - the program holds 115 MB of live data, more than the caches of many machines; for data that size, where it lies in memory can shift the result by several points in ways the interval does not cover, and differently in the next process; run the comparison in several processes with the multiproc package

### point lookups of present keys: numbers against names, both in length-header pages
A = length header (lpage), numbers; B = length header (lpage), names

A: 87.37 ns per operation (median)
B: 83.96 ns per operation (median)
A needs 3.21% more time than B: B takes only 0.97 times as long as A.
With 95% confidence the difference lies between 2.48% and 4.44% more time (B takes between 0.957 and 0.98 times as long as A).
Noise floor: 1.12%. Identical code measured against itself on this setup can look that different, so smaller differences mean nothing.
Neighbouring measurements are nearly independent (autocorrelation +0.19), so they were resampled one by one.

Verdict: RESOLVED. A is slower than B: the difference is real, as its interval excludes zero and it exceeds the noise floor.

Warnings:
  - the program holds 115 MB of live data, more than the caches of many machines; for data that size, where it lies in memory can shift the result by several points in ways the interval does not cover, and differently in the next process; run the comparison in several processes with the multiproc package
  - candidate A drifted during the run: its measurements were 1.47% higher in the second half than in the first, so the machine did not hold still

### point lookups of present keys: names as bytes against names copied to strings
A = length header (lpage), names; B = length header (lpage), names copied to strings

A: 84.16 ns per operation (median)
B: 85.95 ns per operation (median)
A needs 2.33% less time than B: B takes 1.02 times as long as A.
With 95% confidence the difference lies between 1.52% and 3.55% less time (B takes between 1.02 and 1.04 times as long as A).
Noise floor: 0.991%. Identical code measured against itself on this setup can look that different, so smaller differences mean nothing.
Neighbouring measurements are nearly independent (autocorrelation +0.20), so they were resampled one by one.

Verdict: RESOLVED. A is faster than B: the difference is real, as its interval excludes zero and it exceeds the noise floor.

Warnings:
  - the program holds 116 MB of live data, more than the caches of many machines; for data that size, where it lies in memory can shift the result by several points in ways the interval does not cover, and differently in the next process; run the comparison in several processes with the multiproc package

| layout | pages | keys per page | page bytes per key |
|---|--:|--:|--:|
| directory of tags (vpage), numbers | 14622 | 14.5 | 35.2 |
| length header (lpage), numbers | 13762 | 15.4 | 27.3 |
| length header (lpage), names | 27428 | 7.7 | 31.0 |

## dirs: 86215 keys (0 too big for a page, 86215 absent keys)

### point lookups of present keys: tags against lengths, numbers
A = directory of tags (vpage), numbers; B = length header (lpage), numbers

A: 70.1 ns per operation (median)
B: 81.8 ns per operation (median)
A needs 14% less time than B: B takes 1.16 times as long as A.
With 95% confidence the difference lies between 13.6% and 14.8% less time (B takes between 1.16 and 1.17 times as long as A).
Noise floor: 1.31%. Identical code measured against itself on this setup can look that different, so smaller differences mean nothing.
Neighbouring measurements are nearly independent (autocorrelation +0.13), so they were resampled one by one.

Verdict: RESOLVED. A is faster than B: the difference is real, as its interval excludes zero and it exceeds the noise floor.

Warnings:
  - the program holds 161 MB of live data, more than the caches of many machines; for data that size, where it lies in memory can shift the result by several points in ways the interval does not cover, and differently in the next process; run the comparison in several processes with the multiproc package

### point lookups of absent keys: tags against lengths, numbers
A = directory of tags (vpage), numbers; B = length header (lpage), numbers

A: 50.45 ns per operation (median)
B: 51.83 ns per operation (median)
A needs 2.41% less time than B: B takes 1.02 times as long as A.
With 95% confidence the difference lies between 3.34% more time and 6% less time (B takes between 0.968 and 1.06 times as long as A).
Noise floor: 3.53%. Identical code measured against itself on this setup can look that different, so smaller differences mean nothing.
Neighbouring measurements are nearly independent (autocorrelation -0.17), so they were resampled one by one.

Verdict: NOT RESOLVED. This run did not establish which candidate is faster. That is not the same as "equally fast": the warnings below say what is missing.

Warnings:
  - the difference of 2.41% does not clear the 3.53% noise floor, which is what this setup reports between two runs of identical code
  - the interval [-3.34%, 6.00%] includes zero, so a difference in either direction is consistent with these measurements
  - the program holds 161 MB of live data, more than the caches of many machines; for data that size, where it lies in memory can shift the result by several points in ways the interval does not cover, and differently in the next process; run the comparison in several processes with the multiproc package
  - candidate B drifted during the run: its measurements were 5.84% lower in the second half than in the first, so the machine did not hold still

### point lookups of present keys: numbers against names, both in length-header pages
A = length header (lpage), numbers; B = length header (lpage), names

A: 81.29 ns per operation (median)
B: 84.82 ns per operation (median)
A needs 3.56% less time than B: B takes 1.04 times as long as A.
With 95% confidence the difference lies between 2.34% and 4.99% less time (B takes between 1.02 and 1.05 times as long as A).
Noise floor: 0.937%. Identical code measured against itself on this setup can look that different, so smaller differences mean nothing.
Neighbouring measurements are nearly independent (autocorrelation +0.15), so they were resampled one by one.

Verdict: RESOLVED. A is faster than B: the difference is real, as its interval excludes zero and it exceeds the noise floor.

Warnings:
  - the program holds 163 MB of live data, more than the caches of many machines; for data that size, where it lies in memory can shift the result by several points in ways the interval does not cover, and differently in the next process; run the comparison in several processes with the multiproc package
  - in 30% of A/A runs this setup reported a difference between identical code, well above the 10% expected; treat any confidence from it with suspicion

### point lookups of present keys: names as bytes against names copied to strings
A = length header (lpage), names; B = length header (lpage), names copied to strings

A: 84.65 ns per operation (median)
B: 93.57 ns per operation (median)
A needs 9.68% less time than B: B takes 1.11 times as long as A.
With 95% confidence the difference lies between 9.11% and 10.5% less time (B takes between 1.1 and 1.12 times as long as A).
Noise floor: 0.882%. Identical code measured against itself on this setup can look that different, so smaller differences mean nothing.
Neighbouring measurements are correlated (autocorrelation +0.21), so they were resampled in blocks of 4.

Verdict: RESOLVED. A is faster than B: the difference is real, as its interval excludes zero and it exceeds the noise floor.

Warnings:
  - the program holds 162 MB of live data, more than the caches of many machines; for data that size, where it lies in memory can shift the result by several points in ways the interval does not cover, and differently in the next process; run the comparison in several processes with the multiproc package

| layout | pages | keys per page | page bytes per key |
|---|--:|--:|--:|
| directory of tags (vpage), numbers | 9814 | 8.8 | 55.9 |
| length header (lpage), numbers | 8697 | 9.9 | 46.6 |
| length header (lpage), names | 12135 | 7.1 | 59.9 |

checksum 13722112652506622787
