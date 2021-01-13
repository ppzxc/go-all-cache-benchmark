package main

import (
	"crypto/rand"
	"github.com/OrlovEvgeny/go-mcache"
	"github.com/ReneKroon/ttlcache/v2"
	"github.com/akyoto/cache"
	"github.com/allegro/bigcache"
	"github.com/bluele/gcache"
	"github.com/coocood/freecache"
	bg "github.com/dgraph-io/badger/v2"
	"github.com/dgraph-io/ristretto"
	"github.com/muesli/cache2go"
	gc "github.com/patrickmn/go-cache"
	"log"
	"os"
	"strconv"
	"testing"
	"time"
)

var HALF = 512
var RandomBytes = GenerateRandomBytes(HALF)
var LenRandomBytes = int64(len(RandomBytes))

func GenerateRandomBytes(n int) []byte {
	b := make([]byte, n)
	_, _ = rand.Read(b)
	return b
}

func BenchmarkMemoryGCache(b *testing.B) {
	lru := gcache.New(b.N).LRU().Build()

	b.Run("Set_LRU", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			if err := lru.Set(i, RandomBytes); err == nil {
				b.SetBytes(LenRandomBytes)
			}
		}
	})

	b.Run("Get_LRU", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			if value, err := lru.Get(i); err == nil {
				_ = value
				b.SetBytes(LenRandomBytes)
			}
		}
	})
	lru.Purge()

	lfu := gcache.New(b.N).LFU().Build()

	b.Run("Set_LFU", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			if err := lfu.Set(i, RandomBytes); err == nil {
				b.SetBytes(LenRandomBytes)
			}
		}
	})

	b.Run("Get_LFU", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			if value, err := lfu.Get(i); err == nil {
				_ = value
				b.SetBytes(LenRandomBytes)
			}
		}
	})

	lfu.Purge()

	arc := gcache.New(b.N).ARC().Build()

	b.Run("Set_ARC", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			if err := arc.Set(i, RandomBytes); err == nil {
				b.SetBytes(LenRandomBytes)
			}
		}
	})

	b.Run("Get_ARC", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			if value, err := arc.Get(i); err == nil {
				_ = value
				b.SetBytes(LenRandomBytes)
			}
		}
	})

	arc.Purge()
}

func BenchmarkMemoryBigCache(b *testing.B) {
	bc, _ := bigcache.NewBigCache(bigcache.Config{
		// number of shards (must be a power of 2)
		Shards: 1024,
		// time after which entry can be evicted
		LifeWindow: 10 * time.Minute,
		// rps * lifeWindow, used only in initial memory allocation
		MaxEntriesInWindow: 1000 * 10 * 60,
		// max entry size in bytes, used only in initial memory allocation
		MaxEntrySize: 500,
		// cache will not allocate more memory than this limit, value in MB
		// if value is reached then the oldest entries can be overridden for the new ones
		// 0 value means no size limit
		HardMaxCacheSize: 10,
	})
	defer bc.Close()

	b.Run("Set", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			if err := bc.Set(strconv.Itoa(i), RandomBytes); err == nil {
				b.SetBytes(LenRandomBytes)
			}
		}
	})

	b.Run("Get", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			if value, err := bc.Get(strconv.Itoa(i)); err == nil {
				_ = value
				b.SetBytes(LenRandomBytes)
			}
		}
	})
}

func BenchmarkMemoryMCache(b *testing.B) {
	mc := mcache.New()
	defer mc.Close()

	b.Run("Set", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			if err := mc.Set(strconv.Itoa(i), RandomBytes, 5*time.Minute); err == nil {
				b.SetBytes(LenRandomBytes)
			}
		}
	})

	b.Run("Get", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			value, ok := mc.Get(strconv.Itoa(i))
			if ok {
				_ = value
				b.SetBytes(LenRandomBytes)
			}
		}
	})
}

func BenchmarkMemoryTTLCache(b *testing.B) {
	withoutTTL := ttlcache.NewCache()

	b.Run("Set", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			if err := withoutTTL.Set(strconv.Itoa(i), RandomBytes); err == nil {
				b.SetBytes(LenRandomBytes)
			}
		}
	})

	b.Run("Get", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			if value, err := withoutTTL.Get(strconv.Itoa(i)); err == nil {
				b.SetBytes(LenRandomBytes)
				_ = value
			}
		}
	})

	withoutTTL.Close()

	withGlobalTTL := ttlcache.NewCache()

	b.Run("Set_GTTL", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			if err := withGlobalTTL.SetWithTTL(strconv.Itoa(i), RandomBytes, 30*time.Second); err == nil {
				b.SetBytes(LenRandomBytes)
			}
		}
	})

	b.Run("Get_GTTL", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			if value, err := withGlobalTTL.Get(strconv.Itoa(i)); err == nil {
				_ = value
				b.SetBytes(LenRandomBytes)
			}
		}
	})

	withGlobalTTL.Close()

	withTTL := ttlcache.NewCache()

	withTTL.SetTTL(30 * time.Second)
	b.Run("Set_TTL", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			if err := withTTL.Set(strconv.Itoa(i), RandomBytes); err == nil {
				b.SetBytes(LenRandomBytes)
			}
		}
	})

	b.Run("Get_TTL", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			if value, err := withTTL.Get(strconv.Itoa(i)); err == nil {
				_ = value
				b.SetBytes(LenRandomBytes)
			}
		}
	})
}

func benchmarkDiskBadger(b *testing.B) {
	create, _ := os.Getwd()
	diskPath := create + "/tmp/badger"
	db, err := bg.Open(bg.DefaultOptions(diskPath).WithLogger(nil))
	if err != nil {
		log.Fatal(err)
	}

	b.Run("Set", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			db.Update(func(txn *bg.Txn) error {
				if err := txn.Set([]byte(strconv.Itoa(i)), []byte("HI")); err == nil {
					b.SetBytes(LenRandomBytes)
				} else {
					return err
				}
				return nil
			})
		}
	})

	b.Run("Get", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			db.Update(func(txn *bg.Txn) error {
				if value, err := txn.Get([]byte(strconv.Itoa(i))); err == nil {
					_ = value
					b.SetBytes(LenRandomBytes)
				} else {
					return err
				}
				return nil
			})
		}
	})
}

func BenchmarkMemoryBadger(b *testing.B) {
	db, err := bg.Open(bg.DefaultOptions("").WithInMemory(true).WithLogger(nil))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	b.Run("Set", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			db.Update(func(txn *bg.Txn) error {
				if err := txn.Set([]byte(strconv.Itoa(i)), []byte("HI")); err == nil {
					b.SetBytes(LenRandomBytes)
				}
				return err
			})
		}
	})

	b.Run("Get", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			db.Update(func(txn *bg.Txn) error {
				if value, err := txn.Get([]byte(strconv.Itoa(i))); err == nil {
					_ = value
					b.SetBytes(LenRandomBytes)
				} else {
					return err
				}
				return nil
			})
		}
	})
}

func BenchmarkMemoryCache(b *testing.B) {
	c := cache.New(5 * time.Minute)
	defer c.Close()

	b.Run("Set", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			c.Set(strconv.Itoa(i), RandomBytes, 5*time.Minute)
			b.SetBytes(LenRandomBytes)
		}
	})

	b.Run("Get", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			value, found := c.Get(strconv.Itoa(i))
			if found {
				b.SetBytes(LenRandomBytes)
				_ = value
			}
		}
	})
}

func BenchmarkMemoryCache2go(b *testing.B) {
	c2go := cache2go.Cache("myCache")
	defer c2go.Flush()

	b.Run("Set", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			c2go.Add(strconv.Itoa(i), 5*time.Minute, RandomBytes)
			b.SetBytes(LenRandomBytes)
		}
	})

	b.Run("Get", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			value, err := c2go.Value(strconv.Itoa(i))
			if err == nil {
				b.SetBytes(LenRandomBytes)
				_ = value
			}
		}
	})
}

func BenchmarkMemoryGoCache(b *testing.B) {
	goc := gc.New(5*time.Minute, 10*time.Minute)
	defer goc.Flush()

	b.Run("Set", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			goc.Set(strconv.Itoa(i), RandomBytes, 5*time.Minute)
			b.SetBytes(LenRandomBytes)
		}
	})

	b.Run("Get", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			if value, ok := goc.Get(strconv.Itoa(i)); ok {
				b.SetBytes(LenRandomBytes)
				_ = value
			}
		}
	})
}

func BenchmarkMemoryRistretto(b *testing.B) {
	r, _ := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e7,     // Num keys to track frequency of (10M).
		MaxCost:     8 << 30, // Maximum cost of cache (1GB).
		BufferItems: 64,      // Number of keys per Get buffer.
	})
	defer r.Close()
	defer r.Clear()

	b.Run("Set", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			if ok := r.Set(strconv.Itoa(i), RandomBytes, 1); ok {
				b.SetBytes(LenRandomBytes)
			}
		}
	})

	b.Run("Get", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			if value, ok := r.Get(strconv.Itoa(i)); ok {
				b.SetBytes(LenRandomBytes)
				_ = value
			}
		}
	})
}

func BenchmarkMemoryFreeCache(b *testing.B) {
	fc := freecache.NewCache(4000 * 1024 * 1024)
	defer fc.Clear()

	b.Run("Set", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			if err := fc.Set([]byte(strconv.Itoa(i)), RandomBytes, 60); err == nil {
				b.SetBytes(LenRandomBytes)
			}
		}
	})

	b.Run("Get", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			if value, err := fc.Get([]byte(strconv.Itoa(i))); err == nil {
				b.SetBytes(LenRandomBytes)
				_ = value
			}
		}
	})
}
