import sys, collections
sys.path.insert(0, '/tmp/claude-1000/wtools')
import classify
P = '/tmp/claude-1000/wprof/write'
KINDS = ['str','email','url','path','street','dirs','links']
CATS = ['descent','page search','page head','page read','page change','tree change','sets','alloc/GC','mem ops','runtime other','harness','other']
SH = {'descent':'desc','page search':'search','page head':'head','page read':'read','page change':'pchg','tree change':'tree','sets':'sets','alloc/GC':'GC','mem ops':'mem','runtime other':'rt','harness':'harn','other':'oth'}
print('| op | values | n | kind | ' + ' | '.join(SH[c] for c in CATS) + ' |')
print('|---|---|--:|---|' + '--:|'*len(CATS))
for op, suf in (('churn',''),('build','-build')):
  for v in ['single-value','natural']:
    for n in [4096,16384]:
      acc = collections.defaultdict(list)
      for k in KINDS:
        tot, total = classify.classify(f'{P}/{k}-{v}-{n}{suf}.pprof')
        sh = {c: 100*tot[c]/total for c in CATS}
        for c in CATS: acc[c].append(sh[c])
        print(f'| {op} | {v} | {n} | {k} | ' + ' | '.join(f'{sh[c]:.0f}' for c in CATS) + ' |')
      print(f'| **{op}** | **{v}** | **{n}** | **mean** | ' + ' | '.join(f'**{sum(acc[c])/len(acc[c]):.0f}**' for c in CATS) + ' |')
print()
for k, v in classify.unk.most_common(15): print(f'unclassified {k} {v:.2f}s')
