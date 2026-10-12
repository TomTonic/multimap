import sys, importlib.util
sys.argv = ['x', '/tmp/claude-1000/nc/nodes.txt']
src = open('/tmp/claude-1000/nc/classes.py').read().split('TODAY = ')[0]
exec(src)
D = [(12,(64,128,256,512)), (12,(64,)), (12,(128,)), (12,(64,448)), (12,(64,512)), (12,(128,512)), (12,(64,128,512)), (12,(64,256,512)),
     (12,(64,128,256)), (12,(64,128,192,256,512)), (8,(64,128,192,256,512)), (4,(64,128,192,256,512)), (4,(64,128,256,512)), (8,(64,128,256,512)),
     (20,(64,128,256,512)), (28,(64,128,256,512)), (20,(64,128,256,576)), (28,(64,128,256,576))]
for prof in ('single-value',):
    keys = [c for c in cases if c[1] == prof]
    base = {c: evaluate(12,(64,128,256,512), cases[c]) for c in keys}
    print(f'| P | sizes (+256-way) | classes | slots of the classes | routing B a key (mean) | memory vs today | lines 64 B | lines 128 B |')
    print('|--:|---|--:|---|--:|--:|--:|--:|')
    for P, sz in D:
        H, cls, direct = design(P, sz)
        r = [evaluate(P, sz, cases[c]) for c in keys]
        bk = sum(x[0]/c[2] for x, c in zip(r, keys))/len(keys)
        m = sum(x[0]/base[c][0] for x, c in zip(r, keys))/len(keys)
        l64 = sum(x[1]/base[c][1] for x, c in zip(r, keys))/len(keys)
        l128 = sum(x[2]/base[c][2] for x, c in zip(r, keys))/len(keys)
        print(f'| {P} | {"/".join(map(str,sz))} | {len(cls)+1} | {"/".join(str(c[0]) for c in cls)}/256 | {bk:.1f} | {m:.2f} | {l64:.3f} | {l128:.3f} |')
