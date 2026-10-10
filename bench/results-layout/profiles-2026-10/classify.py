import sys,re,subprocess,os,glob,collections
BIN='/tmp/claude-1000/prof.test'
B=[('descent',r'art\.\(\*Tree\)\.find$|art\.findLoc|swar\.(MatchPrefix|Match8|matchLong|Index8|Rank|Has|Word)|art\.prefixLcp|art\.isPage|art\.childAt|art\.as(N5|N12|N26|N58|N256)\b|art\.bitmapOf|art\.sorted|prefixLen|prefixMatches|art\.endPageOf|art\.findSlot|art\.isSingleKey|art\.isMultiKey|art\.asSingleKey|art\.singleKeyHdr'),
   ('page search',r'page\.locate|page\.compare|page\.lcp|page\.\(\*head\)\.Match|keyEndFrom|runEnd|hasValue|page\.indexOf|\(\*singleKeyHead\)\.(matches|stored|rem)'),
   ('scan',r'Cursor|nextChild|firstChild|slotRange|valueRange|ValuesIn|AppendStrings|AppendKeys|cmpKey|Bounds\)\.in|touchChildren|rangeSeq|AllValuesSeq|AllKeysSeq'),
   ('page head',r'\(\*head\)\.(mem|cpl|lay|one|class|Len|Size|setCpl|CP|PrefixLen)|page\.sum$|page\.b2i'),
   ('page read',r'page\.\.?(EachValue|Get|EachSingle|Each)\b|page\.\(\*(Str|Fixed)\)\.(EachValue|Get|EachSingle|Each)|valuesIn|art\.eachValue|art\.\(\*Map\[.*\]\)\.(Each|pageEach|Has|pageHas)|ValuesForSeq|art\.view|art\.strOf|fromStr'),
   ('page change',r'page\.\.?(Add|Remove)\b|page\.\(\*(Str|Fixed)\)\.(Add|Remove|Widen|Skip|Prepend|regrow|pairWith|Used)|page\.(BuildStrings|BuildFixedOf|BuildFixed|NewStr|NewFixed|toOneKey|oneKeyLeft|shrinkClass|classFor|allocRaw|allocPtr|newFixed|regrow|pairWithFixed|further|NeedStrings|NeedFixed)|\(\*head\)\.(setHead|setCpl|KeysUpTo|keys|Keys)|page\.\(\*lay\)|page\.b2i'),
   ('sets',r'Set3|set3\.|vset\.|overflow'),
   ('tree change',r'multimap/internal/art\.'),
   ('alloc/GC',r'runtime\.(mallocgc|nextFreeFast|\(\*mcache\)|\(\*mcentral\)|\(\*mspan\)|\(\*mheap\)|gcBgMarkWorker|gcDrain|scanobject|greyobject|findObject|markBits|memclrNoHeapPointers|heapBits|sweep|bgsweep|wbBuf|bulkBarrier|gcWriteBarrier|typedmemmove|typedslicecopy|newobject|makeslice|growslice|gcmark|scanblock|spanOf|pageIndexOf|deductAssist|gcAssist|markroot|scanstack|publicationBarrier|heapSetType|mallocgcSmall|mallocgcTiny|mallocgcLarge|memclr)'),
   ('mem ops',r'runtime\.(memmove|memequal|memeqbody|cmpbody|cmpstring)|^memeqbody|^cmpbody|bytes\.(Equal|Compare)|internal/bytealg'),
   ('harness',r'bench/cmd/bench|rtcompare|workload|testing\.'),
]
def unit(x):
    m=re.match(r'([\d.]+)(ns|us|µs|ms|s)$',x)
    v=float(m.group(1)); u=m.group(2)
    return v*{'ns':1e-9,'us':1e-6,'µs':1e-6,'ms':1e-3,'s':1}[u]
unk=collections.Counter()
def classify(prof):
    out=subprocess.run(['go','tool','pprof','-top','-nodecount=100000',BIN,prof],capture_output=True,text=True).stdout
    tot=collections.Counter(); total=0
    for line in out.splitlines():
        m=re.match(r'\s*([\d.]+(?:ns|us|µs|ms|s))\s+[\d.]+%\s+[\d.]+%\s+[\d.]+(?:ns|us|µs|ms|s)\s+[\d.]+%\s+(.*)$',line)
        if not m: continue
        f=unit(m.group(1)); name=m.group(2).replace(' (inline)','')
        total+=f
        for b,rx in B:
            if re.search(rx,name): tot[b]+=f; break
        else:
            k='runtime other' if name.startswith('runtime.') else 'other'
            tot[k]+=f
            unk[name]+=f
    return tot,total
if __name__=='__main__':
    tot,total=classify(sys.argv[1])
    print(f"total {total:.2f}s")
    for k,v in tot.most_common(): print(f"{k:14} {100*v/total:5.1f}%")
    for k,v in unk.most_common(8): print(f"   unclassified {k} {v:.3f}s")
