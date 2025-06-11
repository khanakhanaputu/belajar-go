package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

// bikin Public variabel
var DB *sql.DB

func Connect() {
	// bikin syntax query ke database
	connStr := "postgres://postgres:Yhabhu@123@localhost:5432/db-toko?sslmode=disable"

	// connect ke database dan cek error
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatalf("gagal koneksi ke database %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("gagal koneksi ke database %v", err)
	}

	// masukkan kedalam Public variable
	DB = db

	fmt.Print("sukses konek database")
}
