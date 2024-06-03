package solutions

import "container/list"

//Explanation
//LRUCache lRUCache = new LRUCache(2);
//lRUCache.put(1, 1); // cache is {1=1}
//lRUCache.put(2, 2); // cache is {1=1, 2=2}
//lRUCache.get(1);    // return 1
//lRUCache.put(3, 3); // LRU key was 2, evicts key 2, cache is {1=1, 3=3}
//lRUCache.get(2);    // returns -1 (not found)
//lRUCache.put(4, 4); // LRU key was 1, evicts key 1, cache is {4=4, 3=3}
//lRUCache.get(1);    // return -1 (not found)
//lRUCache.get(3);    // return 3
//lRUCache.get(4);    // return 4

type LRU struct {
	capacity int
	list     *list.List
	cash     map[int]*list.Element
}

type entry struct {
	k, v int
}

func NewLRU(capacity int) *LRU {
	return &LRU{
		capacity: capacity,
		cash:     make(map[int]*list.Element),
		list:     list.New(),
	}
}

func (lru *LRU) Put(key int, value int) {
	if elem, ok := lru.cash[key]; ok {
		lru.list.MoveToFront(elem)
		elem.Value.(*entry).v = value
		return
	}

	if lru.list.Len() == lru.capacity {
		back := lru.list.Back()
		if back != nil {
			lru.list.Remove(back)
			delete(lru.cash, back.Value.(*entry).k)
		}
	}

	el := lru.list.PushFront(&entry{key, value})
	lru.cash[key] = el
}

func (lru *LRU) Get(key int) int {
	if elem, ok := lru.cash[key]; ok {
		lru.list.MoveToFront(elem)
		return elem.Value.(*entry).v
	}
	return -1
}
