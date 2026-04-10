package main

import "fmt"

// =============================================================================
// EXERCISE 1: Swap two integers using pointers
// In JavaScript you can't truly swap primitives via a function.
// In Go, you can — do it here.
// =============================================================================
// TODO: Fill in the function body.
// swap should modify a and b IN-PLACE so that after calling swap(&x, &y),
// x and y are swapped without returning anything.
func swap(a *int, b *int) {
	// your code here
	val := *a ;
	*a = *b;
	*b = val;
}

// =============================================================================
// EXERCISE 2: Mutate a struct via a pointer
// A JS object passed to a function is automatically a reference.
// In Go, structs are copied. Use a pointer to mutate the original.
// =============================================================================
type User struct {
	Name string
	Age  int
}

// TODO: Fill in the function body.
// birthday should increment the user's Age by 1, modifying the original struct.
func birthday(u *User) {
	// your code here
	u.Age += 1;
}

// =============================================================================
// EXERCISE 3: Nil Guard — safely dereference or return a default
// A very common pattern in production Go code.
// =============================================================================
// TODO: Fill in the function body.
// safeRead should return the value at ptr if ptr is not nil.
// If ptr IS nil, it should return 0 (the default).
// Hint: use an if statement to check for nil before dereferencing.
func safeRead(ptr *int) int {
	// your code here
	if(ptr != nil){
		return *ptr
	}
	return 0
}

// =============================================================================
// EXERCISE 4: Zero-Value Detective
// Predict the output of this function BEFORE running it.
// Then run the file to check if you were right.
// =============================================================================
func zeroValueDetective() {
	var count int
	var price float64
	var label string
	var active bool
	var ptr *int

	fmt.Printf("count=%d price=%f label=%q active=%v ptr=%v\n",
		count, price, label, active, ptr)
	//will print 0, 0.0, "", false, address
}

// =============================================================================
// EXERCISE 5: Refactor from JS-style to idiomatic Go
// This function returns a map to simulate a JS object — very unidiomatic.
// Refactor it to return a proper Go struct with named return values.
// =============================================================================
func getUserInfo_jsStyle() map[string]interface{} {
	return map[string]interface{}{
		"name":  "Rohan",
		"age":   28,
		"admin": true,
	}
}

// TODO: Create a struct called UserInfo and a new function getUserInfo() that
// returns (UserInfo, error) using named returns. The function should return
// a valid UserInfo with Name="Rohan", Age=28, Admin=true.

type UserInfo struct {
	name string
	age int
	admin bool
}

func getUserInfo() (info UserInfo, err error){
	info.name = "Rohan";
	info.age = 28;
	info.admin = true;
	return 
}


func main() {
	// --- Exercise 1 ---
	x, y := 10, 20
	fmt.Printf("Before swap: x=%d y=%d\n", x, y)
	swap(&x, &y)
	fmt.Printf("After swap:  x=%d y=%d\n\n", x, y) // Expected: x=20 y=10

	// --- Exercise 2 ---
	u := User{Name: "Rohan", Age: 27}
	fmt.Printf("Before birthday: %v\n", u)
	birthday(&u)
	fmt.Printf("After birthday:  %v\n\n", u) // Expected Age: 28

	// --- Exercise 3 ---
	val := 42
	fmt.Println("safeRead(&val):", safeRead(&val))     // Expected: 42
	fmt.Println("safeRead(nil):", safeRead(nil))        // Expected: 0
	fmt.Println()

	// --- Exercise 4 ---
	fmt.Print("Zero Detective: ")
	zeroValueDetective()

	// --- Exercise 5 ---
	// Uncomment once you've written getUserInfo()
	info, err := getUserInfo()
	fmt.Printf("UserInfo: %+v, Err: %v\n", info, err)
}
