package main

import (
	"container/list"
)

type LRUCache struct {
	capacity   int
	cache      *list.List
	keyToValue map[int]*list.Element
}

type entry struct {
	key   int
	value int
}

// Initialize the LRUCache with a given capacity.
func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity:   capacity,
		cache:      list.New(),
		keyToValue: make(map[int]*list.Element),
	}
}

// Get the value (will always be positive) of the key if the key exists in the cache, otherwise return -1.
func (cache *LRUCache) Get(key int) int {
	if el, found := cache.keyToValue[key]; found {
		cache.cache.MoveToFront(el)
		return el.Value.(*entry).value
	}
	return -1
}

// Set or insert the value if the key is not already present. When the cache reached its capacity, it should invalidate the least recently used item before inserting a new item.
func (cache *LRUCache) Put(key int, value int) {
	if el, found := cache.keyToValue[key]; found {
		cache.cache.MoveToFront(el)
		el.Value.(*entry).value = value
	} else {
		if cache.cache.Len() == cache.capacity {
			// Remove the least recently used item
			oldest := cache.cache.Back()
			if oldest != nil {
				cache.cache.Remove(oldest)
				delete(cache.keyToValue, oldest.Value.(*entry).key)
			}
		}
		// Insert the new item at the front
		el := cache.cache.PushFront(&entry{key: key, value: value})
		cache.keyToValue[key] = el
	}
}
