package main

import "fmt"

// =============================================================================
// WEEK 2 — Lesson 3: Structs & Methods
//
// Go has NO classes. Instead:
//   struct  = data
//   method  = function attached to a type (receiver syntax)
//
// This is not a limitation — it's intentional. Go favours composition over
// inheritance (same philosophy as JS prototypes, but more explicit).
// =============================================================================

// -----------------------------------------------------------------------------
// Structs — typed, named collections of fields
// JS equivalent: a class with only properties, or a typed plain object
// -----------------------------------------------------------------------------
type Rectangle struct {
	Width  float64
	Height float64
}

// -----------------------------------------------------------------------------
// Methods — functions with a "receiver" (the thing before the function name)
//
// Value receiver  (r Rectangle)  — works on a COPY; safe for reads
// Pointer receiver (*Rectangle)  — works on the original; use for mutations
//                                  or when the struct is large
//
// Rule of thumb: if any method on a type uses a pointer receiver,
// all methods on that type should use pointer receivers (consistency).
// -----------------------------------------------------------------------------

// Value receiver — r is a copy; fine for read-only
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Pointer receiver — modifies the original Rectangle
func (r *Rectangle) Scale(factor float64) {
	r.Width *= factor
	r.Height *= factor
}

// -----------------------------------------------------------------------------
// Embedding — Go's answer to inheritance
// Embed one struct inside another to "inherit" fields and methods.
// There's no true inheritance — it's just promotion of fields/methods.
// -----------------------------------------------------------------------------
type ColoredRectangle struct {
	Rectangle        // embedded (promoted) — no field name
	Color     string
}

// -----------------------------------------------------------------------------
// Constructor pattern — Go has no `new` keyword for structs.
// Return a pointer from a function named New<Type>.
// -----------------------------------------------------------------------------
func NewRectangle(w, h float64) *Rectangle {
	return &Rectangle{Width: w, Height: h}
}

func main() {
	// Struct literal — two styles
	r1 := Rectangle{Width: 10, Height: 5}
	r2 := Rectangle{10, 5} // positional (less readable, avoid for >2 fields)
	_ = r2

	// Calling methods
	fmt.Printf("Area: %.2f\n", r1.Area())
	fmt.Printf("Perimeter: %.2f\n", r1.Perimeter())

	// Go auto-dereferences — you don't need (*r1).Scale(...)
	r1.Scale(2)
	fmt.Printf("After Scale(2): %+v\n\n", r1)

	// Embedding: ColoredRectangle has Area(), Perimeter(), Scale() promoted from Rectangle
	cr := ColoredRectangle{
		Rectangle: Rectangle{Width: 3, Height: 4},
		Color:     "red",
	}
	fmt.Printf("ColoredRect Area: %.2f Color: %s\n\n", cr.Area(), cr.Color)

	// Constructor pattern
	r3 := NewRectangle(7, 3)
	fmt.Printf("NewRectangle: %+v  Area: %.2f\n", *r3, r3.Area())
}
