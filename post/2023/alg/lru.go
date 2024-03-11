package alg

import (
	"sync"
)

type LRUCache struct {
	capacity, size int
	head, tail     *LinkListNode
	m              map[interface{}]*LinkListNode
	rwmutx         *sync.RWMutex
}

type LinkListNode struct {
	prev, next *LinkListNode
	key, value interface{}
}

func Constructor(capacity int) *LRUCache {
	head := &LinkListNode{}
	tail := &LinkListNode{}
	head.next = tail
	tail.prev = head
	return &LRUCache{
		capacity: capacity,
		head:     head,
		tail:     tail,
		m:        make(map[interface{}]*LinkListNode),
		rwmutx:   &sync.RWMutex{},
	}
}

func (l *LRUCache) Get(key int) int {
	l.rwmutx.RLock()
	defer l.rwmutx.RUnlock()
	v, ok := l.m[key]
	if ok {
		l.MoveToFront(v)
		return v.value.(int)
	}
	return -1
}

func (l *LRUCache) Put(key interface{}, value interface{}) {
	l.rwmutx.Lock()
	defer l.rwmutx.Unlock()
	v, exists := l.m[key]
	if !exists {
		// 容量不够
		if l.size >= l.capacity {
			d := l.tail.prev
			l.Remove(d)
			delete(l.m, d.key)
			l.size--
		}
		e := &LinkListNode{
			key:   key,
			value: value,
		}
		l.AddToFront(e)
		l.m[key] = e
		l.size++
	} else {
		v.value = value
		l.MoveToFront(v)
	}
}

func (l *LRUCache) AddToFront(node *LinkListNode) {
	node.prev = l.head
	node.next = l.head.next
	l.head.next.prev = node
	l.head.next = node
}

func (l *LRUCache) Remove(node *LinkListNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (l *LRUCache) MoveToFront(node *LinkListNode) {
	l.Remove(node)
	l.AddToFront(node)
}

// type LRUCache struct {
// 	capacity int
// 	list     *list.List
// 	maps     map[int]*list.Element
// }

// type Ele struct {
// 	key   int
// 	value int
// }

// func Constructor(capacity int) LRUCache {
// 	return LRUCache{capacity: capacity, list: list.New().Init(), maps: make(map[int]*list.Element)}
// }

// func (cache *LRUCache) Get(key int) int {
// 	e, ok := cache.maps[key]
// 	if !ok {
// 		return -1
// 	}

// 	cache.list.MoveToFront(e)
// 	return e.Value.(Ele).value
// }

// func (cache *LRUCache) Put(key int, value int) {
// 	// if not exists,insert to front
// 	e, ok := cache.maps[key]
// 	if !ok {
// 		// if over cap,delete last
// 		if cache.list.Len() >= cache.capacity {
// 			delEle := cache.list.Remove(cache.list.Back())
// 			delete(cache.maps, delEle.(Ele).key)

// 		}
// 		e = cache.list.PushFront(Ele{key: key, value: value})
// 		cache.maps[key] = e
// 	} else {
// 		// if exists,set new value ,move to front
// 		temp := e.Value.(Ele)
// 		temp.value = value
// 		cache.list.MoveToFront(e)
// 	}
// }
