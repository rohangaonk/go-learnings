package main

import (
	"fmt"
	"sync"
)

var (
	db     string
	once   sync.Once
)

func loadConfig(id int) {
	fmt.Printf("Goroutine %d: trying to load config\n", id)
	once.Do(func() {
		fmt.Printf("--- Goroutine %d: Actually loading config now ---\n", id)
		db = "Postgres-Production"
	})
}

func main() {
	var wg sync.WaitGroup

	// Multiple goroutines try to load config
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			loadConfig(id)
		}(i)
	}

	wg.Wait()
	fmt.Println("Final Config Value:", db)
}
