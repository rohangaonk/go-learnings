package main

import "fmt"

func main() {
	// Create a channel
	ch := make(chan string)

	// Exercise:
	// 1. Start a goroutine that sends "ping" to 'ch'
	// 2. Receive the message from 'ch' in this main function and print it

	go func() {
		fmt.Println("Go routine watinig for receiver")
		ch <- "Hello from go routine"
	}();


	fmt.Println("Waiting for message...")

	msg := <-ch

	fmt.Printf("message is %s", msg)
	
	// YOUR CODE HERE
}
