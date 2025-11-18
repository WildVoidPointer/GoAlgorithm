package lrustack

import "fmt"

// SingleStackLRU 优化的LRU缓存实现
type SingleStackLRU struct {
	capacity int
	stack    []int       // 使用单个栈存储访问顺序
	cache    map[int]int // 用于快速查找已加载页
}

// NewSingleStackLRU 创建优化的LRU缓存
func NewSingleStackLRU(capacity int) *SingleStackLRU {
	return &SingleStackLRU{
		capacity: capacity,
		stack:    make([]int, 0),
		cache:    make(map[int]int),
	}
}

// Get 获取键值，如果存在则移到栈顶
func (lru *SingleStackLRU) Get(key int) int {
	if value, exists := lru.cache[key]; exists {
		lru.moveToTop(key)
		return value
	}
	return -1
}

// Put 插入或更新键值对
func (lru *SingleStackLRU) Put(key int, value int) {
	if _, exists := lru.cache[key]; exists {
		lru.cache[key] = value
		lru.moveToTop(key)
		return
	}

	if len(lru.stack) >= lru.capacity {
		lru.removeLeastRecentlyUsed()
	}

	lru.cache[key] = value
	lru.stack = append(lru.stack, key)
}

// moveToTop 将元素移到栈顶
func (lru *SingleStackLRU) moveToTop(key int) {
	// 找到key在栈中的位置
	index := -1
	for i, k := range lru.stack {
		if k == key {
			index = i
			break
		}
	}

	if index == -1 {
		return
	}

	// 将元素移到栈顶
	lru.stack = append(lru.stack[:index], lru.stack[index+1:]...)
	lru.stack = append(lru.stack, key)
}

// removeLeastRecentlyUsed 移除最久未使用的元素
func (lru *SingleStackLRU) removeLeastRecentlyUsed() {
	if len(lru.stack) == 0 {
		return
	}

	lruKey := lru.stack[0]
	lru.stack = lru.stack[1:]
	delete(lru.cache, lruKey)
}

// Print 打印当前状态
func (lru *SingleStackLRU) Print() {
	fmt.Printf("访问顺序: %v\n", lru.stack)
	fmt.Printf("缓存内容: %v\n", lru.cache)
	fmt.Println("---")
}
