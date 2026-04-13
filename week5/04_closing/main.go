package main

import (
	"fmt"
)

func generateNumbers(ch chan<- int) {
	for i := 1; i <= 3; i++ {
		ch <- i
	}
	// YOUR CODE HERE: What needs to happen right here so the main function knows we are done?
	close(ch)
}

func main() {
	ch := make(chan int)

	go generateNumbers(ch)

	// YOUR CODE HERE: Use a `for val := range ch` loop to read all values 
	// until the channel is closed.

	for val := range ch {
		fmt.Println("Received", val)
	}
	
}
