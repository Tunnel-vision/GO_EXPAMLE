package LRUCache

import (
	"sync"
)

// 设计实体 Entry
type Entry struct {
	Key   string
	Value interface{}
	pre   *Entry
	next  *Entry
}

//
type Cache struct {
	cache    map[string]*Entry
	capacity int
	head     *Entry
	tail     *Entry
}

func NewCache(cap int) *Cache {
	return &Cache{cache: make(map[string]*Entry), capacity: cap}
}

var lock sync.RWMutex

// 把元素放入到缓存中，如果缓存满得的话，则删除最后的，，把新的放入缓存中
// 如果缓存没满，则把新元素放入缓存中 并返回nil
func (self *Cache) Put(key string, Val interface{}) interface{} {
	lock.Lock()
	defer lock.Unlock()

	// 已经存在的话，移到最前面
	if existval, ok := self.cache[key]; ok {
		self.moveToHead(existval)
		return nil
	}
	e := &Entry{Key: key, Value: Val, next: self.head}
	// 不存在的话，// 2 不满的话，则把新的放入最前
	if self.head != nil {
		self.head.pre = e
	}
	self.head = e
	if self.tail == nil {
		self.tail = e
	}
	self.cache[key] = e

	if len(self.cache) <= self.capacity {
		return nil
	}
	// 1 满的话，删除最后的，把新的放入最前
	removeEntry := self.tail
	self.tail = self.tail.pre
	removeEntry.pre = nil
	self.tail.next = nil
	delete(self.cache, removeEntry.Key)
	return removeEntry.Value

}

func (self *Cache) moveToHead(e *Entry) {
	// 元素存在，下面把元素放到最前面去
	if e == self.head {
		return
	}
	// 元素不是head，那么e.pre一定不为nil
	e.pre.next = e.next
	if e == self.tail {
		self.tail = e.pre
	} else {
		e.next.pre = e.pre
	}
	e.pre = nil
	e.next = self.head
	self.head.pre = e
	self.head = e

	//// 根据元素目前存在的位置 来确定怎么移动
	//// 元素存在 把元素放在最前面
	//if entry == self.head {
	//	return
	//}
	//entry.pre.next = entry.next
	//
	//// 元素在队尾
	//if entry == self.tail {
	//	self.tail = entry.pre
	//
	//} else {
	//	entry.next.pre = entry.pre
	//}
	//entry.next.pre = entry
	//entry.pre = nil
	//entry.next = self.head
	//self.head = entry

}

func (self *Cache) Get(Key string) interface{} {
	lock.Lock()
	defer lock.Unlock()
	if existVal, ok := self.cache[Key]; ok {
		self.moveToHead(existVal)
		return existVal.Value
	}
	return nil
}
