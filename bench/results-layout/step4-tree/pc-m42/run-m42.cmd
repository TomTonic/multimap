@echo off
cd /d C:\temp\bench-win
m42-u64.exe -suite dev -minprocs 6 -maxprocs 8 -memn 262144 -memrounds 3 -keys street -values natural -sizes 4096,16384,212449 -vs baseline,btree-sets -out C:\temp\bench-win\m42-u64-street-real > m42-u64-street-real.log 2>&1
echo m42-u64-street-real done >> m42-progress.txt
m42-u64.exe -suite dev -minprocs 6 -maxprocs 8 -memn 262144 -memrounds 3 -keys dirs -values natural -sizes 4096,16384,86215 -vs baseline,btree-sets -out C:\temp\bench-win\m42-u64-dirs-real > m42-u64-dirs-real.log 2>&1
echo m42-u64-dirs-real done >> m42-progress.txt
m42-u64.exe -suite dev -minprocs 6 -maxprocs 8 -memn 262144 -memrounds 3 -keys street -values single-value -sizes 4096,16384,212449 -vs baseline,btree-map -out C:\temp\bench-win\m42-u64-street-single > m42-u64-street-single.log 2>&1
echo m42-u64-street-single done >> m42-progress.txt
m42-u64.exe -suite dev -minprocs 6 -maxprocs 8 -memn 262144 -memrounds 3 -keys dirs -values single-value -sizes 4096,16384,86215 -vs baseline,btree-map -out C:\temp\bench-win\m42-u64-dirs-single > m42-u64-dirs-single.log 2>&1
echo m42-u64-dirs-single done >> m42-progress.txt
m42-str.exe -suite dev -minprocs 6 -maxprocs 8 -memn 262144 -memrounds 3 -keys street -values natural -sizes 4096,16384,212449 -vs baseline,btree-sets -out C:\temp\bench-win\m42-str-street-real > m42-str-street-real.log 2>&1
echo m42-str-street-real done >> m42-progress.txt
m42-str.exe -suite dev -minprocs 6 -maxprocs 8 -memn 262144 -memrounds 3 -keys dirs -values natural -sizes 4096,16384,86215 -vs baseline,btree-sets -out C:\temp\bench-win\m42-str-dirs-real > m42-str-dirs-real.log 2>&1
echo m42-str-dirs-real done >> m42-progress.txt
m42-str.exe -suite dev -minprocs 6 -maxprocs 8 -memn 262144 -memrounds 3 -keys street -values single-value -sizes 4096,16384,212449 -vs baseline,btree-map -out C:\temp\bench-win\m42-str-street-single > m42-str-street-single.log 2>&1
echo m42-str-street-single done >> m42-progress.txt
m42-str.exe -suite dev -minprocs 6 -maxprocs 8 -memn 262144 -memrounds 3 -keys dirs -values single-value -sizes 4096,16384,86215 -vs baseline,btree-map -out C:\temp\bench-win\m42-str-dirs-single > m42-str-dirs-single.log 2>&1
echo m42-str-dirs-single done >> m42-progress.txt
m42-u64.exe -suite dev -minprocs 6 -maxprocs 8 -memn 262144 -memrounds 3 -keys u64 -values natural -sizes 4096,16384,262144 -vs baseline,btree-sets -out C:\temp\bench-win\m42-u64keys-real > m42-u64keys-real.log 2>&1
echo m42-u64keys-real done >> m42-progress.txt
m42-u64.exe -suite dev -minprocs 6 -maxprocs 8 -memn 262144 -memrounds 3 -keys u64 -values single-value -sizes 4096,16384,262144 -vs baseline,btree-map -out C:\temp\bench-win\m42-u64keys-single > m42-u64keys-single.log 2>&1
echo m42-u64keys-single done >> m42-progress.txt
echo done > m42-done.txt
