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
- [CCache](https://github.com/karlseguin/ccache)

## disk
- Todo [Badger](https://github.com/dgraph-io/badger)

## test machine hardware info
```
cat /etc/redhat-release 
CentOS Linux release 7.9.2009 (Core)


cat /proc/cpuinfo
model name	: Intel(R) Core(TM) i5-2500 CPU @ 3.30GHz
cpu cores	: 4
processor	: 3


cat /proc/meminfo
MemTotal:        8080068 kB
```

## normal bench, using 1024 bytes data
```
go test -bench=. -cpu=4 -benchmem
```
```
  ⚡ root@localhost  ~/test  go test -bench=. -cpu=4 -benchmem
goos: linux
goarch: amd64
pkg: test
cpu: Intel(R) Core(TM) i5-2500 CPU @ 3.30GHz
BenchmarkMemoryGCache/Set_LRU-4         	 3576657	       345.3 ns/op	2965.75 MB/s	     143 B/op	       3 allocs/op
BenchmarkMemoryGCache/Get_LRU-4         	14452994	        84.61 ns/op	12102.31 MB/s	       7 B/op	       0 allocs/op
BenchmarkMemoryGCache/Set_LFU-4         	 3311752	       380.0 ns/op	2694.98 MB/s	      96 B/op	       2 allocs/op
BenchmarkMemoryGCache/Get_LFU-4         	13973755	        84.89 ns/op	12062.67 MB/s	       7 B/op	       0 allocs/op
BenchmarkMemoryGCache/Set_ARC-4         	 2102664	       562.3 ns/op	1821.14 MB/s	     143 B/op	       3 allocs/op
BenchmarkMemoryGCache/Get_ARC-4         	12152180	       104.1 ns/op	9834.66 MB/s	       7 B/op	       0 allocs/op
BenchmarkMemoryBigCache/Set-4           	 1947426	       690.9 ns/op	1482.10 MB/s	      24 B/op	       2 allocs/op
BenchmarkMemoryBigCache/Get-4           	 7236795	       181.1 ns/op	5655.36 MB/s	      48 B/op	       1 allocs/op
BenchmarkMemoryMCache/Set-4             	  307893	      3641 ns/op	 281.21 MB/s	    1105 B/op	       9 allocs/op
BenchmarkMemoryMCache/Get-4             	 1397437	       832.1 ns/op	1230.69 MB/s	     127 B/op	       2 allocs/op
BenchmarkMemoryCCache/Set-4             	 2191299	       569.5 ns/op	1797.97 MB/s	     162 B/op	       3 allocs/op
BenchmarkMemoryCCache/Get-4             	 9713568	       143.9 ns/op	7116.77 MB/s	       8 B/op	       1 allocs/op
BenchmarkMemoryTTLCache/Set-4           	  976320	      1474 ns/op	 694.67 MB/s	     282 B/op	       3 allocs/op
BenchmarkMemoryTTLCache/Get-4           	 1211499	       844.6 ns/op	1212.48 MB/s	       7 B/op	       0 allocs/op
BenchmarkMemoryTTLCache/Set_GTTL-4      	  900860	      1596 ns/op	 641.52 MB/s	     286 B/op	       3 allocs/op
BenchmarkMemoryTTLCache/Get_GTTL-4      	 1260369	       962.3 ns/op	1064.12 MB/s	       7 B/op	       0 allocs/op
BenchmarkMemoryTTLCache/Set_TTL-4       	  934764	      1571 ns/op	 651.93 MB/s	     290 B/op	       3 allocs/op
BenchmarkMemoryTTLCache/Get_TTL-4       	 1289794	       941.7 ns/op	1087.45 MB/s	       7 B/op	       0 allocs/op
BenchmarkMemoryBadger/Set-4             	  110365	     11282 ns/op	  90.77 MB/s	    1118 B/op	      29 allocs/op
BenchmarkMemoryBadger/Get-4             	  753824	      1618 ns/op	 632.99 MB/s	     540 B/op	      12 allocs/op
BenchmarkMemoryCache/Set-4              	 2272722	       741.7 ns/op	1380.54 MB/s	     153 B/op	       5 allocs/op
BenchmarkMemoryCache/Get-4              	 4380367	       241.9 ns/op	4233.85 MB/s	       7 B/op	       0 allocs/op
BenchmarkMemoryCache2go/Set-4           	 2303300	       716.9 ns/op	1428.30 MB/s	     269 B/op	       6 allocs/op
BenchmarkMemoryCache2go/Get-4           	 3308176	       330.3 ns/op	3100.48 MB/s	      23 B/op	       1 allocs/op
BenchmarkMemoryGoCache/Set-4            	 3793915	       458.0 ns/op	2235.61 MB/s	     136 B/op	       2 allocs/op
BenchmarkMemoryGoCache/Get-4            	 5524476	       219.0 ns/op	4676.21 MB/s	       7 B/op	       0 allocs/op
BenchmarkMemoryRistretto/Set-4          	 2532092	       673.3 ns/op	1520.95 MB/s	     384 B/op	       4 allocs/op
BenchmarkMemoryRistretto/Get-4          	 3515143	       345.9 ns/op	2960.70 MB/s	      31 B/op	       2 allocs/op
BenchmarkMemoryFreeCache/Set-4          	 2530926	       555.5 ns/op	1843.34 MB/s	      47 B/op	       1 allocs/op
BenchmarkMemoryFreeCache/Get-4          	 1600970	       681.9 ns/op	1501.70 MB/s	    1031 B/op	       1 allocs/op
PASS
ok  	test	86.660s
```

## fix count bench, using 1024 bytes data
```
go test -bench=. -cpu=4  -benchmem -benchtime 1000000x
```
```
⚡ root@localhost  ~/test  go test -bench=. -cpu=4  -benchmem -benchtime 1000000x
goos: linux
goarch: amd64
pkg: test
cpu: Intel(R) Core(TM) i5-2500 CPU @ 3.30GHz
BenchmarkMemoryGCache/Set_LRU-4         	 1000000	       336.7 ns/op	3041.31 MB/s	     143 B/op	       3 allocs/op
BenchmarkMemoryGCache/Get_LRU-4         	 1000000	        85.82 ns/op	11931.73 MB/s	       7 B/op	       0 allocs/op
BenchmarkMemoryGCache/Set_LFU-4         	 1000000	       403.3 ns/op	2538.91 MB/s	      95 B/op	       2 allocs/op
BenchmarkMemoryGCache/Get_LFU-4         	 1000000	        85.51 ns/op	11974.63 MB/s	       7 B/op	       0 allocs/op
BenchmarkMemoryGCache/Set_ARC-4         	 1000000	       619.4 ns/op	1653.19 MB/s	     144 B/op	       3 allocs/op
BenchmarkMemoryGCache/Get_ARC-4         	 1000000	        98.26 ns/op	10421.31 MB/s	       7 B/op	       0 allocs/op
BenchmarkMemoryBigCache/Set-4           	 1000000	       625.3 ns/op	1637.64 MB/s	      23 B/op	       1 allocs/op
BenchmarkMemoryBigCache/Get-4           	 1000000	       290.3 ns/op	3527.09 MB/s	     300 B/op	       1 allocs/op
BenchmarkMemoryMCache/Set-4             	 1000000	      3760 ns/op	 272.33 MB/s	    1142 B/op	       9 allocs/op
BenchmarkMemoryMCache/Get-4             	 1000000	       927.0 ns/op	1104.59 MB/s	     127 B/op	       2 allocs/op
BenchmarkMemoryCCache/Set-4             	 1000000	       551.2 ns/op	1857.68 MB/s	     163 B/op	       4 allocs/op
BenchmarkMemoryCCache/Get-4             	 1000000	       130.4 ns/op	7849.97 MB/s	       7 B/op	       0 allocs/op
BenchmarkMemoryTTLCache/Set-4           	 1000000	      1411 ns/op	 725.69 MB/s	     280 B/op	       3 allocs/op
BenchmarkMemoryTTLCache/Get-4           	 1000000	      1023 ns/op	1001.41 MB/s	       7 B/op	       0 allocs/op
BenchmarkMemoryTTLCache/Set_GTTL-4      	 1000000	      1534 ns/op	 667.61 MB/s	     280 B/op	       3 allocs/op
BenchmarkMemoryTTLCache/Get_GTTL-4      	 1000000	      1259 ns/op	 813.58 MB/s	       7 B/op	       0 allocs/op
BenchmarkMemoryTTLCache/Set_TTL-4       	 1000000	      1481 ns/op	 691.58 MB/s	     280 B/op	       3 allocs/op
BenchmarkMemoryTTLCache/Get_TTL-4       	 1000000	      1263 ns/op	 810.92 MB/s	       7 B/op	       0 allocs/op
BenchmarkMemoryBadger/Set-4             	 1000000	     11670 ns/op	  87.75 MB/s	    1118 B/op	      29 allocs/op
BenchmarkMemoryBadger/Get-4             	 1000000	      1501 ns/op	 681.99 MB/s	     519 B/op	      12 allocs/op
BenchmarkMemoryCache/Set-4              	 1000000	       935.5 ns/op	1094.59 MB/s	     235 B/op	       7 allocs/op
BenchmarkMemoryCache/Get-4              	 1000000	       332.3 ns/op	3081.93 MB/s	       7 B/op	       0 allocs/op
BenchmarkMemoryCache2go/Set-4           	 1000000	       855.9 ns/op	1196.36 MB/s	     339 B/op	       6 allocs/op
BenchmarkMemoryCache2go/Get-4           	 1000000	       368.4 ns/op	2779.27 MB/s	      23 B/op	       1 allocs/op
BenchmarkMemoryGoCache/Set-4            	 1000000	       515.5 ns/op	1986.44 MB/s	     232 B/op	       2 allocs/op
BenchmarkMemoryGoCache/Get-4            	 1000000	       240.1 ns/op	4264.36 MB/s	       7 B/op	       0 allocs/op
BenchmarkMemoryRistretto/Set-4          	 1000000	       464.4 ns/op	2205.21 MB/s	     331 B/op	       4 allocs/op
BenchmarkMemoryRistretto/Get-4          	 1000000	       316.7 ns/op	3233.02 MB/s	      31 B/op	       2 allocs/op
BenchmarkMemoryFreeCache/Set-4          	 1000000	       598.9 ns/op	1709.88 MB/s	      72 B/op	       1 allocs/op
BenchmarkMemoryFreeCache/Get-4          	 1000000	       662.3 ns/op	1546.13 MB/s	    1031 B/op	       1 allocs/op
PASS
ok  	test	55.522s
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