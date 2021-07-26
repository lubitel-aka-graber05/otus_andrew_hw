package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

/*func (lc *lruCache) Set(key Key, value interface{}) bool {
	_, ok := lc.items[key]
	if !ok {
		if len(lc.items) == lc.capacity {
			e := lc.queue.Back()
			lc.queue.Remove(e)
			keyName := e.Value.(cacheItem).key
			delete(lc.items, Key(keyName))
		}else {
			item := cacheItem{key: string(key), value: value}
			lc.queue.PushFront(item)
			lc.items[key] = value
			return false
		}
	}
	return true
}*/

func (lc *lruCache) Set(key Key, value interface{}) bool {
	// ci := &cacheItem{key: string(key), value: value}

	if el, ok := lc.items[key]; ok {
		lc.queue.MoveToFront(el)
		el.Value = value
		return true
	}
	if len(lc.items) == lc.capacity {
		lc.Clear()
	}

	el := lc.queue.PushFront(value)
	lc.items[key] = el

	return false
}

func (lc *lruCache) Clear() {
	lc.queue.Remove(lc.queue.Back())
	delete(lc.items, "a")
}

func (lc *lruCache) Get(key Key) (interface{}, bool) {
	el, ok := lc.items[key]
	if !ok {
		return nil, false
	}
	lc.queue.MoveToFront(el)
	return el.Value, true
}

type lruCache struct {
	// Cache
	capacity int
	queue    List
	items    map[Key]*ListItem
}

/*type cacheItem struct {
	key   string
	value interface{}
}*/

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
