@echo off
cd /d C:\temp\bench-win
m35x-u64.exe -suite dev -minprocs 6 -maxprocs 8 -memn 262144 -memrounds 3 -keys street -values natural -sizes 4096,16384,212449 -vs baseline,btree-sets -out C:\temp\bench-win\m35-x-u64-street-real > m35-x-u64-street-real.log 2>&1
echo m35-x-u64-street-real done >> m35-progress.txt
m35x-u64.exe -suite dev -minprocs 6 -maxprocs 8 -memn 262144 -memrounds 3 -keys dirs -values natural -sizes 4096,16384,86215 -vs baseline,btree-sets -out C:\temp\bench-win\m35-x-u64-dirs-real > m35-x-u64-dirs-real.log 2>&1
echo m35-x-u64-dirs-real done >> m35-progress.txt
m35x-u64.exe -suite dev -minprocs 6 -maxprocs 8 -memn 262144 -memrounds 3 -keys street -values single-value -sizes 4096,16384,212449 -vs baseline,btree-map -out C:\temp\bench-win\m35-x-u64-street-single > m35-x-u64-street-single.log 2>&1
echo m35-x-u64-street-single done >> m35-progress.txt
m35x-u64.exe -suite dev -minprocs 6 -maxprocs 8 -memn 262144 -memrounds 3 -keys dirs -values single-value -sizes 4096,16384,86215 -vs baseline,btree-map -out C:\temp\bench-win\m35-x-u64-dirs-single > m35-x-u64-dirs-single.log 2>&1
echo m35-x-u64-dirs-single done >> m35-progress.txt
m35x-ptr.exe -suite dev -minprocs 6 -maxprocs 8 -memn 262144 -memrounds 3 -keys street -values natural -sizes 4096,16384,212449 -vs baseline,btree-sets -out C:\temp\bench-win\m35-x-ptr-street-real > m35-x-ptr-street-real.log 2>&1
echo m35-x-ptr-street-real done >> m35-progress.txt
m35x-ptr.exe -suite dev -minprocs 6 -maxprocs 8 -memn 262144 -memrounds 3 -keys dirs -values natural -sizes 4096,16384,86215 -vs baseline,btree-sets -out C:\temp\bench-win\m35-x-ptr-dirs-real > m35-x-ptr-dirs-real.log 2>&1
echo m35-x-ptr-dirs-real done >> m35-progress.txt
m35x-ptr.exe -suite dev -minprocs 6 -maxprocs 8 -memn 262144 -memrounds 3 -keys street -values single-value -sizes 4096,16384,212449 -vs baseline,btree-map -out C:\temp\bench-win\m35-x-ptr-street-single > m35-x-ptr-street-single.log 2>&1
echo m35-x-ptr-street-single done >> m35-progress.txt
m35x-ptr.exe -suite dev -minprocs 6 -maxprocs 8 -memn 262144 -memrounds 3 -keys dirs -values single-value -sizes 4096,16384,86215 -vs baseline,btree-map -out C:\temp\bench-win\m35-x-ptr-dirs-single > m35-x-ptr-dirs-single.log 2>&1
echo m35-x-ptr-dirs-single done >> m35-progress.txt
m35x-ptr.exe -suite dev -minprocs 6 -maxprocs 8 -memn 262144 -memrounds 3 -keys u64 -values natural -sizes 4096,16384,262144 -vs baseline,btree-sets -out C:\temp\bench-win\m35-x-ptr-u64keys-real > m35-x-ptr-u64keys-real.log 2>&1
echo m35-x-ptr-u64keys-real done >> m35-progress.txt
m35x-ptr.exe -suite dev -minprocs 6 -maxprocs 8 -memn 262144 -memrounds 3 -keys u64 -values single-value -sizes 4096,16384,262144 -vs baseline,btree-map -out C:\temp\bench-win\m35-x-ptr-u64keys-single > m35-x-ptr-u64keys-single.log 2>&1
echo m35-x-ptr-u64keys-single done >> m35-progress.txt
m35y-u64.exe -suite dev -minprocs 6 -maxprocs 8 -memn 262144 -memrounds 3 -keys street -values natural -sizes 4096,16384,212449 -vs baseline -out C:\temp\bench-win\m35-y-u64-street-real > m35-y-u64-street-real.log 2>&1
echo m35-y-u64-street-real done >> m35-progress.txt
m35y-u64.exe -suite dev -minprocs 6 -maxprocs 8 -memn 262144 -memrounds 3 -keys dirs -values natural -sizes 4096,16384,86215 -vs baseline -out C:\temp\bench-win\m35-y-u64-dirs-real > m35-y-u64-dirs-real.log 2>&1
echo m35-y-u64-dirs-real done >> m35-progress.txt
m35y-u64.exe -suite dev -minprocs 6 -maxprocs 8 -memn 262144 -memrounds 3 -keys street -values single-value -sizes 4096,16384,212449 -vs baseline -out C:\temp\bench-win\m35-y-u64-street-single > m35-y-u64-street-single.log 2>&1
echo m35-y-u64-street-single done >> m35-progress.txt
m35y-u64.exe -suite dev -minprocs 6 -maxprocs 8 -memn 262144 -memrounds 3 -keys dirs -values single-value -sizes 4096,16384,86215 -vs baseline -out C:\temp\bench-win\m35-y-u64-dirs-single > m35-y-u64-dirs-single.log 2>&1
echo m35-y-u64-dirs-single done >> m35-progress.txt
m35y-ptr.exe -suite dev -minprocs 6 -maxprocs 8 -memn 262144 -memrounds 3 -keys street -values natural -sizes 4096,16384,212449 -vs baseline -out C:\temp\bench-win\m35-y-ptr-street-real > m35-y-ptr-street-real.log 2>&1
echo m35-y-ptr-street-real done >> m35-progress.txt
m35y-ptr.exe -suite dev -minprocs 6 -maxprocs 8 -memn 262144 -memrounds 3 -keys dirs -values natural -sizes 4096,16384,86215 -vs baseline -out C:\temp\bench-win\m35-y-ptr-dirs-real > m35-y-ptr-dirs-real.log 2>&1
echo m35-y-ptr-dirs-real done >> m35-progress.txt
m35y-ptr.exe -suite dev -minprocs 6 -maxprocs 8 -memn 262144 -memrounds 3 -keys street -values single-value -sizes 4096,16384,212449 -vs baseline -out C:\temp\bench-win\m35-y-ptr-street-single > m35-y-ptr-street-single.log 2>&1
echo m35-y-ptr-street-single done >> m35-progress.txt
m35y-ptr.exe -suite dev -minprocs 6 -maxprocs 8 -memn 262144 -memrounds 3 -keys dirs -values single-value -sizes 4096,16384,86215 -vs baseline -out C:\temp\bench-win\m35-y-ptr-dirs-single > m35-y-ptr-dirs-single.log 2>&1
echo m35-y-ptr-dirs-single done >> m35-progress.txt
echo done > m35-done.txt
