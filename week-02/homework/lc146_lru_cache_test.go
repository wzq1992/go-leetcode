package homework

import "testing"

func TestLRUCache_Get(t *testing.T) {
	lruCache := NewLRUCache()

	lruCache.Put(1, 1)

	if got := lruCache.Get(1); got != 1 {
		t.Errorf("lruCache.Get(1) error")
	}
}

func TestLRUCache_Put(t *testing.T) {
	lruCache := NewLRUCache()

	lruCache.Put(1, 1)
	lruCache.Put(2, 2)
	lruCache.Put(3, 3)

	if got := lruCache.Get(1); got != -1 {
		t.Error("key: 1, value: 1 expired, but got it")
	}

	lruCache.Put(3, 8)
	if got := lruCache.Get(3); got != 8 {
		t.Error("key: 3, value: 3 update to 8, but got not 8")
	}
}

func NewLRUCache() *LRUCache {
	lruCache := Constructor(2)
	return &lruCache
}
