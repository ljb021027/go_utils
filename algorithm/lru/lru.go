package main

import (
	"fmt"
)

type Node struct {
	key   int
	value int
	prev  *Node
	next  *Node
}

type LRUCache struct {
	capacity int
	cache    map[int]*Node
	head     *Node
	tail     *Node
}

func NewLRUCache(capacity int) *LRUCache {
	head := &Node{}
	tail := &Node{}
	head.next = tail
	tail.prev = head
	return &LRUCache{
		capacity: capacity,
		cache:    make(map[int]*Node),
		head:     head,
		tail:     tail,
	}
}

func (c *LRUCache) Get(key int) int {
	if node, ok := c.cache[key]; ok {
		c.moveToFront(node)
		return node.value
	}
	return -1
}

func (c *LRUCache) Put(key, value int) {
	if node, ok := c.cache[key]; ok {
		node.value = value
		c.moveToFront(node)
	} else {
		node := &Node{key: key, value: value}
		if len(c.cache) >= c.capacity {
			delete(c.cache, c.tail.prev.key)
			c.removeNode(c.tail.prev)
		}
		c.cache[key] = node
		c.addToFront(node)
	}
}

func (c *LRUCache) addToFront(node *Node) {
	node.prev = c.head
	node.next = c.head.next
	c.head.next.prev = node
	c.head.next = node
}

func (c *LRUCache) removeNode(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (c *LRUCache) moveToFront(node *Node) {
	c.removeNode(node)
	c.addToFront(node)
}

func main() {
	cache := NewLRUCache(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	fmt.Println(cache.Get(1))
	cache.Put(3, 3)
	fmt.Println(cache.Get(2))
	cache.Put(4, 4)
	fmt.Println(cache.Get(1))
	fmt.Println(cache.Get(3))
	fmt.Println(cache.Get(4))
}
