package main

import (
	"fmt"
	"time"
)

func main() {
	dbQuery := make(chan string)
	apiCall := make(chan string)

	go func() {
		time.Sleep(200 * time.Millisecond)
		dbQuery <- "DB Result"
	}()

	go func() {
		time.Sleep(100 * time.Millisecond)
		apiCall <- "API Result"
	}()

	// Assessment: We want to retrieve the FIRST result that comes back, 
	// OR timeout and print an error if nothing comes back within 150 milliseconds.
	// How do we use the 'select' statement here?
	
	// YOUR CODE HERE
	fmt.Println("Fill in the select statement!")

	select {
	case <-dbQuery:
		fmt.Println("received db query")
	case <-apiCall:
		fmt.Println("received api call")
	case <-time.After(time.Millisecond * 150):
		fmt.Println("Too slow")
	}
}
