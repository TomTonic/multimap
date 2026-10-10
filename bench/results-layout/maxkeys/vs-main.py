import sys,re,collections
def load(p):
    rows={}
    prec={}
    for line in open(p):
        c=[x.strip() for x in line.strip().strip('|').split('|')]
        if len(c)<16 or c[0] in('values','---'): continue
        vals,keys,n,op,A,B=c[:6]
        r=float(c[9].split('×')[0])
        rows[(vals,keys,int(n),op,B)]=r
        prec[(vals,keys,int(n),op,B)]=c[14]
    return rows,prec
rows,prec=load(sys.argv[1])
cells=sorted({k[:4] for k in rows}, key=lambda k:(k[0],k[1],k[3],k[2]))
mk=['ordered-mk1','ordered-mk2','ordered-mk4']
print("speed vs main (baseline); >1 faster than main. bt = btree-sets/btree-map vs main. '*' = not precise")
print(f"{'values':12} {'keys':6} {'op':13} {'n':>6} | {'ordered':>8} {'mk1':>7} {'mk2':>7} {'mk4':>7} | {'btree':>6}")
for cell in cells:
    vals,keys,n,op=cell
    base=rows.get(cell+('baseline',))
    if base is None: continue
    out=[base]
    for m in mk:
        x=rows.get(cell+(m,))
        out.append(base/x if x else float('nan'))
    bt=rows.get(cell+('btree-sets',)) or rows.get(cell+('btree-map',))
    btv=base/bt if bt else float('nan')
    star=''.join('*' if prec.get(cell+(b,))=='no' else ' ' for b in ['baseline']+mk)
    print(f"{vals:12} {keys:6} {op:13} {n:>6} | {out[0]:8.2f} {out[1]:7.2f} {out[2]:7.2f} {out[3]:7.2f} | {btv:6.2f} {star}")
