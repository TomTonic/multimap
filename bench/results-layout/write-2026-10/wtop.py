# wtop.py SUFFIX VALUES N: top functions by flat share (mean over kinds) with their category
import sys, re, subprocess, collections
sys.path.insert(0, '/tmp/claude-1000/wtools')
import classify
suf, v, n = sys.argv[1], sys.argv[2], sys.argv[3]
KINDS = ['str','email','url','path','street','dirs','links']
acc = collections.Counter()
for k in KINDS:
    out = subprocess.run(['go','tool','pprof','-top','-nodecount=100000',classify.BIN,f'/tmp/claude-1000/wprof/write/{k}-{v}-{n}{suf}.pprof'],capture_output=True,text=True).stdout
    rows=[]; total=0
    for line in out.splitlines():
        m=re.match(r'\s*([\d.]+(?:ns|us|µs|ms|s))\s+[\d.]+%\s+[\d.]+%\s+[\d.]+(?:ns|us|µs|ms|s)\s+[\d.]+%\s+(.*)$',line)
        if not m: continue
        f=classify.unit(m.group(1)); total+=f; rows.append((m.group(2).replace(' (inline)',''),f))
    for name,f in rows: acc[name]+=100*f/total/len(KINDS)
def cat(name):
    for b,rx in classify.B:
        if re.search(rx,name): return b
    return 'runtime other' if name.startswith('runtime.') else 'other'
for name,s in acc.most_common(40):
    print(f'{s:5.1f}  {cat(name):13} {name.replace("github.com/TomTonic/multimap/internal/","")}')
