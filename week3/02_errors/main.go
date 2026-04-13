package main

import (
	"errors"
	"fmt"
)

// ProcessOrder should return an error if amount is negative.
func ProcessOrder(amount float64) (string, error) {
	// TODO: Implement error check
	if amount < 0 {
		return "", errors.New("negative amount")
	}
	return "correct", nil
}

func main() {
	// TODO: Call ProcessOrder with a negative value
	_, err := ProcessOrder(-1)
	if err != nil {
		fmt.Print("correct value")
	}else{
		fmt.Print("negative value")
	}
	// TODO: Handle the error using 'if err != nil'

	val, anotherErr := ProcessOrder(1)
	if anotherErr != nil {
		fmt.Print(val)
	}else{
		fmt.Print("incoreect")
	}
	
	// TODO: Call ProcessOrder with a positive value
	// TODO: Print the result
}
