-- name: CreateUser :one
INSERT INTO users (name, age)
VALUES ($1, $2)
RETURNING id, name, age, created_at;

-- name: GetUser :one
SELECT id, name, age, created_at
FROM users
WHERE id = $1;

-- name: ListUsers :many
SELECT id, name, age, created_at
FROM users
ORDER BY id;

-- name: UpdateUser :one
UPDATE users
SET name = $2, age = $3
WHERE id = $1
RETURNING id, name, age, created_at;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;

-- name: GetUsersByMinAge :many
SELECT * FROM users
WHERE age >= $1
ORDER BY age ASC;