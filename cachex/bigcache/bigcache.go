package bigcache

import (
	"github.com/allegro/bigcache/v2"
	"time"
)

const maxEntrySize = 256

func NewBigCache(entriesInWindow int) *bigcache.BigCache {

	cache, _ := bigcache.NewBigCache(bigcache.Config{
		Shards:             256,
		LifeWindow:         10 * time.Minute,
		MaxEntriesInWindow: entriesInWindow,
		MaxEntrySize:       maxEntrySize,
		Verbose:            true,
	})

	return cache
}
