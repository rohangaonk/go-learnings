package main

import "fmt"

func main() {
	fmt.Println("Starting server...")

	HandleRequest()

	fmt.Println("Server still running! Panic was caught.")
}

func HandleRequest() {
	// recover MUST be in a deferred function
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from panic: %v\n", r)
		}
	}()

	fmt.Println("Processing user request...")
	
	// Simulate a "bug" that triggers a panic
	CausePanic()

	fmt.Println("This line will NEVER run.")
}

func CausePanic() {
	panic("database connection lost!")
}
