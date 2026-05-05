package main

import "fmt"

// Week 9.2 - migrations
// Node.js bridge:
//   This is like keeping Prisma/Knex migrations explicit, versioned, and replayable.
//   Go teams usually run SQL migrations in CI and on deploy startup hooks.
//
// Assessment (tricky snippet): what's wrong here?
//
//   upSQL := "CREATE TABLE users (id SERIAL PRIMARY KEY)"
//   downSQL := "DROP TABLE users"
//   // You shipped only upSQL and forgot downSQL registration in your migration tool.
//
// Explain the operational risk and when this can break rollback strategy.
// Suppose you want to rollback a change which also involves rolling back the migration to bring
// Schema to its original structure then if you dont have down migration then you will have problems.

func main() {
	fmt.Println("week9-02: migrations scaffold ready")
	fmt.Println("next: add 000001_create_users.up.sql and .down.sql")
}

