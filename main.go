package main

import (
	"fmt"
	"log"

	"GO/database"
	"GO/models"
)

func selectAll() {
	database.Connect()
	defer database.DB.Close()

	// jalanin query

	query := "SELECT nama_brg, harga_brg, stok FROM barang"
	rows, err := database.DB.Query(query)

	if err != nil {
		log.Fatalf("gagal query ke database %v", err)
	}
	defer rows.Close()

	var daftarBarang []models.Barang

	for rows.Next() {
		var m models.Barang
		err := rows.Scan(&m.NamaBarang, &m.HargaBarang, &m.Stok)
		if err != nil {
			log.Fatalf("error saat scan rows ada di line 29 %v", err)
		}
		daftarBarang = append(daftarBarang, m)
	}

	err = rows.Err()
	if err != nil {
		log.Fatalf("gagal iterasi pada baris: %v", err)
	}

	fmt.Printf("berhasil mengambil data dari database: %d\n", len(daftarBarang))

	for _, barang := range daftarBarang {
		fmt.Printf("Nama Barang: %s Dengan harga: %d Stok %d \n", barang.NamaBarang, barang.HargaBarang, barang.Stok)
	}
}

func main() {
	selectAll()
}
