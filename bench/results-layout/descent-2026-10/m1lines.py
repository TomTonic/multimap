# m1lines.py LOG: per size, the own time of Tree.find's lines (with inlined code), summed over the kinds, as % of find
import re, sys, collections
txt = open(sys.argv[1]).read()
FIND = ('internal/art/lookup.go', 'internal/art/node.go', 'internal/swar/swar.go')
by = collections.defaultdict(collections.Counter); tot = collections.Counter(); all_ = collections.Counter()
for block in re.split(r'\nlines of ', txt)[1:]:
    head = block.split('\n', 1)[0]
    n = re.search(r'n=(\d+)', head).group(1)
    for l in block.splitlines():
        m = re.match(r'\s*([\d.]+)(m?s)\s+[\d.]+%\s+[\d.]+%\s+[\d.]+m?s\s+[\d.]+%\s+(\S+)\s+(\S+):(\d+)', l)
        if not m: continue
        v = float(m.group(1)) * (0.001 if m.group(2) == 'ms' else 1)
        all_[n] += v
        fn, path, line = m.group(3), m.group(4), int(m.group(5))
        f = next((x for x in FIND if path.endswith(x)), None)
        if f and ('Tree).find' in fn or fn.split('.')[-1] in ('isPage','isSingleKey','childAt','asN5','asN12','asN26','asN58','asN256','endPageOf','singleKeyHdr','Index8','Match8','Has','Rank','Word') ) and not (f.endswith('lookup.go') and line > 95):
            key = f"{fn.split('/')[-1]} {f.split('/')[-1]}:{line}"
            by[n][key] += v; tot[n] += v
for n in sorted(by, key=int):
    print(f'## n={n}: find with inlined code {tot[n]:.2f}s of {all_[n]:.2f}s shown ({100*tot[n]/all_[n]:.0f} %)')
    for k, v in by[n].most_common(25):
        print(f'{100*v/tot[n]:5.1f}  {k}')
