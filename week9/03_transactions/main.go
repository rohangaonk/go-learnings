package main

import "fmt"

// Week 9.3 - transactions
// Node.js bridge:
//   Equivalent to using a pg client transaction block where all writes succeed/fail together.
//   In Go, you pass context.Context and transaction handle explicitly.
//
// Assessment (tradeoff prompt):
// When would you use a single transaction for "create user + write audit log"
// versus an outbox/event-driven approach?
// Answer with latency, consistency, and failure-retry tradeoffs.
// Latency: if you need low latency outbox pattern else you can go for single transaction for low throughput use cases
// Consistency: keeping both statements in txn will gurantee consistency because if outbox failed to create log user creation will not be reverted
// Failure retry: with txn either both passes or fail you can safely retry the operation if any one fails.

func main() {
	fmt.Println("week9-03: transactions scaffold ready")
	fmt.Println("next: implement BeginTx, rollback on error, commit on success")
}
