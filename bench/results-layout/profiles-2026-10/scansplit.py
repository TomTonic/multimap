import re,subprocess,collections,sys
BIN='/tmp/claude-1000/prof.test'
S=[('walk',r'Cursor\[.*\]\)\.(NextPage|enter|frame|push|leaf|page)$|nextChild|firstChild|touchChildren|art\.\(\*Cursor'),
   ('bounds',r'valueRange|slotRange|cmpKey|Bounds\)\.in|bytes\.Compare|cmpbody'),
   ('values',r'ValuesIn|AppendStrings|AppendKeys|valuesIn|memmove|Cursor\[.*\]\)\.strings|slicebytetostring|mallocgc|AllValuesSeq'),
   ('loop',r'rangeSeq|ValuesBetween|bench/cmd/bench'),
   ('sets',r'Set3|set3|vset')]
def unit(x):
    m=re.match(r'([\d.]+)(ns|us|µs|ms|s)$',x); return float(m.group(1))*{'ns':1e-9,'us':1e-6,'µs':1e-6,'ms':1e-3,'s':1}[m.group(2)]
for prof in sys.argv[1:]:
    out=subprocess.run(['go','tool','pprof','-top','-nodecount=100000',BIN,prof],capture_output=True,text=True).stdout
    c=collections.Counter(); tot=0; unk=collections.Counter()
    for line in out.splitlines():
        m=re.match(r'\s*([\d.]+(?:ns|us|µs|ms|s))\s+[\d.]+%\s+[\d.]+%\s+\S+\s+[\d.]+%\s+(.*)$',line)
        if not m: continue
        f=unit(m.group(1)); name=m.group(2).replace(' (inline)',''); tot+=f
        for b,rx in S:
            if re.search(rx,name): c[b]+=f; break
        else: c['other']+=f; unk[name]+=f
    print(prof.split('/')[-1], ' '.join(f"{k} {100*v/tot:.0f}" for k,v in c.most_common()), '| top other:', ', '.join(f"{n.split('/')[-1]} {100*v/tot:.0f}" for n,v in unk.most_common(3)))
