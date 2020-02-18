# Golang SQL App

This simple CLI application demonstrates consistent data modification in a concurrent update scenario using database transactions and locks.

## Instructions

- Ensure you have a postgres database running on port 5433; the easiest way to do this assuming you are running docker is `docker run -p 5433:5432 postgres`
- `go run main.go`

## SQL Compiler

The SQL interop. code was generated using https://github.com/kyleconroy/sqlc