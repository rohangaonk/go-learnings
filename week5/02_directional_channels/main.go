package main

import "fmt"

func worker(jobs chan<- int, results <-chan int) {
	// Assessment: What is wrong here?
	job := <-jobs
	results <- job * 2
}

func main() {
	jobs := make(chan int, 1)
	results := make(chan int, 1)

	// Attempting to use the worker
	go worker(jobs, results)

	jobs <- 5
	fmt.Println(<-results)
}
