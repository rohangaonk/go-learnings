package main

import "fmt"

// =============================================================================
// WEEK 2 — Exercises
//
// Cover: slices, maps, structs + methods, control flow
// Same format as week 1 — fill in the TODOs, then run with `go run main.go`
// =============================================================================

// =============================================================================
// EXERCISE 1: Slice manipulation
// Given a slice of ints, return a NEW slice (independent copy) containing only
// the even numbers, doubled.
// e.g. input: [1, 2, 3, 4, 5] → output: [4, 8]
//
// Requirements:
//   - Use range to iterate
//   - The returned slice must NOT share a backing array with the input
// =============================================================================
func doubledEvens(nums []int) []int {
	// your code here
	arr := []int{}

	for _, val := range nums{
		if val % 2 == 0 {
        	arr = append(arr, val * 2)
		}
		
	}
	return arr;
}

// =============================================================================
// EXERCISE 2: Word frequency counter
// Given a slice of strings, return a map[string]int counting how many times
// each word appears.
// e.g. ["go", "is", "go", "great"] → map[go:2 is:1 great:1]
// =============================================================================
func wordFrequency(words []string) map[string]int {
	// your code here
    ans := map[string]int{}
	for _, val := range words {
		ans[val] = ans[val] + 1;
	}

	return ans
}

// =============================================================================
// EXERCISE 3: Stack — implement using a struct + methods
//
// Build a Stack type (backed by a []int slice) with:
//   Push(val int)         — add to top
//   Pop() (int, bool)     — remove from top; return (value, true) or (0, false) if empty
//   Peek() (int, bool)    — same as Pop but without removing
//   IsEmpty() bool
//   Size() int
//
// All mutating methods (Push, Pop) MUST use pointer receivers.
// =============================================================================
type Stack struct {
	// your fields here
	arr []int

}

func (s *Stack) Push(val int) {
	// your code here
	s.arr = append(s.arr, val)
}

func (s *Stack) Pop() (int, bool) {
	ans, exist := s.Peek();
	if exist {
		s.arr = s.arr[:len(s.arr) - 1];
		return ans, exist
	}
	return 0, false
}

func (s *Stack) Peek() (int, bool) {
	if len(s.arr) == 0 {
		return 0, false;
	}
	ans := s.arr[len(s.arr) - 1];
	return ans, true
}

func (s *Stack) IsEmpty() bool {
	// your code here
	if len(s.arr) == 0 {
		return true
	}
	return false
}

func (s *Stack) Size() int {
	return len(s.arr)
}

// =============================================================================
// EXERCISE 4: FizzBuzz — Go style
// Loop from 1 to 30. Use a switch (not if/else) to print:
//   "FizzBuzz" if divisible by 15
//   "Fizz"     if divisible by 3
//   "Buzz"     if divisible by 5
//   the number  otherwise
// =============================================================================
func fizzBuzz() {
	// your code here — use switch, not if/else

	for i:=0; i<=30; i++ {
		switch  {
		case i % 15 == 0:
			fmt.Printf("FizzBuzz ");
		case i % 5 == 0:
			fmt.Printf("Fizz ");
		case i % 3 == 0:
			fmt.Printf("Buzz ");
		default:
			fmt.Printf("%d ", i)
		}
	}
}

// =============================================================================
// EXERCISE 5: Most common word
// Given []string, return the word that appears most frequently.
// If there's a tie, return any of the tied words.
// If the slice is empty, return "".
// =============================================================================
func mostCommonWord(words []string) string {
	// your code here
	m := map[string]int{}
	for _, val := range words {
		m[val] = m[val] + 1;
	}

	count := 0;
	ans := "";
	for key, val := range m {
		if count < val {
			count = val;
			ans = key
		}
	}

	return ans
}

// =============================================================================
// main — do NOT edit; run this to verify your solutions
// =============================================================================
func main() {
	// Exercise 1
	fmt.Println("=== Exercise 1: Doubled Evens ===")
	result := doubledEvens([]int{1, 2, 3, 4, 5})
	fmt.Println("Expected: [4 8]")
	fmt.Println("Got:     ", result)

	// Exercise 2
	fmt.Println("\n=== Exercise 2: Word Frequency ===")
	freq := wordFrequency([]string{"go", "is", "go", "great", "go"})
	fmt.Println("Expected: map[go:3 is:1 great:1]")
	fmt.Println("Got:     ", freq)

	// Exercise 3
	fmt.Println("\n=== Exercise 3: Stack ===")
	s := &Stack{}
	s.Push(10)
	s.Push(20)
	s.Push(30)
	fmt.Println("Size:", s.Size())         // 3
	top, _ := s.Peek()
	fmt.Println("Peek:", top)              // 30
	val, ok := s.Pop()
	fmt.Println("Pop:", val, ok)           // 30 true
	fmt.Println("Size after pop:", s.Size()) // 2
	s.Pop()
	s.Pop()
	_, ok = s.Pop() // pop from empty
	fmt.Println("Pop empty:", ok)          // false
	fmt.Println("IsEmpty:", s.IsEmpty())   // true

	// Exercise 4
	fmt.Println("\n=== Exercise 4: FizzBuzz ===")
	fizzBuzz()

	// Exercise 5
	fmt.Println("\n=== Exercise 5: Most Common Word ===")
	fmt.Println(mostCommonWord([]string{"cat", "dog", "cat", "bird", "dog", "cat"})) // cat
	fmt.Println(mostCommonWord([]string{})) // ""
}
