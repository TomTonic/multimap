# lifesum.py life.txt: one row per case and stream
import re, sys
txt = open(sys.argv[1]).read()
rows = []
for case in re.split(r'\n## life ', txt)[1:]:
    name = case.split(' (')[0]
    for stream, rt in ((r'\nchurn \(one steady-state cycle\)', 'churn'), (r'\nbuild stream from empty', 'build')):
        m = re.search(stream + r': (\d+) writes; in place ([\d.]+) %, one object for one ([\d.]+) %, other ([\d.]+) %\n'
            r'  objects made ([\d.]+) a write \(([\d.]+) bytes\), dropped ([\d.]+) \(([\d.]+) bytes\)\n'
            r'  LIFO free list: ([\d.]+) % .*?median (-?\d+), p90 (-?\d+); bytes made in between median (-?\d+), p90 (-?\d+); held at most (\d+) bytes', case)
        r = re.search(r'runtime, ' + rt + r': ([\d.]+) allocations and ([\d.]+) bytes', case)
        if not m: continue
        g = m.groups()
        rows.append((name, rt, g[1], g[2], g[3], g[4], g[5], g[8], g[9], g[10], g[11], g[12], g[13], r.group(1) if r else '?', r.group(2) if r else '?'))
print('| case | stream | in place % | 1 for 1 % | other % | objects made /w | bytes made /w | LIFO hit % | dist writes med / p90 | bytes between med / p90 | held max B | runtime allocs /w | runtime B /w |')
print('|---|---|--:|--:|--:|--:|--:|--:|--:|--:|--:|--:|--:|')
for r in rows:
    print(f'| {r[0]} | {r[1]} | {r[2]} | {r[3]} | {r[4]} | {r[5]} | {r[6]} | {r[7]} | {r[8]} / {r[9]} | {r[10]} / {r[11]} | {r[12]} | {r[13]} | {r[14]} |')
