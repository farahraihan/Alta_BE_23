package main

import (
	"fmt"
)

func drawXYZ(n int) string {
	var result string

	for i := 1; i <= n*n; i += n {
		for j := i; j <= i+n-1; j++ {
			if j%3 == 0 {
				result += "X "
			} else if j%2 == 0 {
				result += "Z "
			} else {
				result += "Y "
			}
		}
		result += "\n"
	}
	return result
}

func main() {
	var num int
	fmt.Print("Masukkan sebuah bilangan : ")
	fmt.Scanln(&num)

	fmt.Println("Huruf X mewakili kelipatan 3")
	fmt.Println("Huruf Z mewakili kelipatan genap")
	fmt.Println("Huruf Y mewakili kelipatan ganjil")
	fmt.Println("_______________________________________________")
	fmt.Println(drawXYZ(num))
}
