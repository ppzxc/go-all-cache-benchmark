# golang benchmark all cache lib
## memory
- [GCache](https://github.com/bluele/gcache)
- [BigCache](https://github.com/allegro/bigcache)
- [TTLCache](https://github.com/ReneKroon/ttlcache)
- [Cache](https://github.com/akyoto/cache)
- [Cache2go](https://github.com/muesli/cache2go)
- [Go-Cache](https://github.com/patrickmn/go-cache)

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

## normal bench, using 32,768 bytes data
```
go test -bench=. -cpu=4 -benchmem
```
```
[root@localhost go-all-cache-benchmark]# go test -bench=. -cpu=4 -benchmem
goos: linux
goarch: amd64
pkg: github.com/ppzxc/go-all-cache-benchmark
BenchmarkMemoryGCache/Set_LRU-4         	 3168678	       353 ns/op	92738.68 MB/s	     151 B/op	       3 allocs/op
BenchmarkMemoryGCache/Get_LRU-4         	13445673	        85.4 ns/op	383621.50 MB/s	       7 B/op	       0 allocs/op
BenchmarkMemoryGCache/Set_LFU-4         	 3018172	       413 ns/op	79394.20 MB/s	     103 B/op	       2 allocs/op
BenchmarkMemoryGCache/Get_LFU-4         	14789017	        91.3 ns/op	358906.98 MB/s	       7 B/op	       0 allocs/op
BenchmarkMemoryGCache/Set_ARC-4         	 2065030	       654 ns/op	50133.42 MB/s	     151 B/op	       3 allocs/op
BenchmarkMemoryGCache/Get_ARC-4         	12596788	        95.1 ns/op	344491.08 MB/s	       7 B/op	       0 allocs/op
BenchmarkMemoryBigCache/Set-4           	  137250	      7657 ns/op	4279.36 MB/s	     633 B/op	       2 allocs/op
BenchmarkMemoryBigCache/Get-4           	 7181403	       167 ns/op	195786.87 MB/s	      41 B/op	       1 allocs/op
BenchmarkMemoryMCache/Set-4             	  324243	      3882 ns/op	8440.76 MB/s	    1113 B/op	       9 allocs/op
BenchmarkMemoryMCache/Get-4             	 1365093	       812 ns/op	40349.48 MB/s	     135 B/op	       2 allocs/op
BenchmarkMemoryTTLCache/Set-4           	  947601	      1530 ns/op	21413.70 MB/s	     295 B/op	       3 allocs/op
BenchmarkMemoryTTLCache/Get-4           	 1522677	       684 ns/op	47940.26 MB/s	       7 B/op	       0 allocs/op
BenchmarkMemoryTTLCache/Set_GTTL-4      	  849854	      1505 ns/op	21768.35 MB/s	     237 B/op	       3 allocs/op
BenchmarkMemoryTTLCache/Get_GTTL-4      	 1398577	       789 ns/op	41554.85 MB/s	       7 B/op	       0 allocs/op
BenchmarkMemoryTTLCache/Set_TTL-4       	  961063	      1561 ns/op	20987.92 MB/s	     293 B/op	       3 allocs/op
BenchmarkMemoryTTLCache/Get_TTL-4       	 1370934	       887 ns/op	36958.46 MB/s	       7 B/op	       0 allocs/op
BenchmarkDiskBadger/Set-4               	    5086	    232799 ns/op	 140.76 MB/s	    1295 B/op	      35 allocs/op
BenchmarkDiskBadger/Get-4               	  818130	      1689 ns/op	19395.38 MB/s	     543 B/op	      12 allocs/op
BenchmarkMemoryBadger/Set-4             	  110594	     11401 ns/op	2874.10 MB/s	    1127 B/op	      29 allocs/op
BenchmarkMemoryBadger/Get-4             	  761905	      1691 ns/op	19379.81 MB/s	     540 B/op	      12 allocs/op
BenchmarkMemoryCache/Set-4              	 2212816	       755 ns/op	43394.21 MB/s	     169 B/op	       5 allocs/op
BenchmarkMemoryCache/Get-4              	 4605349	       228 ns/op	143609.56 MB/s	       7 B/op	       0 allocs/op
BenchmarkMemoryCache2go/Set-4           	 2261769	       717 ns/op	45671.89 MB/s	     278 B/op	       6 allocs/op
BenchmarkMemoryCache2go/Get-4           	 3305065	       323 ns/op	101323.90 MB/s	      23 B/op	       1 allocs/op
BenchmarkMemoryGoCache/Set-4            	 3800065	       456 ns/op	71865.78 MB/s	     143 B/op	       2 allocs/op
BenchmarkMemoryGoCache/Get-4            	 5515010	       218 ns/op	150650.07 MB/s	       7 B/op	       0 allocs/op
PASS
ok  	github.com/ppzxc/go-all-cache-benchmark	78.609s
```

## fix count bench, using 32,768 bytes data
```
go test -bench=. -cpu=4  -benchmem -benchtime 1000000x
```
```
[root@localhost go-all-cache-benchmark]# go test -bench=. -cpu=4  -benchmem -benchtime 1000000x
goos: linux
goarch: amd64
pkg: github.com/ppzxc/go-all-cache-benchmark
BenchmarkMemoryGCache/Set_LRU-4         	 1000000	       394 ns/op	83070.85 MB/s	     151 B/op	       3 allocs/op
BenchmarkMemoryGCache/Get_LRU-4         	 1000000	        84.9 ns/op	385984.49 MB/s	       7 B/op	       0 allocs/op
BenchmarkMemoryGCache/Set_LFU-4         	 1000000	       454 ns/op	72237.06 MB/s	     103 B/op	       2 allocs/op
BenchmarkMemoryGCache/Get_LFU-4         	 1000000	        80.0 ns/op	409703.96 MB/s	       7 B/op	       0 allocs/op
BenchmarkMemoryGCache/Set_ARC-4         	 1000000	       636 ns/op	51526.68 MB/s	     152 B/op	       3 allocs/op
BenchmarkMemoryGCache/Get_ARC-4         	 1000000	        94.5 ns/op	346801.68 MB/s	       7 B/op	       0 allocs/op
BenchmarkMemoryBigCache/Set-4           	 1000000	      7639 ns/op	4289.38 MB/s	     180 B/op	       1 allocs/op
BenchmarkMemoryBigCache/Get-4           	 1000000	       248 ns/op	131983.70 MB/s	     246 B/op	       1 allocs/op
BenchmarkMemoryMCache/Set-4             	 1000000	      3934 ns/op	8328.58 MB/s	    1158 B/op	       9 allocs/op
BenchmarkMemoryMCache/Get-4             	 1000000	       896 ns/op	36559.51 MB/s	     135 B/op	       2 allocs/op
BenchmarkMemoryTTLCache/Set-4           	 1000000	      1399 ns/op	23423.92 MB/s	     288 B/op	       3 allocs/op
BenchmarkMemoryTTLCache/Get-4           	 1000000	      1018 ns/op	32185.94 MB/s	       7 B/op	       0 allocs/op
BenchmarkMemoryTTLCache/Set_GTTL-4      	 1000000	      1492 ns/op	21959.03 MB/s	     288 B/op	       3 allocs/op
BenchmarkMemoryTTLCache/Get_GTTL-4      	 1000000	      1232 ns/op	26600.88 MB/s	       7 B/op	       0 allocs/op
BenchmarkMemoryTTLCache/Set_TTL-4       	 1000000	      1445 ns/op	22683.62 MB/s	     288 B/op	       3 allocs/op
BenchmarkMemoryTTLCache/Get_TTL-4       	 1000000	      1226 ns/op	26717.17 MB/s	       7 B/op	       0 allocs/op
BenchmarkMemoryCache/Set-4              	 1000000	       934 ns/op	35080.38 MB/s	     251 B/op	       7 allocs/op
BenchmarkMemoryCache/Get-4              	 1000000	       332 ns/op	98763.58 MB/s	       7 B/op	       0 allocs/op
BenchmarkMemoryCache2go/Set-4           	 1000000	       834 ns/op	39272.71 MB/s	     347 B/op	       6 allocs/op
BenchmarkMemoryCache2go/Get-4           	 1000000	       346 ns/op	94623.88 MB/s	      23 B/op	       1 allocs/op
BenchmarkMemoryGoCache/Set-4            	 1000000	       505 ns/op	64824.72 MB/s	     240 B/op	       2 allocs/op
BenchmarkMemoryGoCache/Get-4            	 1000000	       224 ns/op	145993.42 MB/s	       7 B/op	       0 allocs/op
PASS
ok  	github.com/ppzxc/go-all-cache-benchmark	38.561s
```

## fix time bench, using 32,768 bytes data
```
go test -bench=. -cpu=4  -benchmem -benchtime 10s
```
```
[root@localhost go-all-cache-benchmark]# go test -bench=. -cpu=4  -benchmem -benchtime 5s
goos: linux
goarch: amd64
pkg: github.com/ppzxc/go-all-cache-benchmark
BenchmarkMemoryGCache/Set_LRU-4         	14555043	       359 ns/op	91173.89 MB/s	     152 B/op	       3 allocs/op
BenchmarkMemoryGCache/Get_LRU-4         	70477158	        82.8 ns/op	395733.21 MB/s	       7 B/op	       0 allocs/op
BenchmarkMemoryGCache/Set_LFU-4         	16452160	       418 ns/op	78483.56 MB/s	     104 B/op	       2 allocs/op
BenchmarkMemoryGCache/Get_LFU-4         	69063326	        84.0 ns/op	389960.03 MB/s	       7 B/op	       0 allocs/op
BenchmarkMemoryGCache/Set_ARC-4         	10514586	       649 ns/op	50509.77 MB/s	     152 B/op	       3 allocs/op
BenchmarkMemoryGCache/Get_ARC-4         	59669031	       104 ns/op	315883.75 MB/s	       7 B/op	       0 allocs/op
BenchmarkMemoryBigCache/Set-4           	  674002	      7626 ns/op	4296.63 MB/s	     148 B/op	       2 allocs/op
BenchmarkMemoryBigCache/Get-4           	37730102	       161 ns/op	203358.78 MB/s	      14 B/op	       1 allocs/op
BenchmarkMemoryMCache/Set-4             	 1465910	      3638 ns/op	9006.91 MB/s	     997 B/op	       9 allocs/op
BenchmarkMemoryMCache/Get-4             	 6596492	       843 ns/op	38858.36 MB/s	     135 B/op	       2 allocs/op
BenchmarkMemoryTTLCache/Set-4           	 4298888	      1382 ns/op	23712.53 MB/s	     217 B/op	       2 allocs/op
BenchmarkMemoryTTLCache/Get-4           	 9467466	       558 ns/op	58735.84 MB/s	       7 B/op	       0 allocs/op
BenchmarkMemoryTTLCache/Set_GTTL-4      	 3984811	      1726 ns/op	18989.38 MB/s	     225 B/op	       2 allocs/op
BenchmarkMemoryTTLCache/Get_GTTL-4      	 5898128	       935 ns/op	35030.85 MB/s	       7 B/op	       0 allocs/op
BenchmarkMemoryTTLCache/Set_TTL-4       	 4150930	      1679 ns/op	19520.15 MB/s	     221 B/op	       2 allocs/op
BenchmarkMemoryTTLCache/Get_TTL-4       	 5611174	      1020 ns/op	32130.82 MB/s	       7 B/op	       0 allocs/op
BenchmarkMemoryCache/Set-4              	 6203104	       821 ns/op	39902.66 MB/s	     186 B/op	       6 allocs/op
BenchmarkMemoryCache/Get-4              	26141553	       224 ns/op	146519.80 MB/s	       7 B/op	       0 allocs/op
BenchmarkMemoryCache2go/Set-4           	 7041307	      1009 ns/op	32466.63 MB/s	     347 B/op	       6 allocs/op
BenchmarkMemoryCache2go/Get-4           	15188438	       336 ns/op	97631.89 MB/s	      23 B/op	       1 allocs/op
BenchmarkMemoryGoCache/Set-4            	11838270	       481 ns/op	68164.77 MB/s	     162 B/op	       2 allocs/op
BenchmarkMemoryGoCache/Get-4            	28197944	       208 ns/op	157537.42 MB/s	       7 B/op	       0 allocs/op
PASS
ok  	github.com/ppzxc/go-all-cache-benchmark	300.788s
```

## How to use this benchmark
- git clone https://github.com/ppzxc/go-all-cache-benchmark
- go test -bench=.

## Todo
- Disk Base KV, Cache Lib Benchmark