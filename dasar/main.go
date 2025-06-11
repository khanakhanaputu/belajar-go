package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	// format penulisan connection postgre postgres://user:password@host:port/dbname
	netStr := "postgres://postgres:Yhabhu@123@localhost:5432/db-toko?sslmode=disable"

	db, err := sql.Open("pgx", netStr)
	if err != nil {
		log.Fatalf("Gagal koneksi ke database %v", err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf("Gagal koneksi ke database %v", err)
	}

	fmt.Print("sukses konek ke database")
}
