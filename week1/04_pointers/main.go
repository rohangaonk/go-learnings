package main

import "fmt"

func main() {
	// 1. Everything is passed by value (copied)
	age := 25
	updateAgeValue(age)
	fmt.Println("Age after updateAgeValue (copy):", age) // Still 25!

	// 2. To modify the ORIGINAL, we pass a Pointer
	// &age gives us the memory address of age (e.g., 0xc0000120b8)
	updateAgePointer(&age)
	fmt.Println("Age after updateAgePointer (reference):", age) // Now 30!

	// 3. Pointer Declaration
	// *int means "a pointer to an integer"
	var p *int = &age
	fmt.Printf("Pointer p looks like this (memory address): %v\n", p)
	fmt.Printf("The value at pointer p (*p) is: %d\n", *p) // "Dereferencing"

	// 4. Zero value of a pointer is nil
	var emptyPointer *string
	fmt.Printf("Empty pointer value: %v\n", emptyPointer)

	var empty *int
	*empty = 100

}

func updateAgeValue(a int) {
	a = 30 // This 'a' is a copy. The original is untouched.
}

func updateAgePointer(a *int) {
	// *a means "the value at address a"
	*a = 30 // We are modifying the memory address directly.
}
