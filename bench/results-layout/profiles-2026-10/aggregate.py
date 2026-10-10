import sys,re,os,collections
sys.path.insert(0,'/tmp/claude-1000')
import classify
P='/tmp/claude-1000/prof-all'
KINDS=['u64','str','uuid','email','url','path','street','dirs']
CATS=['descent','page search','page head','page read','page change','tree change','scan','sets','alloc/GC','mem ops','runtime other','harness','other']
SHORT={'descent':'desc','page search':'search','page head':'head','page read':'read','page change':'pchg','tree change':'tree','scan':'scan','sets':'sets','alloc/GC':'GC','mem ops':'mem','runtime other':'rt','harness':'harn','other':'oth'}
def num(path,rx):
    try: t=open(path).read()
    except: return None
    m=re.search(rx,t)
    return float(m.group(1)) if m else None
rows=collections.defaultdict(list)
agg=collections.defaultdict(lambda: collections.defaultdict(list))
for k in KINDS:
  for v in ['single-value','natural']:
    for n in [4096,16384,65536]:
      c=f"{k}-{v}-{n}"
      o=f"{P}/out"
      ops={
       'churn':(f"{P}/write/{c}.pprof", num(f"{o}/probe-{c}.txt",r'plain replay of a cycle: ([\d.]+) ns'), None),
       'build':(f"{P}/write/{c}-build.pprof", num(f"{o}/probe-{c}.txt",r'build stream: ([\d.]+) ns'), None),
       'lookup':(f"{P}/look/{c}-ordered.pprof", num(f"{o}/look-{c}.txt",r'ordered ([\d.]+) ns'), num(f"{o}/look-{c}.txt",r'baseline \(main [0-9a-f]+\) ([\d.]+) ns')),
       'range':(f"{P}/range/{c}-range.pprof", num(f"{o}/range-{c}.txt",r'ordered ([\d.]+) ns a value'), num(f"{o}/range-{c}.txt",r'btree-map ([\d.]+) ns a value')),
      }
      for op,(prof,ns,ref) in ops.items():
        if not os.path.exists(prof): continue
        tot,total=classify.classify(prof)
        if total == 0: continue
        sh={k2:100*tot[k2]/total for k2 in CATS}
        rows[op].append((k,v,n,ns,ref,total,sh))
        agg[op][(v,n)].append(sh)
out=[]
for op,ref in [('lookup','main'),('churn',None),('build',None),('range','btree-map')]:
  out.append(f"\n### {op} (ns {'a value' if op=='range' else 'an operation'}; shares of the CPU time in %)\n")
  hdr=f"| keys | values | n | "+" | ".join(SHORT[c] for c in CATS)+" | s |"
  out.append(hdr); out.append("|"+"---|"*(hdr.count('|')-1))
  for (k,v,n,ns,r,total,sh) in rows[op]:
    line=f"| {k} | {v} | {n} | "+" | ".join(f"{sh[c]:.0f}" for c in CATS)+f" | {total:.1f} |"
    out.append(line)
  out.append(f"\nmean over the key kinds:\n")
  out.append("| values | n | "+" | ".join(SHORT[c] for c in CATS)+" |"); out.append("|"+"---|"*(len(CATS)+2))
  for (v,n),lst in sorted(agg[op].items()):
    out.append(f"| {v} | {n} | "+" | ".join(f"{sum(s[c] for s in lst)/len(lst):.0f}" for c in CATS)+" |")
print("\n".join(out))
