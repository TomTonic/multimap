import sys
cats = {
 'type byte of the next object': ['5e0b08','5e0b0c'],
 'child pointer in the node (N5,N12,N26,N58,N256)': ['5e0c8e','5e0d2e','5e10c2','5e101b','5e0e6d'],
 'other loads in the node (prefix length, prefix word, key words, bitmap)': ['5e0b17','5e0b78','5e0c46','5e0cb2','5e0cf2','5e0d68','5e0df9'],
}
for f in sys.argv[1:]:
    rows = {}
    for l in open(f):
        if l.startswith('#'): print(l.strip()); continue
        p = l.split()
        rows[p[1]] = float(p[0])
    tot = sum(rows.values()); used = 0
    for c, a in cats.items():
        s = sum(rows.get(x,0) for x in a); used += s
        print(f"  {s:5.1f}  {c}")
    print(f"  {tot-used:5.1f}  rest (switch, loop, SWAR work, page check)")
