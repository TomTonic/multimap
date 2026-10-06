@echo off
cd /d C:\temp\bench-win
m45-u64.exe -suite dev -minprocs 6 -maxprocs 8 -memn 262144 -memrounds 3 -keys dirs -values natural -sizes 4096,16384,86215 -vs baseline,btree-sets -out C:\temp\bench-win\m45-u64-dirs-real > m45-u64-dirs-real.log 2>&1
echo m45-u64-dirs-real done >> m45-progress.txt
m45-str.exe -suite dev -minprocs 6 -maxprocs 8 -memn 262144 -memrounds 3 -keys dirs -values natural -sizes 4096,16384,86215 -vs baseline,btree-sets -out C:\temp\bench-win\m45-str-dirs-real > m45-str-dirs-real.log 2>&1
echo m45-str-dirs-real done >> m45-progress.txt
m45-ptr.exe -suite dev -minprocs 6 -maxprocs 8 -memn 262144 -memrounds 3 -keys dirs -values natural -sizes 4096,16384,86215 -vs baseline,btree-sets -out C:\temp\bench-win\m45-ptr-dirs-real > m45-ptr-dirs-real.log 2>&1
echo m45-ptr-dirs-real done >> m45-progress.txt
echo done > m45-done.txt
