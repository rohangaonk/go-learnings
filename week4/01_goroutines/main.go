package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// A WaitGroup waits for a collection of goroutines to finish.
	// The main goroutine calls Add to set the number of
	// goroutines to wait for. Then each of the goroutines
	// runs and calls Done when finished. At the same time,
	// Wait can be used to block until all goroutines have finished.
	var wg sync.WaitGroup

	numWorkers := 5

	// 1. Tell the WaitGroup we are waiting for numWorkers goroutines
	wg.Add(numWorkers)

	for i := 1; i <= numWorkers; i++ {
		// Launch a goroutine for each worker
		// ⚠️ Gotcha: We pass 'i' as an argument to the anonymous function.
		// If we used 'i' directly inside the closure, all goroutines
		// might see the same value (the final value of i) because they
		// share the same variable. This is a classic Go concurrency pitfall!
		go func(workerID int) {
			// 2. Schedule wg.Done() to be called when the function exits
			defer wg.Done()

			fmt.Printf("Worker %d started\n", workerID)

			// Simulate some work
			time.Sleep(time.Millisecond * 100)

			fmt.Printf("Worker %d done\n", workerID)
		}(i)
	}

	// 3. Block here until the WaitGroup counter goes back to 0
	fmt.Println("Main: Waiting for workers...")
	wg.Wait()

	fmt.Println("Main: All workers finished. Exiting.")
}
