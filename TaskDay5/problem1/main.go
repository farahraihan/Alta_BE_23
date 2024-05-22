package main

import (
	"fmt"
)

func playWithAsterik(n int) {
	for i := 1; i <= n; i++ {
		for j := 0; j < n-i; j++ {
			fmt.Print(" ")
		}
		for k := 0; k < i; k++ {
			fmt.Printf("* ")
		}
		fmt.Println()
	}
}

func main() {
	var num int
	fmt.Print("Masukkan sebuah bilangan : ")
	fmt.Scanln(&num)
	fmt.Println("_______________________________________________")
	playWithAsterik(num)
}
