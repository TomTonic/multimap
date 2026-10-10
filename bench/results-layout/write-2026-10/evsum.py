# evsum.py files...: merge tries and pairs per case, from TestProbe output with MKPROBE_HIST and mkstats
import re, sys
NEXT = {32: 64, 64: 128, 128: 256, 256: 384, 384: 512, 512: 512}
def hist(block, key):
    m = re.search(r'histogram \| ' + re.escape(key) + r'[^|]*\|(.*)', block)
    if not m: return {}
    return {int(a): int(b) for a, b in re.findall(r'(\d+):(\d+)', m.group(1))}
def tot(h): return sum(h.values())
def med(h):
    t, s = tot(h), 0
    for k in sorted(h):
        s += h[k]
        if 2 * s >= t: return k
    return -1
print('| case | merge up a 1,000 writes | from 1 entry left | tries a merge up | refused: >12 children / child no page / >64 values / too big | too big: need / mergeFill median, share under 1.5x | merges done | pairs a 1,000 writes | pair fits the old object | single-key pages one class larger: B a key |')
print('|---|--:|--:|--:|---|---|--:|--:|--:|--:|')
for f in sys.argv[1:]:
    txt = open(f).read()
    for case in re.split(r'\n## ', txt)[1:]:
        name = case.split(' (')[0]
        m = re.search(r'events of one steady-state cycle \((\d+) operations\):(.*?)\n\n(histogram.*?)\n\n', case, re.S)
        if not m: continue
        ops, block = int(m.group(1)), m.group(3)
        up = hist(block, 'merge up started'); tries = hist(block, 'merge tried'); ok = hist(block, 'merge done')
        wide = hist(block, 'merge refused: the node has more'); nopage = hist(block, 'merge refused: a child is')
        slots = hist(block, 'merge refused: the children hold more'); big = hist(block, 'merge refused: the merged page')
        pair = hist(block, 'pair: single-key page and new key'); fit = hist(block, 'pair: the page of two keys')
        nup = tot(up); ref = [tot(wide), tot(nopage), tot(slots), tot(big)]; allref = sum(ref) or 1
        under = sum(v for k, v in big.items() if k <= 96)
        # census after one cycle
        c = re.search(r'objects of the tree after one cycle:.*?\n\n(.*?)\n\n', case, re.S)
        keys = sum(int(r.split('|')[4]) for r in c.group(1).splitlines()[2:]) if c else 0
        s = re.search(r'tree after one cycle.*?single-key pages by size of the object:([^\n]*)', case, re.S)
        extra = sum((NEXT[int(a)] - int(a)) * int(b) for a, b in re.findall(r'(\d+): (\d+);', s.group(1))) if s else 0
        print(f'| {name} | {1000*nup/ops:.0f} | {100*up.get(1,0)/max(nup,1):.0f} % | {(tot(tries)+tot(wide))/max(nup,1):.2f} | '
              f'{100*ref[0]/allref:.0f} / {100*ref[1]/allref:.0f} / {100*ref[2]/allref:.0f} / {100*ref[3]/allref:.0f} % | '
              f'{med(big)/64:.1f}x, {100*under/max(tot(big),1):.0f} % | {tot(ok)} | {1000*tot(pair)/ops:.0f} | '
              f'{100*sum(v for k,v in fit.items() if k<=64)/max(tot(fit),1):.0f} % | {extra/max(keys,1):.1f} |')
