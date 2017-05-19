package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

// getConnection obtiene una conexión a la BD
func getConnection() *sql.DB {
	dsn := "postgres://postgres:postgres@127.0.0.1:5432/postgres?sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	return db
}