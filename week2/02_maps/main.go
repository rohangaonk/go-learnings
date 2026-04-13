package main

import "fmt"

// =============================================================================
// WEEK 2 — Lesson 2: Maps
//
// Go maps are like JS objects/Maps — key-value stores with O(1) average lookup.
// Key differences:
//   - Keys must be a comparable type (string, int, etc. — NOT slices/maps)
//   - Accessing a missing key returns the zero value, NOT undefined
//   - Always use the "comma ok" idiom to distinguish "zero" from "not present"
//   - Maps are reference types — passing to a func mutates the original
// =============================================================================

func basicMaps() {
	// Literal
	scores := map[string]int{
		"alice": 95,
		"bob":   82,
	}

	// Add / update
	scores["charlie"] = 78
	scores["alice"] = 100

	// Delete
	delete(scores, "bob")

	fmt.Println("Scores:", scores)
}

// -----------------------------------------------------------------------------
// The comma-ok idiom — never skip this when the key might be absent
// In JS: obj.missing === undefined
// In Go: scores["missing"] returns 0 (the int zero value) — ambiguous!
// -----------------------------------------------------------------------------
func commaOk() {
	scores := map[string]int{"alice": 95}

	// BAD — you can't tell if alice scored 0 or is absent
	val := scores["bob"]
	fmt.Println("Raw lookup (bob):", val) // prints 0

	// GOOD — use comma-ok
	score, ok := scores["alice"]
	if ok {
		fmt.Println("alice score:", score)
	}

	score, ok = scores["bob"]
	if !ok {
		fmt.Println("bob not found, default:", score) // score=0, ok=false
	}
}

// -----------------------------------------------------------------------------
// Maps are reference types — mutations inside functions affect the original
// Same mental model as JS objects passed to functions
// -----------------------------------------------------------------------------
func addEntry(m map[string]int, key string, value int) {
	m[key] = value // mutates the original map
}

func mapsAreReferences() {
	m := map[string]int{"a": 1}
	addEntry(m, "b", 2)
	fmt.Println("After addEntry:", m) // map[a:1 b:2]
}

// -----------------------------------------------------------------------------
// Iterating — order is NOT guaranteed (intentionally randomised by Go runtime)
// This is different from JS Map which preserves insertion order
// -----------------------------------------------------------------------------
func iterateMap() {
	population := map[string]int{
		"India":  1400,
		"USA":    330,
		"Brazil": 215,
	}

	for country, pop := range population {
		fmt.Printf("  %s: %dM\n", country, pop)
	}
}

func main() {
	fmt.Println("=== Basic Map Operations ===")
	basicMaps()

	fmt.Println("\n=== Comma-Ok Idiom ===")
	commaOk()

	fmt.Println("\n=== Maps Are References ===")
	mapsAreReferences()

	fmt.Println("\n=== Iterating (unordered) ===")
	iterateMap()
}
