package lru

import (
	"container/list"
)

type Cache struct {
	maxBytes int64 // 允许使用的最大内存
	nbytes int64   // 已经使用的内存
	ll *list.List
	cache map[string]*list.Element
	OnEvicted func(key string, value Value) // 记录移除回调
}



type Value interface {
	Len() int
}

type entry struct {
	key string
	value Value
}

func New(maxBytes int64, onEvicted func(string, Value)) *Cache {
	return &Cache{
		maxBytes:  maxBytes,
		ll:        list.New(),
		cache:     make(map[string]*list.Element),
		OnEvicted: onEvicted,
	}
}

func (c *Cache) Get(key string) (value Value, ok bool)  {
	ele, ok := c.cache[key]
	if ok {
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		return kv.value, ok
	}
	return
}

func (c *Cache) Len() int {
	return c.ll.Len()
}

func (c *Cache) RemoveOldest() {
	ele := c.ll.Back()
	if ele != nil {
		c.ll.Remove(ele)
		kv := ele.Value.(*entry)
		delete(c.cache, kv.key)
		c.nbytes -= getValueNbytes(kv)
		if c.OnEvicted != nil {
			c.OnEvicted(kv.key, kv.value)
		}
	}
}

func getValueNbytes(kv *entry) int64 {
	return int64(len(kv.key)) + int64(kv.value.Len())
}

func (c *Cache) Add(key string, value Value) {
	ele, ok := c.cache[key]
	if ok {
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		c.nbytes += getValueNbytes(kv)
		kv.value = value
	} else {
		entryP := &entry{
			key:  key,
			value: value,
		}
		ele := c.ll.PushFront(entryP)
		c.cache[key] = ele
		c.nbytes += getValueNbytes(entryP)
	}
	for c.maxBytes != 0 && c.maxBytes < c.nbytes {
		c.RemoveOldest()
	}
}










