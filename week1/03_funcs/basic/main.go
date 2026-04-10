package main

import (
	"fmt"
	"strings"
)

// 1. Basic Function
func add(a int, b int) int {
	return a + b
}

// 2. Multiple Return Values (The Go Way)
// In Node.js, you'd return { upper, lower }
func transform(name string) (string, string) {
	return strings.ToUpper(name), strings.ToLower(name)
}

// 3. Named Return Values (Useful for documentation)
func divide(a, b float64) (result float64, err error) {
	if b == 0 {
		// We'll learn real errors next week, for now just a mockup
		return 0, fmt.Errorf("cannot divide by zero")
	}
	result = a / b
	return // Naked return: automatically returns result and err
}

func main() {
	sum := add(10, 5)
	fmt.Println("Sum:", sum)

	// Multiple returns destructuring
	up, low := transform("GoLanguage")
	fmt.Printf("Upper: %s, Lower: %s\n", up, low)

	// Ignoring a return value with "_"
	upperOnly, _ := transform("NodeEngineer")
	fmt.Println("Upper only:", upperOnly)

	// Named returns in action
	res, err := divide(10, 2)
	fmt.Printf("Division result: %f, Error: %v\n", res, err)
}
