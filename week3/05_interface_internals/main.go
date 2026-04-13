package main

import (
	"fmt"
)

type MyError struct{}

func (e *MyError) Error() string { return "mine" }

func returnsNilConcrete() *MyError {
	return nil
}

func returnsNilInterface() error {
	return nil
}

func main() {
	// 1. Returns nil interface (Type: nil, Value: nil)
	err1 := returnsNilInterface()
	fmt.Printf("err1: Type=%T, Val=%v, IsNil?=%v\n", err1, err1, err1 == nil)

	// 2. Returns nil concrete pointer (Type: *MyError, Value: nil)
	err2 := error(returnsNilConcrete()) 
	fmt.Printf("err2: Type=%T, Val=%v, IsNil?=%v\n", err2, err2, err2 == nil)

	if err2 != nil {
		fmt.Println("CRITICAL GOTCHA: err2 is NOT nil because it has a Type!")
	}
}
