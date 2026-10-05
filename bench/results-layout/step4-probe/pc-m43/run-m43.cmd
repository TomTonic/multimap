@echo off
cd /d C:\temp\bench-win
m43-u64.exe -suite dev -minprocs 6 -maxprocs 8 -memn 262144 -memrounds 3 -keys street -values natural -sizes 4096,16384,212449 -vs baseline,btree-sets -out C:\temp\bench-win\m43-u64-street-real > m43-u64-street-real.log 2>&1
echo m43-u64-street-real done >> m43-progress.txt
m43-u64.exe -suite dev -minprocs 6 -maxprocs 8 -memn 262144 -memrounds 3 -keys dirs -values natural -sizes 4096,16384,86215 -vs baseline,btree-sets -out C:\temp\bench-win\m43-u64-dirs-real > m43-u64-dirs-real.log 2>&1
echo m43-u64-dirs-real done >> m43-progress.txt
m43-u64.exe -suite dev -minprocs 6 -maxprocs 8 -memn 262144 -memrounds 3 -keys street -values single-value -sizes 4096,16384,212449 -vs baseline,btree-map -out C:\temp\bench-win\m43-u64-street-single > m43-u64-street-single.log 2>&1
echo m43-u64-street-single done >> m43-progress.txt
m43-u64.exe -suite dev -minprocs 6 -maxprocs 8 -memn 262144 -memrounds 3 -keys dirs -values single-value -sizes 4096,16384,86215 -vs baseline,btree-map -out C:\temp\bench-win\m43-u64-dirs-single > m43-u64-dirs-single.log 2>&1
echo m43-u64-dirs-single done >> m43-progress.txt
m43-str.exe -suite dev -minprocs 6 -maxprocs 8 -memn 262144 -memrounds 3 -keys street -values natural -sizes 4096,16384,212449 -vs baseline,btree-sets -out C:\temp\bench-win\m43-str-street-real > m43-str-street-real.log 2>&1
echo m43-str-street-real done >> m43-progress.txt
m43-str.exe -suite dev -minprocs 6 -maxprocs 8 -memn 262144 -memrounds 3 -keys dirs -values natural -sizes 4096,16384,86215 -vs baseline,btree-sets -out C:\temp\bench-win\m43-str-dirs-real > m43-str-dirs-real.log 2>&1
echo m43-str-dirs-real done >> m43-progress.txt
m43-str.exe -suite dev -minprocs 6 -maxprocs 8 -memn 262144 -memrounds 3 -keys street -values single-value -sizes 4096,16384,212449 -vs baseline,btree-map -out C:\temp\bench-win\m43-str-street-single > m43-str-street-single.log 2>&1
echo m43-str-street-single done >> m43-progress.txt
m43-str.exe -suite dev -minprocs 6 -maxprocs 8 -memn 262144 -memrounds 3 -keys dirs -values single-value -sizes 4096,16384,86215 -vs baseline,btree-map -out C:\temp\bench-win\m43-str-dirs-single > m43-str-dirs-single.log 2>&1
echo m43-str-dirs-single done >> m43-progress.txt
m43-u64.exe -suite dev -minprocs 6 -maxprocs 8 -memn 262144 -memrounds 3 -keys u64 -values natural -sizes 4096,16384,262144 -vs baseline,btree-sets -out C:\temp\bench-win\m43-u64keys-real > m43-u64keys-real.log 2>&1
echo m43-u64keys-real done >> m43-progress.txt
m43-u64.exe -suite dev -minprocs 6 -maxprocs 8 -memn 262144 -memrounds 3 -keys u64 -values single-value -sizes 4096,16384,262144 -vs baseline,btree-map -out C:\temp\bench-win\m43-u64keys-single > m43-u64keys-single.log 2>&1
echo m43-u64keys-single done >> m43-progress.txt
echo done > m43-done.txt
