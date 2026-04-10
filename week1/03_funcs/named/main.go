package main

import (
	"errors"
	"fmt"
)

// 1. Documentation & Default initialization
// You know exactly what the return values represent just by reading the signature.
func fetchUserSettings(userID int) (theme string, notificationsEnabled bool, err error) {
	// At the start of the function, because they are named returns,
	// Go automatically initializes them to their zero values:
	// theme = ""
	// notificationsEnabled = false
	// err = nil

	if userID < 0 {
		// We can assign directly to the return variables
		err = errors.New("invalid user ID")
		return // "Naked return" picks up the current values of theme, notificationsEnabled, and err
	}

	theme = "dark"
	notificationsEnabled = true
	// Naked return automatically packages up the modified variables
	return 
}

// 2. The Shadowing Gotcha (A very common bug)
func process(val int) (result int, err error) {
	if val > 10 {
		// DANGER: We used ':=' which creates a NEW variable 'result' 
		// inside this 'if' scope. It "shadows" the named return variable.
		result := val * 2 
		res:= 7
		fmt.Printf("Inner result: %d\n", result)
		
		// If we do a naked 'return' here, Go will throw a compile error: 
		// "result is shadowed during return"
		// To fix it, we have to either use '=' instead of ':=', 
		// or explicitly return: 'return result, err'
		
		// We explicitly return the inner shadowed result to satisfy the compiler
		return res, nil 
	}

	result = val
	return
}

func main() {
	t, n, e := fetchUserSettings(-1)
	fmt.Printf("Theme: %q, Notifications: %v, Error: %v\n", t, n, e)

	t, n, e = fetchUserSettings(42)
	fmt.Printf("Theme: %q, Notifications: %v, Error: %v\n", t, n, e)
    
	res, err := process(15)
	fmt.Printf("Process Result: %d, Error: %v\n", res, err)
}
