package main

import (
	"fmt"
	"sync"
)

type LogRequest struct {
	ID      int
	Message string
}

var logPool = sync.Pool{
	New: func() any {
		fmt.Println(">> Allocating new LogRequest")
		return &LogRequest{}
	},
}

func main() {
	// 1. First use
	req1 := logPool.Get().(*LogRequest)
	req1.ID = 101
	req1.Message = "Hello from User A"
	fmt.Printf("User A using: %+v\n", req1)

	// Put it back WITHOUT resetting
	logPool.Put(req1)

	// 2. Second use (highly likely to get the same pointer back)
	req2 := logPool.Get().(*LogRequest)
	
	fmt.Printf("User B got data from pool: %+v\n", req2)

	if req2.Message != "" {
		fmt.Println("!! ALERT: User B received dirty data from User A !!")
	}

	// Proper idiom: Always Zero the struct before putting it back
	req2.ID = 0
	req2.Message = ""
	logPool.Put(req2)
}
