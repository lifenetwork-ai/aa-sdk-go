package aasdk

import (
	"fmt"

	lru "github.com/hashicorp/golang-lru"
)

// LRUCache is an interface for caching smart account addresses.
// The implementation should be thread-safe.
type LRUCache interface {
	// Get retrieves the address from the cache.
	Get(key string) (any, bool)

	// Set stores the address in the cache.
	// Returns true if the key evicted an existing value.
	Set(key string, value any) bool
}

type lruCache struct {
	inner *lru.Cache
}

func NewLRUCache(maxSize int) LRUCache {
	cache, err := lru.New(maxSize)
	if err != nil {
		panic(fmt.Errorf("failed to create LRU cache: %w, maxSize: %d", err, maxSize))
	}
	return &lruCache{
		inner: cache,
	}
}

// Get implements LRUCache.
func (l *lruCache) Get(key string) (any, bool) {
	value, ok := l.inner.Get(key)
	return value, ok
}

// Set implements LRUCache.
func (l *lruCache) Set(key string, value any) bool {
	evicted := l.inner.Add(key, value)
	return evicted
}

var _ LRUCache = &lruCache{}
