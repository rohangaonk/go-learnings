package main

import "fmt"

// =============================================================================
// WEEK 2 — Lesson 1: Slices
//
// In JS, arrays are dynamic by default. In Go, you have two distinct things:
//   - Arrays  — fixed size, value type, rarely used directly
//   - Slices  — dynamic view over an underlying array; what you use 99% of the time
//
// A slice is a 3-field descriptor: (pointer, length, capacity)
// Understanding this mental model prevents most Go bugs.
// =============================================================================

func arrayVsSlice() {
	// Array — fixed size, type includes length [3]int != [4]int
	arr := [3]int{1, 2, 3}
	fmt.Println("Array:", arr)

	// Slice — dynamic, declared without a size
	s := []int{1, 2, 3}
	fmt.Println("Slice:", s)
	fmt.Printf("len=%d cap=%d\n", len(s), cap(s))
}

// -----------------------------------------------------------------------------
// append — the key operation on slices
// When capacity is exceeded, Go allocates a new (larger) backing array.
// This is important: after a reallocation, the new slice no longer shares
// memory with the original.
// -----------------------------------------------------------------------------
func appendDemo() {
	s := []int{1, 2, 3}
	fmt.Printf("Before append: len=%d cap=%d ptr=%p\n", len(s), cap(s), s)

	s = append(s, 4)
	fmt.Printf("After append:  len=%d cap=%d ptr=%p\n\n", len(s), cap(s), s)
	// Notice: capacity typically doubles, and pointer may change (new backing array)
}

// -----------------------------------------------------------------------------
// Shared backing array — the most common source of bugs
// -----------------------------------------------------------------------------
func sharedBackingArray() {
	original := []int{10, 20, 30}
	sub := original[1:] // sub shares original's backing array

	fmt.Println("Before mutation:", original)
	sub[0] = 99 // mutates original[1]
	fmt.Println("After sub[0]=99:", original) // [10 99 30]
}

// -----------------------------------------------------------------------------
// copy — use this when you want an independent slice
// -----------------------------------------------------------------------------
func copyDemo() {
	original := []int{1, 2, 3}
	clone := make([]int, len(original))
	copy(clone, original)

	clone[0] = 999
	fmt.Println("original:", original) // [1 2 3] — untouched
	fmt.Println("clone:   ", clone)    // [999 2 3]
}

// -----------------------------------------------------------------------------
// make — allocate a slice with a known size/capacity upfront
// Preferred over repeated append when final size is known (avoids reallocations)
// -----------------------------------------------------------------------------
func makeDemo() {
	// make([]T, len, cap)
	s := make([]int, 0, 5) // len=0, cap=5 — no reallocations until 5 elements
	for i := 0; i < 6; i++ {
		s = append(s, i*10)
		fmt.Printf("  len=%d cap=%d %v\n", len(s), cap(s), s)
	}
}

func reslicingDemo() {
    full := []int{0, 1, 2, 3, 4, 5}
    
    // Window starts at index 2
    window := full[1:3] 
    
    // 1. You CAN grow to the right (within capacity)
    expanded := window[:4] 
    fmt.Println(expanded)
    
    // 2. You CANNOT grow to the left
    // back := window[-1:] // Syntax Error: Negative index
    // back := window[0-1 : 2] // Panic: runtime error: slice bounds out of range
}


func main() {
	// fmt.Println("=== Array vs Slice ===")
	// arrayVsSlice()

	// fmt.Println("\n=== Append & Capacity Growth ===")
	// appendDemo()

	// fmt.Println("=== Shared Backing Array ===")
	// sharedBackingArray()

	// fmt.Println("\n=== copy() for Independence ===")
	// copyDemo()

	// fmt.Println("\n=== make() for Pre-allocation ===")
	// makeDemo()

	reslicingDemo()
}
