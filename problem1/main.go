package main

import (
	"fmt"
)

func main() {
	var nama string
	var nilai float32
	fmt.Print("Masukkan nama mahasiswa : ")
	fmt.Scanln(&nama)
	fmt.Print("Masukkan nilai : ")
	fmt.Scanln(&nilai)

	var grade string
	switch {
	case nilai >= 80 && nilai <= 100:
		grade = "A"
	case nilai >= 65 && nilai <= 79:
		grade = "B"
	case nilai >= 50 && nilai <= 64:
		grade = "C"
	case nilai >= 35 && nilai <= 49:
		grade = "D"
	case nilai >= 0 && nilai <= 34:
		grade = "E"
	default:
		grade = "Invalid input"
	}

	// Menampilkan hasil
	fmt.Printf("Nilai %s: %2.f, Grade: %s\n", nama, nilai, grade)
}
