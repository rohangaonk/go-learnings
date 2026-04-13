package main

import "fmt"

// Talker is our interface. 
// In Node.js, you might use a TypeScript interface or just check if a method exists.
type Talker interface {
	Talk() string
}

type Dog struct {
	Name string
}

// Dog satisfies Talker implicitly. 
// No "implements" keyword needed.
func (d Dog) Talk() string {
	return "Woof!"
}

type Human struct {
	Name string
}

func (h Human) Talk() string {
	return "Hello!"
}

// Shout accepts the interface. It doesn't care WHAT you are, 
// as long as you can Talk().
func Shout(t Talker) {
	fmt.Println(t.Talk())
}

func main() {
	d := Dog{Name: "Buddy"}
	h := Human{Name: "Alice"}

	Shout(d)
	Shout(h)
}
