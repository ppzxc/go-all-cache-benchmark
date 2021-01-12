package main

import (
	"testing"
	"time"
	"github.com/bluele/gcache"
)

func BenchmarkGCacheMemoryGetSet(b *testing.B) {
	gc := gcache.New(b.N).
		ARC().
		Build()
	defer gc.Purge()
	for i := 0; i < b.N; i++ {
		err := gc.SetWithExpire(i, []byte("HI"), 10 * time.Second)
		if err != nil {
			b.Error(err)
		}
	}

	for i := 0; i < b.N; i++ {
		_, err := gc.Get(i)
		if err != nil {
			b.Error(err)
		}
	}
}
