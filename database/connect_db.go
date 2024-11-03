package database

import (
	"database/sql"
	"fmt"
	"log"
)

func InitDb() (*sql.DB, error) {
	connStr := "user=postgres dbname=Restaurant sslmode=disable password=postgres host=localhost"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Gagal terhubung ke database:", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal("Database tidak merespons:", err)
	}
	fmt.Println("connected")

	return db, err
}
