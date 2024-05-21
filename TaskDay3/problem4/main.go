package main

import "fmt"

func main() {
    var alas float64 = 20
    var tinggi float64 = 25

    luas := 0.5 * alas * tinggi

    fmt.Printf("Luas segitiga dengan alas %.2f dan tinggi %.2f adalah %.2f\n", alas, tinggi, luas)
}
