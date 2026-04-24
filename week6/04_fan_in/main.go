package main

import (
	"fmt"
	"sync"
	"time"
)

// fanIn takes any number of receive-only channels and merges them into one
func fanIn(sources ...<-chan string) <-chan string {
	var wg sync.WaitGroup
	out := make(chan string)

	// Helper function that forwards data from one channel
	forward := func(c <-chan string) {
		defer wg.Done()
		for val := range c {
			out <- val
		}
	}

	wg.Add(len(sources))
	for _, c := range sources {
		// Fan-out a goroutine for each source channel
		go forward(c)
	}

	// Wait for all forwarding goroutines to finish in the background
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func produce(name string, delay time.Duration) <-chan string {
	c := make(chan string)
	go func() {
		for i := 1; i <= 3; i++ {
			time.Sleep(delay)
			c <- fmt.Sprintf("%s: Msg %d", name, i)
		}
		close(c)
	}()
	return c
}

func main() {
	// Create two distinct sources
	c1 := produce("Source A", 100*time.Millisecond)
	c2 := produce("Source B", 150*time.Millisecond)

	// Merge them together
	merged := fanIn(c1, c2)

	// Consume the unified stream
	for msg := range merged {
		fmt.Println(msg)
	}
	fmt.Println("All done!")
}
