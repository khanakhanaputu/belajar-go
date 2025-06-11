package main

import (
	"fmt"
)

type Mahasiswa struct {
	Nama  string
	Umur  int
	Kelas string
	NIM   int
	Aktif bool
}

func masins() {

	// slice
	daftarMahasiswa := []Mahasiswa{}

	mahasiswaPertama := Mahasiswa{
		Nama:  "Putu Khana Rahayu Putra",
		Umur:  19,
		Kelas: "D",
		NIM:   24010193,
		Aktif: true,
	}

	mahasiswaKedua := Mahasiswa{
		Nama:  "Putu Ayam Goreng Telur",
		Umur:  19,
		Kelas: "D",
		NIM:   24010192,
		Aktif: false,
	}

	daftarMahasiswa = append(daftarMahasiswa, mahasiswaPertama)
	daftarMahasiswa = append(daftarMahasiswa, mahasiswaKedua)

	// map di go mirip kek array assosiatif di php
	infoSingkat := make(map[string]bool)

	// infoSingkat["Putu Ayam Goreng Telur"] = False
	infoSingkat[mahasiswaKedua.Nama] = mahasiswaKedua.Aktif

	ok := infoSingkat[mahasiswaKedua.Nama]
	if ok {
		fmt.Printf("%s Masih aktif dalam kemahasiswaan", mahasiswaKedua.Nama)
	} else {
		fmt.Printf("%s Sudah DO atau membeli susu", mahasiswaKedua.Nama)
	}

}
