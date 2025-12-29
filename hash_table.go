package dbmsgo

import (
	"fmt"
)

type HashEntry struct {
	Key   string
	Value string
	Next  *HashEntry
}

type HashTable struct {
	name     string
	buckets  []*HashEntry
	capacity int
	size     int
}

func NewHashTable(name string) *HashTable {
	capacity := 10
	return &HashTable{
		name:     name,
		buckets:  make([]*HashEntry, capacity),
		capacity: capacity,
		size:     0,
	}
}

func (h *HashTable) hashFunction(key string) int {
	hash := 5381
	for _, c := range key {
		hash = ((hash << 5) + hash) + int(c)
	}
	// Исправление: убедимся, что индекс неотрицательный
	result := hash % h.capacity
	if result < 0 {
		return -result
	}
	return result
}

func (h *HashTable) hashFunctionWithCapacity(key string, cap int) int {
	hash := 5381
	for _, c := range key {
		hash = ((hash << 5) + hash) + int(c)
	}
	result := hash % cap
	if result < 0 {
		return -result
	}
	return result
}

func (h *HashTable) resize(newCapacity int) {
	newBuckets := make([]*HashEntry, newCapacity)
	
	for i := 0; i < h.capacity; i++ {
		current := h.buckets[i]
		for current != nil {
			next := current.Next
			newIndex := h.hashFunctionWithCapacity(current.Key, newCapacity)
			current.Next = newBuckets[newIndex]
			newBuckets[newIndex] = current
			current = next
		}
	}
	
	h.buckets = newBuckets
	h.capacity = newCapacity
}

func (h *HashTable) Insert(key, value string) {
	if h.size >= int(float64(h.capacity)*0.7) {
		h.resize(h.capacity * 2)
	}
	
	index := h.hashFunction(key)
	current := h.buckets[index]
	
	for current != nil {
		if current.Key == key {
			current.Value = value
			return
		}
		current = current.Next
	}
	
	newEntry := &HashEntry{Key: key, Value: value, Next: h.buckets[index]}
	h.buckets[index] = newEntry
	h.size++
}

func (h *HashTable) Search(key string) (string, bool) {
	index := h.hashFunction(key)
	current := h.buckets[index]
	
	for current != nil {
		if current.Key == key {
			return current.Value, true
		}
		current = current.Next
	}
	
	return "", false
}

func (h *HashTable) Remove(key string) bool {
	index := h.hashFunction(key)
	current := h.buckets[index]
	var prev *HashEntry
	
	for current != nil {
		if current.Key == key {
			if prev != nil {
				prev.Next = current.Next
			} else {
				h.buckets[index] = current.Next
			}
			h.size--
			return true
		}
		prev = current
		current = current.Next
	}
	
	return false
}

func (h *HashTable) Print() {
	fmt.Printf("Хеш-таблица '%s':\n", h.name)
	for i := 0; i < h.capacity; i++ {
		fmt.Printf("  [%d]: ", i)
		current := h.buckets[i]
		if current == nil {
			fmt.Printf("NULL")
		} else {
			for current != nil {
				fmt.Printf("{%s: %s}", current.Key, current.Value)
				if current.Next != nil {
					fmt.Printf(" -> ")
				}
				current = current.Next
			}
		}
		fmt.Println()
	}
}

func (h *HashTable) GetSize() int {
	return h.size
}

func (h *HashTable) IsEmpty() bool {
	return h.size == 0
}

func (h *HashTable) GetName() string {
	return h.name
}

func (h *HashTable) GetCapacity() int {
	return h.capacity
}

func (h *HashTable) GetBuckets() []*HashEntry {
	return h.buckets
}

func (h *HashTable) Cleanup() {
	for i := 0; i < h.capacity; i++ {
		h.buckets[i] = nil
	}
	h.size = 0
}
