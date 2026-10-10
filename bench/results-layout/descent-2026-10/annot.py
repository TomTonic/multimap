# annot.py BIN SIZE PROFILES... : per-instruction flat samples of Tree.find (with its inlined code), summed over the profiles
import re, subprocess, sys, collections
bin_, size, profs = sys.argv[1], sys.argv[2], sys.argv[3:]
asm = subprocess.run(['go','tool','objdump','-s',r'^github.com/TomTonic/multimap/internal/art.\(\*Tree\).find$',bin_],capture_output=True,text=True).stdout
ins = []
for l in asm.splitlines():
    m = re.match(r'\s+(\S+:\d+)\s+0x([0-9a-f]+)\s+[0-9a-f]+\s+(.*)', l)
    if m: ins.append((int(m.group(2),16), m.group(1), m.group(3).strip()))
lo, hi = ins[0][0], ins[-1][0]
cnt = collections.Counter(); tot = 0.0
for p in profs:
    out = subprocess.run(['go','tool','pprof','-top','-addresses','-nodecount=1000000',bin_,p],capture_output=True,text=True).stdout
    for l in out.splitlines():
        m = re.match(r'\s*([\d.]+)(m?s)\s+[\d.]+%\s+[\d.]+%\s+[\d.]+m?s\s+[\d.]+%\s+([0-9a-f]{8,})\s', l)
        if not m: continue
        v = float(m.group(1)) * (0.001 if m.group(2)=='ms' else 1)
        a = int(m.group(3),16)
        tot += v
        if lo <= a <= hi: cnt[a] += v
f = sum(cnt.values())
print(f"# size {size}: find (incl. inlined) {f:.2f}s of {tot:.2f}s = {100*f/tot:.1f}% of the lookup")
for a, src, txt in ins:
    s = cnt.get(a, 0)
    print(f"{100*s/f if f else 0:5.1f}  {a:x}  {src:14s} {txt}")
