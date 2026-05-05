-- Run this against your local Postgres instance before starting the exercise.
-- Example: psql -U postgres -d go_learning -f schema.sql
--
-- Create DB first if needed:
--   createdb go_learning
--   OR: psql -U postgres -c "CREATE DATABASE go_learning;"

CREATE TABLE IF NOT EXISTS users (
    id   SERIAL PRIMARY KEY,
    name TEXT        NOT NULL,
    age  INTEGER     NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
