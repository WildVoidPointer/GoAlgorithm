package lrustack

import "fmt"

// DoubleStackLRU 使用两个栈实现的LRU缓存
type DoubleStackLRU struct {
	capacity int
	// 主栈，存储键值对，栈顶是最新访问的元素
	mainStack []int
	// 辅助栈，用于临时存储
	tempStack []int
	// 存储键值对
	cache map[int]int
}

// NewDoubleStackLRU 创建新的LRU缓存
func NewDoubleStackLRU(capacity int) *DoubleStackLRU {
	return &DoubleStackLRU{
		capacity:  capacity,
		mainStack: make([]int, 0),
		tempStack: make([]int, 0),
		cache:     make(map[int]int),
	}
}

// Get 获取键对应的值，如果存在则将其移到栈顶
func (lru *DoubleStackLRU) Get(key int) int {
	if value, exists := lru.cache[key]; exists {
		// 将key移到栈顶
		lru.moveToTop(key)
		return value
	}
	return -1
}

// Put 插入键值对，如果键已存在则更新并移到栈顶
func (lru *DoubleStackLRU) Put(key int, value int) {
	// 如果键已存在，更新值并移到栈顶
	if _, exists := lru.cache[key]; exists {
		lru.cache[key] = value
		lru.moveToTop(key)
		return
	}

	// 如果缓存已满，移除最久未使用的元素（栈底）
	if len(lru.mainStack) >= lru.capacity {
		lru.removeLeastRecentlyUsed()
	}

	// 添加新元素到栈顶
	lru.cache[key] = value
	lru.mainStack = append(lru.mainStack, key)
}

// moveToTop 将指定key移到栈顶
func (lru *DoubleStackLRU) moveToTop(key int) {
	// 使用辅助栈找到目标key
	for len(lru.mainStack) > 0 {
		top := lru.mainStack[len(lru.mainStack)-1]
		lru.mainStack = lru.mainStack[:len(lru.mainStack)-1]

		if top == key {
			// 找到目标key，将其放在栈顶
			break
		}
		// 其他元素暂存到辅助栈
		lru.tempStack = append(lru.tempStack, top)
	}

	// 将辅助栈中的元素放回主栈
	for len(lru.tempStack) > 0 {
		top := lru.tempStack[len(lru.tempStack)-1]
		lru.tempStack = lru.tempStack[:len(lru.tempStack)-1]
		lru.mainStack = append(lru.mainStack, top)
	}

	// 将目标key放到栈顶
	lru.mainStack = append(lru.mainStack, key)
}

// removeLeastRecentlyUsed 移除最久未使用的元素（栈底）
func (lru *DoubleStackLRU) removeLeastRecentlyUsed() {
	if len(lru.mainStack) == 0 {
		return
	}

	// 栈底元素是最久未使用的
	lruKey := lru.mainStack[0]

	// 移除栈底元素
	lru.mainStack = lru.mainStack[1:]

	// 从缓存中删除
	delete(lru.cache, lruKey)
}

// Print 打印当前缓存状态
func (lru *DoubleStackLRU) Print() {
	fmt.Printf("Stack: %v\n", lru.mainStack)
	fmt.Printf("Cache: %v\n", lru.cache)
	fmt.Println("---")
}
