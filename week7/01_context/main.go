package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

// ---------------------------------------------------------------------------
// 1. context.WithCancel — manual cancellation
// ---------------------------------------------------------------------------

// worker listens on ctx.Done() to know when to stop.
// This is the fundamental pattern: every blocking operation should respect ctx.
func worker(ctx context.Context, id int) {
	for {
		select {
		case <-ctx.Done():
			// ctx.Err() tells you WHY it was cancelled: Canceled or DeadlineExceeded
			fmt.Printf("worker %d stopping: %v\n", id, ctx.Err())
			return
		default:
			fmt.Printf("worker %d doing work\n", id)
			time.Sleep(200 * time.Millisecond)
		}
	}
}

func demoCancel() {
	fmt.Println("=== WithCancel ===")
	// context.Background() is the root context — use it at the top of main or
	// at the entry point of an HTTP handler / goroutine tree.
	ctx, cancel := context.WithCancel(context.Background())

	go worker(ctx, 1)
	go worker(ctx, 2)

	time.Sleep(500 * time.Millisecond)

	// cancel() signals all children that they should stop.
	// ALWAYS call cancel to release resources — even if the context times out.
	cancel()

	time.Sleep(50 * time.Millisecond) // let goroutines print their exit message
	fmt.Println()
}

// ---------------------------------------------------------------------------
// 2. context.WithTimeout — automatic cancellation after a duration
// ---------------------------------------------------------------------------

// fetchData simulates a slow I/O call (e.g. DB query, HTTP request).
// It honours the deadline embedded in ctx.
func fetchData(ctx context.Context) (string, error) {
	// Simulate work that takes 300ms
	select {
	case <-time.After(300 * time.Millisecond):
		return "result", nil
	case <-ctx.Done():
		// Wrap the context error so callers can use errors.Is for matching.
		return "", fmt.Errorf("fetchData: %w", ctx.Err())
	}
}

func demoTimeout() {
	fmt.Println("=== WithTimeout ===")

	// Case 1: deadline is generous enough
	ctx1, cancel1 := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel1()
	result, err := fetchData(ctx1)
	if err != nil {
		fmt.Println("case 1 error:", err)
	} else {
		fmt.Println("case 1 result:", result)
	}

	// Case 2: deadline is too tight — fetchData will be cut off
	ctx2, cancel2 := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel2()
	result, err = fetchData(ctx2)
	if err != nil {
		fmt.Println("case 2 error:", err)
		// Use errors.Is to distinguish cancellation reasons
		if errors.Is(err, context.DeadlineExceeded) {
			fmt.Println("  -> deadline exceeded (not a bug, retry or return 504)")
		}
	} else {
		fmt.Println("case 2 result:", result)
	}
	fmt.Println()
}

// ---------------------------------------------------------------------------
// 3. context.WithDeadline — same as Timeout but with an absolute time
// ---------------------------------------------------------------------------

func demoDeadline() {
	fmt.Println("=== WithDeadline ===")
	deadline := time.Now().Add(150 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	// Deadline() returns the absolute time and whether one was set.
	if d, ok := ctx.Deadline(); ok {
		fmt.Printf("deadline set: %v (in %v)\n", d.Format(time.RFC3339), time.Until(d).Round(time.Millisecond))
	}

	_, err := fetchData(ctx) // fetchData takes 300ms, deadline is 150ms
	fmt.Println("deadline result:", err)
	fmt.Println()
}

// ---------------------------------------------------------------------------
// 4. context.WithValue — passing request-scoped values down the call chain
// ---------------------------------------------------------------------------

// Use an unexported key type to avoid collisions with other packages.
// NEVER use a built-in type (string, int) as a context key.
type contextKey string

const requestIDKey contextKey = "requestID"

func processRequest(ctx context.Context) {
	// Retrieve the value; type-assert it.
	reqID, ok := ctx.Value(requestIDKey).(string)
	if !ok {
		fmt.Println("no request ID in context")
		return
	}
	fmt.Printf("processing request ID: %s\n", reqID)
	logRequest(ctx)
}

func logRequest(ctx context.Context) {
	// Context values propagate to all children automatically — no need to
	// pass reqID as a separate argument.
	reqID := ctx.Value(requestIDKey).(string)
	fmt.Printf("logging request ID: %s\n", reqID)
}

func demoValue() {
	fmt.Println("=== WithValue ===")
	ctx := context.WithValue(context.Background(), requestIDKey, "req-abc-123")
	processRequest(ctx)
	fmt.Println()
}

// ---------------------------------------------------------------------------
// 5. Propagating context through a call chain — the real-world pattern
// ---------------------------------------------------------------------------

// In production every function that does I/O takes ctx as its FIRST argument.
// Here we simulate: HTTP handler -> service layer -> repository layer.

func repository(ctx context.Context, userID int) (string, error) {
	// Simulate DB query respecting ctx deadline/cancellation
	select {
	case <-time.After(50 * time.Millisecond):
		return fmt.Sprintf("user-%d-data", userID), nil
	case <-ctx.Done():
		return "", fmt.Errorf("repository: %w", ctx.Err())
	}
}

func service(ctx context.Context, userID int) (string, error) {
	data, err := repository(ctx, userID)
	if err != nil {
		return "", fmt.Errorf("service: %w", err)
	}
	return "processed:" + data, nil
}

// httpHandler simulates an HTTP handler that receives a request with a timeout.
func httpHandler(userID int) {
	// In net/http, the framework gives you r.Context() which is already
	// cancelled when the client disconnects. Here we simulate that with
	// a 200ms timeout.
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel() // critical: prevents goroutine/timer leak

	result, err := service(ctx, userID)
	if err != nil {
		fmt.Println("handler error:", err)
		return
	}
	fmt.Println("handler response:", result)
}

func demoCallChain() {
	fmt.Println("=== Call Chain Propagation ===")
	httpHandler(42)
	fmt.Println()
}

// ---------------------------------------------------------------------------
// Main
// ---------------------------------------------------------------------------

func main() {
	demoCancel()
	demoTimeout()
	demoDeadline()
	demoValue()
	demoCallChain()
}
