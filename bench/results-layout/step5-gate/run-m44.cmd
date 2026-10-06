@echo off
cd /d C:\temp\bench-win
m44-u64.exe -suite dev -minprocs 6 -maxprocs 8 -memn 262144 -memrounds 3 -keys street -values natural -sizes 4096,16384,212449 -vs baseline,btree-sets -out C:\temp\bench-win\m44-u64-street-real > m44-u64-street-real.log 2>&1
echo m44-u64-street-real done >> m44-progress.txt
m44-u64.exe -suite dev -minprocs 6 -maxprocs 8 -memn 262144 -memrounds 3 -keys dirs -values natural -sizes 4096,16384,86215 -vs baseline,btree-sets -out C:\temp\bench-win\m44-u64-dirs-real > m44-u64-dirs-real.log 2>&1
echo m44-u64-dirs-real done >> m44-progress.txt
m44-u64.exe -suite dev -minprocs 6 -maxprocs 8 -memn 262144 -memrounds 3 -keys street -values single-value -sizes 4096,16384,212449 -vs baseline,btree-map -out C:\temp\bench-win\m44-u64-street-single > m44-u64-street-single.log 2>&1
echo m44-u64-street-single done >> m44-progress.txt
m44-u64.exe -suite dev -minprocs 6 -maxprocs 8 -memn 262144 -memrounds 3 -keys dirs -values single-value -sizes 4096,16384,86215 -vs baseline,btree-map -out C:\temp\bench-win\m44-u64-dirs-single > m44-u64-dirs-single.log 2>&1
echo m44-u64-dirs-single done >> m44-progress.txt
m44-u64.exe -suite dev -minprocs 6 -maxprocs 8 -memn 262144 -memrounds 3 -keys u64 -values natural -sizes 4096,16384,262144 -vs baseline,btree-sets -out C:\temp\bench-win\m44-u64-u64keys-real > m44-u64-u64keys-real.log 2>&1
echo m44-u64-u64keys-real done >> m44-progress.txt
m44-u64.exe -suite dev -minprocs 6 -maxprocs 8 -memn 262144 -memrounds 3 -keys u64 -values single-value -sizes 4096,16384,262144 -vs baseline,btree-map -out C:\temp\bench-win\m44-u64-u64keys-single > m44-u64-u64keys-single.log 2>&1
echo m44-u64-u64keys-single done >> m44-progress.txt
m44-str.exe -suite dev -minprocs 6 -maxprocs 8 -memn 262144 -memrounds 3 -keys street -values natural -sizes 4096,16384,212449 -vs baseline,btree-sets -out C:\temp\bench-win\m44-str-street-real > m44-str-street-real.log 2>&1
echo m44-str-street-real done >> m44-progress.txt
m44-str.exe -suite dev -minprocs 6 -maxprocs 8 -memn 262144 -memrounds 3 -keys dirs -values natural -sizes 4096,16384,86215 -vs baseline,btree-sets -out C:\temp\bench-win\m44-str-dirs-real > m44-str-dirs-real.log 2>&1
echo m44-str-dirs-real done >> m44-progress.txt
m44-str.exe -suite dev -minprocs 6 -maxprocs 8 -memn 262144 -memrounds 3 -keys street -values single-value -sizes 4096,16384,212449 -vs baseline,btree-map -out C:\temp\bench-win\m44-str-street-single > m44-str-street-single.log 2>&1
echo m44-str-street-single done >> m44-progress.txt
m44-str.exe -suite dev -minprocs 6 -maxprocs 8 -memn 262144 -memrounds 3 -keys dirs -values single-value -sizes 4096,16384,86215 -vs baseline,btree-map -out C:\temp\bench-win\m44-str-dirs-single > m44-str-dirs-single.log 2>&1
echo m44-str-dirs-single done >> m44-progress.txt
m44-ptr.exe -suite dev -minprocs 6 -maxprocs 8 -memn 262144 -memrounds 3 -keys street -values natural -sizes 4096,16384,212449 -vs baseline,btree-sets -out C:\temp\bench-win\m44-ptr-street-real > m44-ptr-street-real.log 2>&1
echo m44-ptr-street-real done >> m44-progress.txt
m44-ptr.exe -suite dev -minprocs 6 -maxprocs 8 -memn 262144 -memrounds 3 -keys dirs -values natural -sizes 4096,16384,86215 -vs baseline,btree-sets -out C:\temp\bench-win\m44-ptr-dirs-real > m44-ptr-dirs-real.log 2>&1
echo m44-ptr-dirs-real done >> m44-progress.txt
m44-ptr.exe -suite dev -minprocs 6 -maxprocs 8 -memn 262144 -memrounds 3 -keys street -values single-value -sizes 4096,16384,212449 -vs baseline,btree-map -out C:\temp\bench-win\m44-ptr-street-single > m44-ptr-street-single.log 2>&1
echo m44-ptr-street-single done >> m44-progress.txt
m44-ptr.exe -suite dev -minprocs 6 -maxprocs 8 -memn 262144 -memrounds 3 -keys dirs -values single-value -sizes 4096,16384,86215 -vs baseline,btree-map -out C:\temp\bench-win\m44-ptr-dirs-single > m44-ptr-dirs-single.log 2>&1
echo m44-ptr-dirs-single done >> m44-progress.txt
m44-ptr.exe -suite dev -minprocs 6 -maxprocs 8 -memn 262144 -memrounds 3 -keys u64 -values natural -sizes 4096,16384,262144 -vs baseline,btree-sets -out C:\temp\bench-win\m44-ptr-u64keys-real > m44-ptr-u64keys-real.log 2>&1
echo m44-ptr-u64keys-real done >> m44-progress.txt
m44-ptr.exe -suite dev -minprocs 6 -maxprocs 8 -memn 262144 -memrounds 3 -keys u64 -values single-value -sizes 4096,16384,262144 -vs baseline,btree-map -out C:\temp\bench-win\m44-ptr-u64keys-single > m44-ptr-u64keys-single.log 2>&1
echo m44-ptr-u64keys-single done >> m44-progress.txt
echo done > m44-done.txt
