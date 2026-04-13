package main

import (
	"fmt"
)

// OrderError is a custom error type.
type OrderError struct {
	Code    int
	Message string
}

// Error implements the error interface.
func (e *OrderError) Error() string {
	return fmt.Sprintf("Code %d: %s", e.Code, e.Message)
}

func ValidateOrder(id int) error {
	if id < 0 {
		return &OrderError{
			Code:    400,
			Message: "Invalid Order ID",
		}
	}
	return nil
}

func main() {
	err := ValidateOrder(-5)
	if err != nil {
		fmt.Println("Caught custom error:", err)
		
		// Type assertion: similar to 'if (err instanceof OrderError)' in JS
		if oe, ok := err.(*OrderError); ok {
			fmt.Printf("Extracted Code: %d\n", oe.Code)
		}
	}
}
