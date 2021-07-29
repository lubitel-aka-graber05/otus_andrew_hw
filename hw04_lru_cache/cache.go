package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

func (lc *lruCache) Set(key Key, value interface{}) bool {
	if node, ok := lc.items[key]; ok {
		lc.queue.MoveToFront(node)
		node.Value.(*ListItem).Value = cacheItem{key: string(key), value: value}
		return true
	}
	if lc.queue.Len() == lc.capacity {
		lc.delLastElement()
	}
	node := &ListItem{
		Value: cacheItem{
			key:   string(key),
			value: value,
		},
	}
	ptr := lc.queue.PushFront(node)
	lc.items[key] = ptr

	return false
}

func (lc *lruCache) delLastElement() {
	k := lc.queue.Back().Value.(*ListItem).Value.(cacheItem).key
	delete(lc.items, Key(k))
	lc.queue.Remove(lc.queue.Back())
}

func (lc *lruCache) Clear() {
	lc.capacity = 0
	lc.queue = nil
	lc.items = nil
}

func (lc *lruCache) Get(key Key) (interface{}, bool) {
	if node, ok := lc.items[key]; ok {
		val := node.Value.(*ListItem).Value.(cacheItem).value
		lc.queue.MoveToFront(node)
		return val, true
	}
	return nil, false
}

type lruCache struct {
	// Cache
	capacity int
	queue    List
	items    map[Key]*ListItem
}

type cacheItem struct {
	key   string
	value interface{}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
