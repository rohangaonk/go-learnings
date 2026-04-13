package main

import "fmt"

// =============================================================================
// WEEK 2 — Lesson 4: Control Flow
//
// Go has exactly ONE loop keyword: `for`
// It replaces: for, while, do-while, for..of, for..in from JS
//
// `switch` is cleaner than JS — no fallthrough by default, no break needed.
// =============================================================================

// -----------------------------------------------------------------------------
// for — three forms
// -----------------------------------------------------------------------------
func forLoops() {
	// 1. Classic C-style for
	for i := 0; i < 3; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()

	// 2. While-style (condition only) — equivalent to JS while(cond)
	n := 1
	for n < 100 {
		n *= 2
	}
	fmt.Println("First power of 2 >= 100:", n)

	// 3. Infinite loop — equivalent to JS while(true)
	count := 0
	for {
		count++
		if count == 3 {
			break
		}
	}
	fmt.Println("Broke at count:", count)
}

// -----------------------------------------------------------------------------
// range — iterate over slices, maps, strings, channels
// JS equivalent: for..of (slices), Object.entries() (maps)
// -----------------------------------------------------------------------------
func rangeExamples() {
	// Slice — index + value
	fruits := []string{"apple", "banana", "cherry"}
	for i, v := range fruits {
		fmt.Printf("  [%d] %s\n", i, v)
	}

	// Ignore index with _
	for _, v := range fruits {
		fmt.Println(" ", v)
	}

	// Map — key + value (unordered)
	scores := map[string]int{"alice": 90, "bob": 75}
	for name, score := range scores {
		fmt.Printf("  %s: %d\n", name, score)
	}

	// String — ranges over Unicode runes (not bytes!)
	for i, ch := range "Go🚀" {
		fmt.Printf("  index=%d rune=%c\n", i, ch)
	}
}

// -----------------------------------------------------------------------------
// switch — no fallthrough by default (unlike JS!)
// No `break` needed. Use `fallthrough` keyword explicitly if you want it.
// Cases can have multiple values, expressions, or no condition at all.
// -----------------------------------------------------------------------------
func switchExamples() {
	// Basic switch
	day := "Monday"
	switch day {
	case "Saturday", "Sunday":
		fmt.Println(day, "is a weekend")
	case "Monday":
		fmt.Println(day, "is the worst")
	default:
		fmt.Println(day, "is a weekday")
	}

	// Switch without a condition — cleaner if/else chain
	score := 82
	switch {
	case score >= 90:
		fmt.Println("Grade: A")
	case score >= 80:
		fmt.Println("Grade: B")
	case score >= 70:
		fmt.Println("Grade: C")
	default:
		fmt.Println("Grade: F")
	}

	// Type switch — very useful with interfaces (preview for Week 3)
	values := []interface{}{42, "hello", true, 3.14}
	for _, v := range values {
		switch t := v.(type) {
		case int:
			fmt.Printf("  int: %d\n", t)
		case string:
			fmt.Printf("  string: %q\n", t)
		case bool:
			fmt.Printf("  bool: %v\n", t)
		default:
			fmt.Printf("  other: %T\n", t)
		}
	}
}

func main() {
	fmt.Println("=== For Loops ===")
	forLoops()

	fmt.Println("\n=== Range ===")
	rangeExamples()

	fmt.Println("\n=== Switch ===")
	switchExamples()
}
