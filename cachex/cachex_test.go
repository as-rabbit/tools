package concx

import (
	"fmt"
	"github.com/coocood/freecache"
	"testing"
	"tool/cachex/bigcache"
)

const maxEntrySize = 256

func key(i int) string {
	return fmt.Sprintf("key-%010d", i)
}

func value() []byte {
	return make([]byte, 100)
}

func BenchmarkFreeCacheSet(b *testing.B) {
	cache := freecache.NewCache(b.N * maxEntrySize)
	for i := 0; i < b.N; i++ {
		cache.Set([]byte(key(i)), value(), 0)
	}

}

func BenchmarkBigCacheSet(b *testing.B) {

	cache := bigcache.NewBigCache(b.N)
	for i := 0; i < b.N; i++ {
		cache.Set(key(i), value())
	}
}
