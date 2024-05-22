package main

import (
	"fmt"
	"strconv"
)

func isPrimeDigit(digit int) bool {
	switch digit {
	case 2, 3, 5, 7:
		return true
	default:
		return false
	}
}

func isPrime(num int) bool {
	if num <= 1 {
		return false
	} else if num == 2 {
		return true
	} else if num%2 == 0 {
		return false
	} else {
		isPrime := true
		for i := 3; i < num; i++ {
			if num%i == 0 {
				return false
			}
		}
		if isPrime {
			return true
		} else {
			return false
		}

	}
}

func main() {
	var num int

	fmt.Print("Masukkan sebuah bilangan: ")
	fmt.Scanln(&num)

	checkPrima := isPrime(num)
	isFullPrima := true

	if checkPrima {
		numStr := strconv.Itoa(num)
		for i := 0; i < len(numStr); i++ {
			digitChar := numStr[i]
			digitStr := string(digitChar)
			digitInt, err := strconv.Atoi(digitStr)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			digitPrima := isPrimeDigit(digitInt)
			if !digitPrima {
				isFullPrima = false
				break
			}
		}
		if isFullPrima {
			fmt.Println("True (Bilangan full prima)")
		} else {
			fmt.Println("False (Bukan bilangan full prima)")
		}

	} else {
		fmt.Println("False (Bukan bilangan full prima)")
	}

}
