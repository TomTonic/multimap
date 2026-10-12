# classes.py nodes.txt: evaluate node designs (inline prefix P, node sizes) on today's tree shape; see node-classes-analysis.md
import sys, itertools, collections, math
GO = [8,16,24,32,48,64,80,96,112,128,144,160,176,192,208,224,240,256,288,320,352,384,416,448,480,512,576,640,704,768,896,1024,
      1152,1280,1408,1536,1792,2048,2304,2688,3072,3200,3456,4096]
def goclass(b):
    for c in GO:
        if c >= b: return c
    return (b + 8191) // 8192 * 8192
def a8(x): return (x + 7) // 8 * 8
# --- read and aggregate: per case, groups (slots, end, plen) -> [count, prefix sums of keys by slot index (end page last)]
cases = collections.OrderedDict(); cur = None
for line in open(sys.argv[1]):
    f = line.split()
    if f[0] == 'case':
        cur = (f[1], f[2], int(f[3])); cases[cur] = {}
        continue
    plen, end, ks = int(f[1]), int(f[2]), list(map(int, f[3:]))
    g = cases[cur].setdefault((len(ks), end, plen), [0, [0] * len(ks)])
    g[0] += 1
    for j, k in enumerate(ks): g[1][j] += k
for c in cases.values():
    for g in c.values():
        ps = [0]
        for k in g[1]: ps.append(ps[-1] + k)
        g.append(ps)  # g[2] prefix sums
TAILS = (16, 48, 112)
def design(P, sizes):
    H = 4 + P
    cls = []
    for s in sorted(sizes):
        best = None
        for c in range(1, 25):  # sorted node
            if a8(H + c) + 8 * c <= s: best = (c, 'sorted', a8(H + c))
        cb = (s - a8(H + 32)) // 8
        if cb > 0 and (best is None or cb > best[0]): best = (cb, 'bitmap', a8(H + 32))
        if best: cls.append((best[0], s, best[1], best[2]))
    cls.sort()
    direct = (257, a8(H + 8) + 257 * 8, 'direct', a8(H + 8))
    return H, cls, direct
def evaluate(P, sizes, case):
    H, cls, direct = design(P, sizes)
    mem = lines64 = lines128 = 0.0
    for (s, end, plen), (cnt, ks, ps) in case.items():
        cap, size, kind, base = next((c for c in cls if c[0] >= s), direct)
        need = plen - P
        tail = 0; extra = 0
        if need > 0:
            tail = next((t for t in TAILS if t >= need), 16)
            if need > 112: extra = goclass(need)
        mem += cnt * (goclass(size + tail) + extra)
        K = ps[-1]
        for L, acc in ((64, 'l64'), (128, 'l128')):
            # keys whose child slot is outside the first line: byte slots j with base+8j+8 > L, the end page in slot cap-1
            nb = s - end
            if kind == 'direct':
                far = K * (1 - min(1.0, L / size))
            else:
                first = max(0, min(nb, (L - base) // 8))  # byte slots in the first line
                far = ps[nb] - ps[first]
                if end and base + 8 * (cap - 1) + 8 > L: far += ks[-1]
            t = K if (need > 0 and (size // L) > 0) else 0  # the tail starts at size: another line unless the node is shorter than a line
            v = K + far + t
            if L == 64: lines64 += v
            else: lines128 += v
    return mem, lines64, lines128
TODAY = (12, (64, 128, 256, 512))
MENU = (64, 96, 128, 192, 256, 384, 512, 768, 1024)
designs = [(P, sz) for P in (4, 8, 12, 16, 20, 24, 28) for k in range(1, 6) for sz in itertools.combinations(MENU, k)]
if TODAY not in designs: designs.append(TODAY)
res = {}
for prof in ('single-value', 'natural'):
    keys = [c for c in cases if c[1] == prof]
    base = {c: evaluate(*TODAY, cases[c]) for c in keys}
    out = []
    for d in designs:
        r = [evaluate(*d, cases[c]) for c in keys]
        m = sum(x[0] / base[c][0] for x, c in zip(r, keys)) / len(keys)
        l64 = sum(x[1] / base[c][1] for x, c in zip(r, keys)) / len(keys)
        l128 = sum(x[2] / base[c][2] for x, c in zip(r, keys)) / len(keys)
        out.append((m, l64, l128, d))
    res[prof] = (out, base, keys)
    n = {c[2] for c in keys}
    print(f'\n## {prof}: today (P=12, 64/128/256/512): routing bytes a key and lines a lookup (64 / 128 B), by case')
    for c in keys:
        b = base[c]; print(f'  {c[0]:7} {c[2]:6}: {b[0]/c[2]:6.1f} B/key  {b[1]/c[2]:5.2f} / {b[2]/c[2]:5.2f} lines')
    # Pareto front over (memory, lines64)
    out.sort()
    front, best = [], 9e9
    for x in out:
        if x[1] < best - 1e-9: front.append(x); best = x[1]
    print(f'\n### Pareto front (memory and 64-B lines, mean ratio to today over the {len(keys)} cases)')
    print('| P | sizes | memory | lines 64 B | lines 128 B |\n|--:|---|--:|--:|--:|')
    for m, l64, l128, (P, sz) in front[:40]:
        print(f'| {P} | {"/".join(map(str, sz))} | {m:.3f} | {l64:.3f} | {l128:.3f} |')
    print('\n### best by P (lowest memory with lines 64 B at most today\'s)')
    for P in (4, 8, 12, 16, 20, 24, 28):
        cand = [x for x in out if x[3][0] == P and x[1] <= 1.0001]
        if cand:
            m, l64, l128, d = min(cand)
            print(f'  P={P:2}: {"/".join(map(str, d[1])):24} memory {m:.3f}  lines {l64:.3f} / {l128:.3f}')
    t = [x for x in out if x[3] == TODAY][0]
    print(f'  today: memory {t[0]:.3f} lines {t[1]:.3f} / {t[2]:.3f}')
