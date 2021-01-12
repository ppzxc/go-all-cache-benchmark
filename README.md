# go all cache lib benchmark

- [GCache](https://github.com/bluele/gcache)
- [BigCache](https://github.com/allegro/bigcache)
- [TTLCache](https://github.com/ReneKroon/ttlcache)
- [Badger](https://github.com/dgraph-io/badger)
- [Cache](https://github.com/akyoto/cache)
- [Cache2go](https://github.com/muesli/cache2go)
- [Go-Cache](https://github.com/patrickmn/go-cache)

## normal bench
```
go test -bench=. -cpu=4 -benchmem
```
```

```

## fix execute bench count
```
go test -bench=. -cpu=4  -benchmem -benchtime 10000000x
```

## fix execute bench time
```
go test -bench=. -cpu=4  -benchmem -benchtime 10s
```