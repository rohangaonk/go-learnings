package main

// Week 9.1 — sqlc
//
// sqlc reads schema.sql + query.sql and generates:
//   db/models.go      — User struct (typed, json-tagged)
//   db/query.sql.go   — CreateUser, GetUser, ListUsers, UpdateUser, DeleteUser
//   db/db.go          — Queries struct + DBTX interface
//
// Node.js bridge:
//   db.New(pool)              ≈  knex({ client: 'pg', connection: pool })
//   q.CreateUser(ctx, params) ≈  knex('users').insert(params).returning('*')
//   No .Scan() anywhere — sqlc writes it for you, statically verified at generate time.
//
// Run:
//   export DATABASE_URL="postgres://localhost:5432/go_learning?sslmode=disable"
//   go run .

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"

	"go-learning/week9/01_sqlc/db"
)

func main() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "postgres://localhost:5432/go_learning?sslmode=disable"
	}

	ctx := context.Background()

	// pgxpool satisfies sqlc's DBTX interface directly.
	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		log.Fatalf("pool: %v", err)
	}
	defer pool.Close()

	q := db.New(pool) // db.New accepts any DBTX — pool, conn, or tx

	// Create
	user, err := q.CreateUser(ctx, db.CreateUserParams{Name: "Rohan", Age: 30})
	if err != nil {
		log.Fatalf("create: %v", err)
	}
	fmt.Printf("created: %+v\n", user)

	// Get
	fetched, err := q.GetUser(ctx, user.ID)
	if err != nil {
		log.Fatalf("get: %v", err)
	}
	fmt.Printf("fetched: %+v\n", fetched)

	// List
	users, err := q.ListUsers(ctx)
	if err != nil {
		log.Fatalf("list: %v", err)
	}
	fmt.Printf("total users: %d\n", len(users))

	// Delete
	if err := q.DeleteUser(ctx, user.ID); err != nil {
		log.Fatalf("delete: %v", err)
	}
	fmt.Println("deleted ok")
}