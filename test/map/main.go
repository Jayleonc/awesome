package main

import (
	"fmt"
	"sync"
)

// MapWithSlice 是包含切片的 Map 结构体
type MapWithSlice struct {
	mu   sync.Mutex
	keys []string
	data map[string]interface{}
}

// NewMapWithSlice 创建一个新的 MapWithSlice 实例
func NewMapWithSlice() *MapWithSlice {
	return &MapWithSlice{
		keys: make([]string, 0),
		data: make(map[string]interface{}),
	}
}

// Set 设置键值对
func (m *MapWithSlice) Set(key string, value interface{}) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.keys = append(m.keys, key)
	m.data[key] = value
}

// GetByKey 根据键获取值
func (m *MapWithSlice) GetByKey(key string) interface{} {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.data[key]
}

// GetByIndex 根据切片索引获取值
func (m *MapWithSlice) GetByIndex(index int) (string, interface{}) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if index < 0 || index >= len(m.keys) {
		return "", nil
	}
	key := m.keys[index]
	value := m.data[key]
	return key, value
}

// Size 返回 MapWithSlice 的大小
func (m *MapWithSlice) Size() int {
	m.mu.Lock()
	defer m.mu.Unlock()
	return len(m.keys)
}

func main() {
	mws := NewMapWithSlice()

	// 设置键值对
	mws.Set("key1", 123)
	mws.Set("key2", "hello")
	mws.Set("key3", []int{1, 2, 3})

	// 遍历切片并获取 Map 的值
	size := mws.Size()
	for i := 0; i < size; i++ {
		key, value := mws.GetByIndex(i)
		fmt.Printf("Key: %s, Value: %v\n", key, value)
	}
}
