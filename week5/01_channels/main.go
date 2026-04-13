package main

import "fmt"

func main() {
	// Create a channel
	ch := make(chan string)

	// Exercise:
	// 1. Start a goroutine that sends "ping" to 'ch'
	// 2. Receive the message from 'ch' in this main function and print it

	fmt.Println("Waiting for message...")
	
	// YOUR CODE HERE
}
