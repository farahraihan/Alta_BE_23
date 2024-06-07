package main

import (
	"fmt"
)

func main() {
	var x, n int

	fmt.Print("Masukkan nilai x : ")
	fmt.Scanln(&x)
	fmt.Print("Masukkan nilai n : ")
	fmt.Scanln(&n)

	result := 1
	for i := 0; i < n; i++ {
		result *= x
	}

	fmt.Printf("%d pangkat %d adalah %d\n", x, n, result)
}
