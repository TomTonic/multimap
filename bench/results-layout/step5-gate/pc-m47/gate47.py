import re,os
BASE='/mnt/c/temp/code/multi_map-layout/bench/results-layout/'
R43=BASE+'step4-probe/pc-m43/'
R44=BASE+'step5-gate/pc-m44/'; R45=BASE+'step5-gate/pc-m45/'; R46=BASE+'step5-gate/pc-m46/'; R47=BASE+'step5-gate/pc-m47/'
def ns(s):
    m=re.match(r'([\d.]+)\s*(ns|µs|ms|s)?',s); v=float(m.group(1)); u=m.group(2) or 'ns'
    return v*{'ns':1,'µs':1e3,'ms':1e6,'s':1e9}[u]
def cells(fn):
    for line in open(fn):
        c=[x.strip() for x in line.strip().strip('|').split('|')]
        if len(c)<16 or not c[2].isdigit(): continue
        yield c
def times(fn):
    out={}
    if not os.path.exists(fn): return out
    for c in cells(fn):
        if c[5]=='baseline': continue
        out[(c[3],int(c[2]))]=ns(c[7])
    return out
def over(fn):
    out={}
    if not os.path.exists(fn): return out
    for c in cells(fn):
        m=re.match(r'([\d.]+)×',c[9])
        if m: out[(c[3],int(c[2]),c[5])]=float(m.group(1))
    return out
def mem(fn):
    out={}
    if not os.path.exists(fn): return out
    for line in open(fn):
        c=[x.strip() for x in line.strip().strip('|').split('|')]
        if len(c)<9 or c[0]=='values' or not c[2].isdigit(): continue
        out[c[3]]=(int(c[2]),float(c[5]),float(c[6]),float(c[8]))
    return out
CASES=['street-real','dirs-real','street-single','dirs-single','u64keys-real','u64keys-single']
def sub(t,c): return t+'-'+c
def d43(t,c): return R43+(('u64-' if t=='u64' else 'str-')+c if not c.startswith('u64keys') else c)
def d44(t,c): return (R45 if c=='dirs-real' else R44)+sub(t,c)
def d46(t,c): return R47+sub(t,c)
def d46o(t,c): return R46+sub(t,c)
OPS=['valuesFor','valuesBetween','prefix','churn','build']
def ratio_table(title,da,db,builds):
    print('\n'+title); print('| build | case | '+' | '.join(OPS)+' |'); print('|---|---|'+'--|'*5)
    for t in builds:
        for c in CASES:
            if t=='str' and c.startswith('u64keys'): continue
            a=times(da(t,c)+'/speed-summary.md'); b=times(db(t,c)+'/speed-summary.md')
            row=[]
            for op in OPS:
                rs=[b[(op,n)]/a[(op,n)] for (o,n) in sorted(a) if o==op and (o,n) in b]
                row.append('%.2f–%.2f'%(min(rs),max(rs)) if rs else '–')
            print(f'| {t} | {c} | '+' | '.join(row)+' |')
ratio_table('time m47 over m43 (below 1.00 is faster)',d43,d46,('u64','str'))
ratio_table('time m47 over m46 (below 1.00 is faster)',d46o,d46,('u64','str','ptr'))
ratio_table('time m47 over m44/m45 (below 1.00 is faster)',d44,d46,('u64','str','ptr'))
print('\nmemory: heap B/key and scannable, after removing half. step3.5 baseline | m43 | m44 | m46')
print('| build | case | n | baseline | m43 | m44 | m46 | scannable m44 -> m46 | after half m44 -> m46 |'); print('|---|---|--:|--|--|--|--|--|--|')
for t in ('u64','str','ptr'):
    for c in CASES:
        if t=='str' and c.startswith('u64keys'): continue
        m46=mem(d46(t,c)+'/mem-summary.md'); m44=mem(d44(t,c)+'/mem-summary.md'); m43=mem(d43(t,c)+'/mem-summary.md') if t!='ptr' else {}
        o=m46.get('ordered'); 
        if not o: continue
        b=m46.get('baseline'); p4=m44.get('ordered'); p3=m43.get('ordered')
        f=lambda x:('%.0f'%x[1]) if x else '–'
        print(f'| {t} | {c} | {o[0]} | {f(b)} | {f(p3)} | {f(p4)} | **{f(o)}** | {("%.0f → "%p4[2]) if p4 else ""}{o[2]:.0f} | {("%.0f → "%p4[3]) if p4 else ""}{o[3]:.0f} |')
for kind,vs in (('real','btree-sets'),('single','btree-map')):
    print(f'\ncredo: ordered speed over {vs} (churn / build), m44 -> m46')
    print('| build | keys | 4,096 | 16,384 | large |'); print('|---|---|--|--|--|')
    for t in ('u64','str','ptr'):
        for k in ('street','dirs','u64keys'):
            if t=='str' and k=='u64keys': continue
            c=f'{k}-{kind}'
            r4=over(d44(t,c)+'/speed-summary.md'); r6=over(d46(t,c)+'/speed-summary.md')
            ns_=sorted({n for (_,n,_) in r6})
            cs=[]
            for n in ns_:
                g=lambda r,op:('%.2f'%r[(op,n,vs)]) if (op,n,vs) in r else '–'
                cs.append(f'{g(r4,"churn")} / {g(r4,"build")} → **{g(r6,"churn")} / {g(r6,"build")}**')
            print(f'| {t} | {k} | '+' | '.join(cs)+' |')
print('\ncredo 2: ordered over btree for the ranges, m46')
for t in ('u64','str','ptr'):
    for c in ('street-real','dirs-real','u64keys-real'):
        if t=='str' and c.startswith('u64keys'): continue
        r=over(d46(t,c)+'/speed-summary.md')
        print(t,c,{op:'%.1f–%.1f'%(min(v for (o,n,b),v in r.items() if o==op and b=='btree-sets'),max(v for (o,n,b),v in r.items() if o==op and b=='btree-sets')) for op in ('valuesBetween','prefix') if any(o==op for (o,n,b) in r)})
