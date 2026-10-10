import sys, collections
rows = {}; prec = {}
for line in open(sys.argv[1]):
    c = [x.strip() for x in line.strip().strip('|').split('|')]
    if len(c) < 16 or c[0] in ('values', '---'): continue
    vals, keys, n, op, A, B = c[:6]
    rows[(vals, keys, op, int(n), B)] = float(c[9].split('×')[0]); prec[(vals, keys, op, int(n), B)] = c[14]
for vals, bt in (('natural', 'btree-sets'), ('single-value', 'btree-map')):
    print(f'\n{vals}: ordered speed vs main / vs {bt} (>1 = ordered faster; * = not precise)\n')
    print('| keys | op | 4,096 | 16,384 | 65,536 |\n|---|---|---|---|---|')
    for keys in ('street', 'dirs', 'links', 'url'):
        for op in ('valuesFor', 'valuesBetween', 'churn', 'build'):
            cells = []
            for n in (4096, 16384, 65536):
                a = rows.get((vals, keys, op, n, 'baseline')); b = rows.get((vals, keys, op, n, bt))
                s = lambda x, k: '–' if x is None else f"{x:.2f}{'*' if prec.get((vals, keys, op, n, k)) == 'no' else ''}"
                cells.append(f'{s(a, "baseline")} / {s(b, bt)}')
            print(f'| {keys} | {op} | ' + ' | '.join(cells) + ' |')
