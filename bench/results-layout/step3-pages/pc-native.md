## street: 212449 keys (0 too big for a page, 212449 absent keys)

### point lookups of present keys: tags against lengths, numbers
A = directory of tags (vpage), numbers; B = length header (lpage), numbers

A: 69 ns per operation (median)
B: 85.61 ns per operation (median)
A needs 16% less time than B: B takes 1.19 times as long as A.
With 95% confidence the difference lies between 13.7% and 16.6% less time (B takes between 1.16 and 1.20 times as long as A).
Noise floor: 3.55%. Identical code measured against itself on this setup can look that different, so smaller differences mean nothing.
Neighbouring measurements are correlated (autocorrelation +0.21), so they were resampled in blocks of 4.

Verdict: RESOLVED. A is faster than B: the difference is real, as its interval excludes zero and it exceeds the noise floor.

Warnings:
  - the program holds 115 MB of live data, more than the caches of many machines; for data that size, where it lies in memory can shift the result by several points in ways the interval does not cover, and differently in the next process; run the comparison in several processes with the multiproc package
  - candidate A drifted during the run: its measurements were 5.37% higher in the second half than in the first, so the machine did not hold still

### point lookups of absent keys: tags against lengths, numbers
A = directory of tags (vpage), numbers; B = length header (lpage), numbers

A: 33.45 ns per operation (median)
B: 51.25 ns per operation (median)
A needs 34.7% less time than B: B takes 1.53 times as long as A.
With 95% confidence the difference lies between 33.8% and 35.4% less time (B takes between 1.51 and 1.55 times as long as A).
Noise floor: 2%. Identical code measured against itself on this setup can look that different, so smaller differences mean nothing.
Neighbouring measurements are nearly independent (autocorrelation +0.17), so they were resampled one by one.

Verdict: RESOLVED. A is faster than B: the difference is real, as its interval excludes zero and it exceeds the noise floor.

Warnings:
  - the program holds 116 MB of live data, more than the caches of many machines; for data that size, where it lies in memory can shift the result by several points in ways the interval does not cover, and differently in the next process; run the comparison in several processes with the multiproc package

### point lookups of present keys: numbers against names, both in length-header pages
A = length header (lpage), numbers; B = length header (lpage), names

A: 71.9 ns per operation (median)
B: 68.09 ns per operation (median)
A needs 5.44% more time than B: B takes only 0.95 times as long as A.
With 95% confidence the difference lies between 3.72% and 6.8% more time (B takes between 0.936 and 0.96 times as long as A).
Noise floor: 2.41%. Identical code measured against itself on this setup can look that different, so smaller differences mean nothing.
Neighbouring measurements are nearly independent (autocorrelation +0.14), so they were resampled one by one.

Verdict: RESOLVED. A is slower than B: the difference is real, as its interval excludes zero and it exceeds the noise floor.

Warnings:
  - the program holds 116 MB of live data, more than the caches of many machines; for data that size, where it lies in memory can shift the result by several points in ways the interval does not cover, and differently in the next process; run the comparison in several processes with the multiproc package

### point lookups of present keys: names as bytes against names copied to strings
A = length header (lpage), names; B = length header (lpage), names copied to strings

A: 69.02 ns per operation (median)
B: 74.87 ns per operation (median)
A needs 6.96% less time than B: B takes 1.07 times as long as A.
With 95% confidence the difference lies between 6.2% and 8.02% less time (B takes between 1.07 and 1.09 times as long as A).
Noise floor: 2.21%. Identical code measured against itself on this setup can look that different, so smaller differences mean nothing.
Neighbouring measurements are correlated (autocorrelation +0.25), so they were resampled in blocks of 4.

Verdict: RESOLVED. A is faster than B: the difference is real, as its interval excludes zero and it exceeds the noise floor.

Warnings:
  - the program holds 117 MB of live data, more than the caches of many machines; for data that size, where it lies in memory can shift the result by several points in ways the interval does not cover, and differently in the next process; run the comparison in several processes with the multiproc package
  - candidate A drifted during the run: its measurements were 4.66% lower in the second half than in the first, so the machine did not hold still
  - in 35% of A/A runs this setup reported a difference between identical code, well above the 10% expected; treat any confidence from it with suspicion

| layout | pages | keys per page | page bytes per key |
|---|--:|--:|--:|
| directory of tags (vpage), numbers | 14622 | 14.5 | 35.2 |
| length header (lpage), numbers | 13762 | 15.4 | 27.3 |
| length header (lpage), names | 27428 | 7.7 | 31.0 |

## dirs: 86215 keys (0 too big for a page, 86215 absent keys)

### point lookups of present keys: tags against lengths, numbers
A = directory of tags (vpage), numbers; B = length header (lpage), numbers

A: 62.48 ns per operation (median)
B: 73.07 ns per operation (median)
A needs 15.2% less time than B: B takes 1.18 times as long as A.
With 95% confidence the difference lies between 13.1% and 16.6% less time (B takes between 1.15 and 1.20 times as long as A).
Noise floor: 1.77%. Identical code measured against itself on this setup can look that different, so smaller differences mean nothing.
Neighbouring measurements are nearly independent (autocorrelation +0.14), so they were resampled one by one.

Verdict: RESOLVED. A is faster than B: the difference is real, as its interval excludes zero and it exceeds the noise floor.

Warnings:
  - the program holds 163 MB of live data, more than the caches of many machines; for data that size, where it lies in memory can shift the result by several points in ways the interval does not cover, and differently in the next process; run the comparison in several processes with the multiproc package

### point lookups of absent keys: tags against lengths, numbers
A = directory of tags (vpage), numbers; B = length header (lpage), numbers

A: 48.28 ns per operation (median)
B: 48.21 ns per operation (median)
A needs 0.0154% less time than B: B takes 1.00 times as long as A.
With 95% confidence the difference lies between 1.66% more time and 1.51% less time (B takes between 0.984 and 1.02 times as long as A).
Noise floor: 2.5%. Identical code measured against itself on this setup can look that different, so smaller differences mean nothing.
Neighbouring measurements are nearly independent (autocorrelation +0.09), so they were resampled one by one.

Verdict: NOT RESOLVED. This run did not establish which candidate is faster. That is not the same as "equally fast": the warnings below say what is missing.

Warnings:
  - the difference of 0.02% does not clear the 2.50% noise floor, which is what this setup reports between two runs of identical code
  - the interval [-1.66%, 1.51%] includes zero, so a difference in either direction is consistent with these measurements
  - the program holds 163 MB of live data, more than the caches of many machines; for data that size, where it lies in memory can shift the result by several points in ways the interval does not cover, and differently in the next process; run the comparison in several processes with the multiproc package
  - in 35% of A/A runs this setup reported a difference between identical code, well above the 10% expected; treat any confidence from it with suspicion

### point lookups of present keys: numbers against names, both in length-header pages
A = length header (lpage), numbers; B = length header (lpage), names

A: 73.1 ns per operation (median)
B: 72.94 ns per operation (median)
A needs 0.567% less time than B: B takes 1.01 times as long as A.
With 95% confidence the difference lies between 1.86% more time and 2.11% less time (B takes between 0.982 and 1.02 times as long as A).
Noise floor: 1.64%. Identical code measured against itself on this setup can look that different, so smaller differences mean nothing.
Neighbouring measurements are nearly independent (autocorrelation +0.11), so they were resampled one by one.

Verdict: NOT RESOLVED. This run did not establish which candidate is faster. That is not the same as "equally fast": the warnings below say what is missing.

Warnings:
  - the difference of 0.57% does not clear the 1.64% noise floor, which is what this setup reports between two runs of identical code
  - the interval [-1.86%, 2.11%] includes zero, so a difference in either direction is consistent with these measurements
  - the program holds 163 MB of live data, more than the caches of many machines; for data that size, where it lies in memory can shift the result by several points in ways the interval does not cover, and differently in the next process; run the comparison in several processes with the multiproc package

### point lookups of present keys: names as bytes against names copied to strings
A = length header (lpage), names; B = length header (lpage), names copied to strings

A: 72.9 ns per operation (median)
B: 80.1 ns per operation (median)
A needs 8.59% less time than B: B takes 1.09 times as long as A.
With 95% confidence the difference lies between 7.51% and 9.66% less time (B takes between 1.08 and 1.11 times as long as A).
Noise floor: 1.9%. Identical code measured against itself on this setup can look that different, so smaller differences mean nothing.
Neighbouring measurements are nearly independent (autocorrelation +0.14), so they were resampled one by one.

Verdict: RESOLVED. A is faster than B: the difference is real, as its interval excludes zero and it exceeds the noise floor.

Warnings:
  - the program holds 161 MB of live data, more than the caches of many machines; for data that size, where it lies in memory can shift the result by several points in ways the interval does not cover, and differently in the next process; run the comparison in several processes with the multiproc package
  - in 30% of A/A runs this setup reported a difference between identical code, well above the 10% expected; treat any confidence from it with suspicion

| layout | pages | keys per page | page bytes per key |
|---|--:|--:|--:|
| directory of tags (vpage), numbers | 9814 | 8.8 | 55.9 |
| length header (lpage), numbers | 8697 | 9.9 | 46.6 |
| length header (lpage), names | 12135 | 7.1 | 59.9 |

checksum 17517123388901241254
