package main

import "fmt"

func main() {
	fmt.Println("Start")

	// Defers are executed in LIFO (Last In, First Out) order.
	defer fmt.Println("Deferred: Cleanup 1 (First called, last run)")
	defer fmt.Println("Deferred: Cleanup 2")
	defer fmt.Println("Deferred: Cleanup 3 (Last called, first run)")

	fmt.Println("Doing some work...")
	
	PerformAction()

	fmt.Println("End of main")
}

func PerformAction() {
	defer fmt.Println("Deferred: Action cleanup")
	fmt.Println("Performing action...")
}
