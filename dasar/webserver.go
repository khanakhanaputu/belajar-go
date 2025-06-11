package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

type Mobil struct {
	Merk  string `json:"merk"`
	Tipe  string `json:"tipe"`
	Harga int    `json:"harga"`
}

var datamobil = map[string]Mobil{
	"bmw":    {Merk: "BMW", Tipe: "E46 M3", Harga: 300000000},
	"toyota": {Merk: "Toyota", Tipe: "Supra MK4", Harga: 2000000000},
	"honda":  {Merk: "Honda", Tipe: "Civic Type R", Harga: 1400000000},
	"ford":   {Merk: "Ford", Tipe: "Mustang GT 3.0", Harga: 3000000000},
}

func findCarDB(merkMobil string) (Mobil, bool) {
	merkMobil = strings.ToLower(merkMobil)
	mobil, ditemukan := datamobil[merkMobil]
	return mobil, ditemukan
}

func findCar(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	cleanedPath := strings.TrimPrefix(path, "/")
	partPath := strings.Split(cleanedPath, "/")

	if len(partPath) > 1 && partPath[1] != "" {
		car := partPath[1]
		mobil, ketemu := findCarDB(car)
		if ketemu {
			fmt.Fprintf(w, "Kamu mencari %s %s, harganya Rp.%d", mobil.Merk, mobil.Tipe, mobil.Harga)
		} else {
			fmt.Fprint(w, "g ad y cri yg lain y")
		}
	} else {
		fmt.Fprintln(w, "Silahkan ketik di url y syg, contoh: find/merk")
	}

}
func webserver() {
	// http.HandleFunc("/", showHandler)
	http.HandleFunc("/find/", findCar)
	fmt.Println("Server sudah jalan di http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

}
