package main

import "fmt"

func main() {
	var hargaAwal, diskon float64

	// Menerima input harga awal dan diskon dari pengguna
	fmt.Print("Masukkan harga awal barang: ")
	fmt.Scanln(&hargaAwal)

	fmt.Print("Masukkan persentase diskon (dalam persen): ")
	fmt.Scanln(&diskon)

	// Menghitung harga akhir setelah diskon
	hargaAkhir := hargaAwal * (1 - diskon/100)

	// Mencetak hasil
	fmt.Printf("Harga akhir setelah diskon adalah: %.2f\n", hargaAkhir)
}
