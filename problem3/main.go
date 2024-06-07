package main

import (
	"fmt"
	"strconv"
)

func cetakTabelPerkalian(n int) {
	strn := strconv.Itoa(n * n)
	digitn := len(strn)

	for i := 1; i <= n; i++ {
		for j := i; j <= n*i; j += i {
			fmt.Print(j)
			// mencetak spasi
			strj := strconv.Itoa(j)
			digitj := len(strj)
			space := (digitn - digitj) + 1
			for k := 0; k < space; k++ {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func main() {
	var num int
	fmt.Print("Masukkan sebuah bilangan : ")
	fmt.Scanln(&num)
	fmt.Print("\n Tabel Perkalian")
	fmt.Println("_______________________________________________")
	cetakTabelPerkalian(num)
}
