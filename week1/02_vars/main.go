package main

import "fmt"

func main() {
	// 1. Explicit declaration with var
	var name string = "Go"
	var version float64 = 1.22
	
	// 2. Type inference with var
	var isAwesome = true // bool inferred
	
	// 3. Short declaration (only inside functions)
	// Equivalent to 'let' in JS, but type-safe
	age := 10 
	
	// 4. Zero Values
	// In JS: let x; -> undefined
	// In Go: everything has a default value
	var i int     // 0
	var f float64 // 0
	var b bool    // false
	var s string  // "" (empty string, not null/undefined)
	
	fmt.Printf("Name: %s, Version: %f, Age: %d, Awesome: %v\n", name, version, age, isAwesome)
	fmt.Printf("Zero Values -> int: %d, float: %f, bool: %v, string: %q\n", i, f, b, s)
}
