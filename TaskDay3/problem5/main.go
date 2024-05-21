package main

import "fmt"

func main() {
    var r, T float64

    fmt.Print("Masukkan jari-jari tabung (r): ")
    fmt.Scanln(&r)

    fmt.Print("Masukkan tinggi tabung (T): ")
    fmt.Scanln(&T)

    pi := 3.141592653589793
    luasPermukaan := 2 * pi * r * (r + T)

    fmt.Printf("Luas permukaan tabung dengan jari-jari %.2f dan tinggi %.2f adalah %.2f\n", r, T, luasPermukaan)
}
