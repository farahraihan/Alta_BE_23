package main

import (
	"fmt"
)

func main() {
	var bilangan int

	fmt.Print("Masukkan sebuah bilangan : ")
	fmt.Scanln(&bilangan)
	fmt.Printf("Faktor dari %d adalah :\n", bilangan)

	for i := bilangan; i > 0; i-- {
		if bilangan%i == 0 {
			fmt.Println(i)
		}
	}

}
