package main

import (
	"fmt"
	"strconv"
)

func isPrimeDigit(digit rune) bool {
	switch digit {
	case '2', '3', '5', '7':
		return true
	default:
		return false
	}
}

func main() {
	var num int

	fmt.Print("Masukkan sebuah bilangan: ")
	fmt.Scanln(&num)

	numStr := strconv.Itoa(num)
	isFullPrime := true

	for _, digit := range numStr {
		if !isPrimeDigit(digit) {
			isFullPrime = false
			break
		}
	}

	if isFullPrime {
		fmt.Println("Bilangan tersebut adalah bilangan full prima.")
	} else {
		fmt.Println("Bilangan tersebut bukan bilangan full prima.")
	}
}
