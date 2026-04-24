package main

import (
	"fmt"
	"sync"
	"time"
)

// 1. Using a standard sync.Mutex
type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.v[key]++
}

// 2. Using an RWMutex for cache pattern
type SafeCache struct {
	mu   sync.RWMutex
	data map[string]string
}

func (c *SafeCache) Get(key string) string {
	c.mu.RLock() // Multiple goroutines can read at the same time
	defer c.mu.RUnlock()
	return c.data[key]
}

func (c *SafeCache) Set(key, val string) {
	c.mu.Lock() // Exclusive write lock block everyone else
	defer c.mu.Unlock()
	c.data[key] = val
}

func main() {
	// Example usage of SafeCounter
	counter := SafeCounter{v: make(map[string]int)}
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Inc("somekey")
		}()
	}
	wg.Wait()
	fmt.Printf("Counter value (Protected by Mutex): %d\n", counter.v["somekey"])

	// Example usage of SafeCache
	cache := SafeCache{data: make(map[string]string)}
	cache.Set("status", "Learning Sync Primitives")
	
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("Goroutine %d reads from cache (RWMutex): %s\n", id, cache.Get("status"))
		}(i)
	}
	wg.Wait()
	
	// Adding a small delay just to keep main running for goroutine output flush in some systems
	time.Sleep(10 * time.Millisecond)
}
