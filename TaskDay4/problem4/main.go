package main

import (
	"fmt"
)

func main() {
	var n int

	fmt.Print("Masukkan sebuah bilangan : ")
	fmt.Scanln(&n)

	if n <= 1 {
		fmt.Println("false (bukan prima)")
	} else if n == 2 {
		fmt.Println("true (bilangan prima)")
	} else if n%2 == 0 {
		fmt.Println("false (bukan prima)")
	} else {
		isPrime := true
		for i := 3; i < n; i++ {
			if n%i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			fmt.Println("true (bilangan prima)")
		} else {
			fmt.Println("false (bukan prima)")
		}

	}

}
