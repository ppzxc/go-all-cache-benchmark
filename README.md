# golang benchmark all cache lib
## reference
- [Awesome go](https://github.com/avelino/awesome-go)

## memory
- [GCache](https://github.com/bluele/gcache)
- [BigCache](https://github.com/allegro/bigcache)
- [TTLCache](https://github.com/ReneKroon/ttlcache)
- [Cache](https://github.com/akyoto/cache)
- [Cache2go](https://github.com/muesli/cache2go)
- [Go-Cache](https://github.com/patrickmn/go-cache)
- [Ristretto](https://github.com/dgraph-io/ristretto)
- [FreeCache](https://github.com/coocood/freecache)

## disk
- Todo [Badger](https://github.com/dgraph-io/badger)

## test machine hardware info
```
cat /etc/redhat-release 
CentOS Linux release 7.9.2009 (Core)


cat /proc/cpuinfo
model name	: Intel(R) Core(TM) i7-2600 CPU @ 3.40GHz
physical id	: 0
cpu cores	: 4
processor	: 7


cat /proc/meminfo
MemTotal    : 11983936 kB
```

## normal bench, using 1024 bytes data
```
go test -bench=. -cpu=4 -benchmem
```
```
 ⚡ root@localhost  ~/go-all-cache-benchmark   master  go test -bench=. -cpu=4 -benchmem
goos: linux
goarch: amd64
pkg: github.com/ppzxc/go-all-cache-benchmark
BenchmarkMemoryGCache/Set_LRU-4         	 3569758	       325 ns/op	3150.99 MB/s	     151 B/op	       3 allocs/op
BenchmarkMemoryGCache/Get_LRU-4         	14463192	        82.6 ns/op	12390.76 MB/s	       7 B/op	       0 allocs/op
BenchmarkMemoryGCache/Set_LFU-4         	 2909048	       427 ns/op	2400.70 MB/s	     103 B/op	       2 allocs/op
BenchmarkMemoryGCache/Get_LFU-4         	14945622	        80.8 ns/op	12676.45 MB/s	       7 B/op	       0 allocs/op
BenchmarkMemoryGCache/Set_ARC-4         	 1918798	       555 ns/op	1844.37 MB/s	     151 B/op	       3 allocs/op
BenchmarkMemoryGCache/Get_ARC-4         	12585079	       105 ns/op	9768.09 MB/s	       7 B/op	       0 allocs/op
BenchmarkMemoryBigCache/Set-4           	 1980724	       660 ns/op	1550.42 MB/s	      24 B/op	       2 allocs/op
BenchmarkMemoryBigCache/Get-4           	 7500758	       174 ns/op	5873.94 MB/s	      46 B/op	       1 allocs/op
BenchmarkMemoryMCache/Set-4             	  290188	      3807 ns/op	 268.99 MB/s	    1129 B/op	       9 allocs/op
BenchmarkMemoryMCache/Get-4             	 1397383	       855 ns/op	1198.14 MB/s	     135 B/op	       2 allocs/op
BenchmarkMemoryTTLCache/Set-4           	  948326	      1547 ns/op	 661.74 MB/s	     295 B/op	       3 allocs/op
BenchmarkMemoryTTLCache/Get-4           	 1523881	       684 ns/op	1496.64 MB/s	       7 B/op	       0 allocs/op
BenchmarkMemoryTTLCache/Set_GTTL-4      	  873480	      1577 ns/op	 649.17 MB/s	     300 B/op	       3 allocs/op
BenchmarkMemoryTTLCache/Get_GTTL-4      	 1345822	       860 ns/op	1190.03 MB/s	       7 B/op	       0 allocs/op
BenchmarkMemoryTTLCache/Set_TTL-4       	  954970	      1586 ns/op	 645.56 MB/s	     294 B/op	       3 allocs/op
BenchmarkMemoryTTLCache/Get_TTL-4       	 1367469	       889 ns/op	1152.50 MB/s	       7 B/op	       0 allocs/op
BenchmarkMemoryBadger/Set-4             	  109624	     11191 ns/op	  91.50 MB/s	    1127 B/op	      29 allocs/op
BenchmarkMemoryBadger/Get-4             	  712833	      1683 ns/op	 608.30 MB/s	     540 B/op	      12 allocs/op
BenchmarkMemoryCache/Set-4              	 2308275	       742 ns/op	1380.81 MB/s	     168 B/op	       5 allocs/op
BenchmarkMemoryCache/Get-4              	 4503571	       237 ns/op	4326.22 MB/s	       7 B/op	       0 allocs/op
BenchmarkMemoryCache2go/Set-4           	 2093635	       737 ns/op	1389.98 MB/s	     283 B/op	       6 allocs/op
BenchmarkMemoryCache2go/Get-4           	 3320185	       350 ns/op	2926.86 MB/s	      23 B/op	       1 allocs/op
BenchmarkMemoryGoCache/Set-4            	 3894265	       464 ns/op	2208.31 MB/s	     141 B/op	       2 allocs/op
BenchmarkMemoryGoCache/Get-4            	 5531332	       219 ns/op	4668.39 MB/s	       7 B/op	       0 allocs/op
BenchmarkMemoryRistretto/Set-4          	 2519400	       706 ns/op	1450.13 MB/s	     393 B/op	       4 allocs/op
BenchmarkMemoryRistretto/Get-4          	 3466861	       346 ns/op	2963.54 MB/s	      31 B/op	       2 allocs/op
BenchmarkMemoryFreeCache/Set-4          	 2582163	       550 ns/op	1860.30 MB/s	      47 B/op	       1 allocs/op
BenchmarkMemoryFreeCache/Get-4          	 1645215	       660 ns/op	1550.44 MB/s	    1031 B/op	       1 allocs/op
PASS
ok  	github.com/ppzxc/go-all-cache-benchmark	86.170s
```

## fix count bench, using 1024 bytes data
```
go test -bench=. -cpu=4  -benchmem -benchtime 1000000x
```
```
 ⚡ root@localhost  ~/go-all-cache-benchmark   master  go test -bench=. -cpu=4  -benchmem -benchtime 1000000x
goos: linux
goarch: amd64
pkg: github.com/ppzxc/go-all-cache-benchmark
BenchmarkMemoryGCache/Set_LRU-4         	 1000000	       346 ns/op	2963.61 MB/s	     151 B/op	       3 allocs/op
BenchmarkMemoryGCache/Get_LRU-4         	 1000000	        82.4 ns/op	12427.33 MB/s	       7 B/op	       0 allocs/op
BenchmarkMemoryGCache/Set_LFU-4         	 1000000	       379 ns/op	2699.59 MB/s	     104 B/op	       2 allocs/op
BenchmarkMemoryGCache/Get_LFU-4         	 1000000	        93.8 ns/op	10913.52 MB/s	       7 B/op	       0 allocs/op
BenchmarkMemoryGCache/Set_ARC-4         	 1000000	       529 ns/op	1937.45 MB/s	     151 B/op	       3 allocs/op
BenchmarkMemoryGCache/Get_ARC-4         	 1000000	        94.7 ns/op	10809.33 MB/s	       7 B/op	       0 allocs/op
BenchmarkMemoryBigCache/Set-4           	 1000000	       608 ns/op	1684.67 MB/s	      23 B/op	       1 allocs/op
BenchmarkMemoryBigCache/Get-4           	 1000000	       274 ns/op	3739.49 MB/s	     300 B/op	       1 allocs/op
BenchmarkMemoryMCache/Set-4             	 1000000	      3941 ns/op	 259.82 MB/s	    1158 B/op	       9 allocs/op
BenchmarkMemoryMCache/Get-4             	 1000000	       899 ns/op	1139.43 MB/s	     135 B/op	       2 allocs/op
BenchmarkMemoryTTLCache/Set-4           	 1000000	      1432 ns/op	 715.20 MB/s	     288 B/op	       3 allocs/op
BenchmarkMemoryTTLCache/Get-4           	 1000000	      1012 ns/op	1012.14 MB/s	       7 B/op	       0 allocs/op
BenchmarkMemoryTTLCache/Set_GTTL-4      	 1000000	      1525 ns/op	 671.28 MB/s	     288 B/op	       3 allocs/op
BenchmarkMemoryTTLCache/Get_GTTL-4      	 1000000	      1222 ns/op	 837.75 MB/s	       7 B/op	       0 allocs/op
BenchmarkMemoryTTLCache/Set_TTL-4       	 1000000	      1449 ns/op	 706.66 MB/s	     288 B/op	       3 allocs/op
BenchmarkMemoryTTLCache/Get_TTL-4       	 1000000	      1222 ns/op	 837.74 MB/s	       7 B/op	       0 allocs/op
BenchmarkMemoryBadger/Set-4             	 1000000	     11483 ns/op	  89.18 MB/s	    1127 B/op	      29 allocs/op
BenchmarkMemoryBadger/Get-4             	 1000000	      1658 ns/op	 617.70 MB/s	     519 B/op	      12 allocs/op
BenchmarkMemoryCache/Set-4              	 1000000	       938 ns/op	1092.15 MB/s	     251 B/op	       7 allocs/op
BenchmarkMemoryCache/Get-4              	 1000000	       335 ns/op	3058.36 MB/s	       7 B/op	       0 allocs/op
BenchmarkMemoryCache2go/Set-4           	 1000000	       854 ns/op	1199.11 MB/s	     347 B/op	       6 allocs/op
BenchmarkMemoryCache2go/Get-4           	 1000000	       356 ns/op	2877.47 MB/s	      23 B/op	       1 allocs/op
BenchmarkMemoryGoCache/Set-4            	 1000000	       488 ns/op	2099.24 MB/s	     240 B/op	       2 allocs/op
BenchmarkMemoryGoCache/Get-4            	 1000000	       222 ns/op	4602.69 MB/s	       7 B/op	       0 allocs/op
BenchmarkMemoryRistretto/Set-4          	 1000000	       522 ns/op	1962.86 MB/s	     339 B/op	       4 allocs/op
BenchmarkMemoryRistretto/Get-4          	 1000000	       308 ns/op	3329.98 MB/s	      31 B/op	       2 allocs/op
BenchmarkMemoryFreeCache/Set-4          	 1000000	       574 ns/op	1784.71 MB/s	      72 B/op	       1 allocs/op
BenchmarkMemoryFreeCache/Get-4          	 1000000	       648 ns/op	1579.50 MB/s	    1031 B/op	       1 allocs/op
PASS
ok  	github.com/ppzxc/go-all-cache-benchmark	54.495s
```

## fix time bench, using 1024 bytes data
```
go test -bench=. -cpu=4  -benchmem -benchtime 5s
```
```
 ✘ ⚡ root@localhost  ~/go-all-cache-benchmark   master  go test -bench=. -cpu=4  -benchmem -benchtime 5s 
goos: linux
goarch: amd64
pkg: github.com/ppzxc/go-all-cache-benchmark
BenchmarkMemoryGCache/Set_LRU-4         	16504808	       331 ns/op	3090.35 MB/s	     152 B/op	       3 allocs/op
BenchmarkMemoryGCache/Get_LRU-4         	73070196	        85.2 ns/op	12018.03 MB/s	       7 B/op	       0 allocs/op
BenchmarkMemoryGCache/Set_LFU-4         	12171514	       414 ns/op	2474.63 MB/s	     104 B/op	       2 allocs/op
BenchmarkMemoryGCache/Get_LFU-4         	65472864	        83.5 ns/op	12257.32 MB/s	       7 B/op	       0 allocs/op
BenchmarkMemoryGCache/Set_ARC-4         	 8904129	       591 ns/op	1734.03 MB/s	     152 B/op	       3 allocs/op
BenchmarkMemoryGCache/Get_ARC-4         	63415984	        99.9 ns/op	10252.52 MB/s	       7 B/op	       0 allocs/op
BenchmarkMemoryBigCache/Set-4           	10006041	       663 ns/op	1543.65 MB/s	      24 B/op	       2 allocs/op
BenchmarkMemoryBigCache/Get-4           	36587178	       168 ns/op	6099.65 MB/s	      15 B/op	       1 allocs/op
BenchmarkMemoryMCache/Set-4             	 1448583	      3488 ns/op	 293.58 MB/s	     998 B/op	       9 allocs/op
BenchmarkMemoryMCache/Get-4             	 6313848	       848 ns/op	1206.94 MB/s	     135 B/op	       2 allocs/op
BenchmarkMemoryTTLCache/Set-4           	 4187820	      1417 ns/op	 722.64 MB/s	     220 B/op	       2 allocs/op
BenchmarkMemoryTTLCache/Get-4           	 9617332	       547 ns/op	1873.42 MB/s	       7 B/op	       0 allocs/op
BenchmarkMemoryTTLCache/Set_GTTL-4      	 3957517	      1750 ns/op	 585.12 MB/s	     225 B/op	       2 allocs/op
BenchmarkMemoryTTLCache/Get_GTTL-4      	 5981611	       923 ns/op	1109.16 MB/s	       7 B/op	       0 allocs/op
BenchmarkMemoryTTLCache/Set_TTL-4       	 4108146	      1709 ns/op	 599.20 MB/s	     222 B/op	       2 allocs/op
BenchmarkMemoryTTLCache/Get_TTL-4       	 5763652	       991 ns/op	1033.33 MB/s	       7 B/op	       0 allocs/op
BenchmarkMemoryBadger/Set-4             	  553262	     11280 ns/op	  90.78 MB/s	    1127 B/op	      29 allocs/op
BenchmarkMemoryBadger/Get-4             	 3656499	      1700 ns/op	 602.52 MB/s	     540 B/op	      12 allocs/op
BenchmarkMemoryCache/Set-4              	10212211	       762 ns/op	1343.58 MB/s	     158 B/op	       5 allocs/op
BenchmarkMemoryCache/Get-4              	22535676	       236 ns/op	4345.72 MB/s	       7 B/op	       0 allocs/op
BenchmarkMemoryCache2go/Set-4           	 6940514	       952 ns/op	1075.76 MB/s	     348 B/op	       6 allocs/op
BenchmarkMemoryCache2go/Get-4           	17639505	       324 ns/op	3162.79 MB/s	      23 B/op	       1 allocs/op
BenchmarkMemoryGoCache/Set-4            	11678008	       496 ns/op	2065.33 MB/s	     163 B/op	       2 allocs/op
BenchmarkMemoryGoCache/Get-4            	26572488	       226 ns/op	4540.56 MB/s	       7 B/op	       0 allocs/op
BenchmarkMemoryRistretto/Set-4          	12783396	       661 ns/op	1548.59 MB/s	     376 B/op	       4 allocs/op
BenchmarkMemoryRistretto/Get-4          	16195382	       386 ns/op	2652.43 MB/s	      31 B/op	       2 allocs/op
BenchmarkMemoryFreeCache/Set-4          	10401394	       986 ns/op	1039.05 MB/s	      27 B/op	       1 allocs/op
BenchmarkMemoryFreeCache/Get-4          	20908278	       376 ns/op	2725.81 MB/s	     202 B/op	       1 allocs/op
PASS
ok  	github.com/ppzxc/go-all-cache-benchmark	400.426s
```

## How to use
1. git clone https://github.com/ppzxc/go-all-cache-benchmark
2. go test -bench=.
3. go test -bench=. -cpu=4 -benchmem
4. go test -bench=. -cpu=4 -benchmem -benchtime 1000000x
5. go test -bench=. -cpu=4 -benchmem -benchtime 5s


## Todo
- Disk Base KV, Cache Lib Benchmark