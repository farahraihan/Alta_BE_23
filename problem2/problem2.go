package main

import (
	"fmt"
)

func caesar(offset int, input string) string {
	var result string
	var ascii int
	var char rune

	for i := 0; i < len(input); i++ {
		if offset > 122 {
			offset = offset % 26
		}
		ascii = int(input[i]) + offset
		if ascii > 122 {
			ascii = ascii - 122
			ascii = 96 + ascii
			char = rune(ascii)
			result += string(char)
		} else {
			char = rune(ascii)
			result += string(char)
		}
	}
	return result
}

func main() {
	fmt.Println(caesar(3, "abc"))             //def
	fmt.Println(caesar(2, "alta"))            //cnvc
	fmt.Println(caesar(10, "alterraacademy")) //kvdobbkkmknowi
	fmt.Println(caesar(1, "abcdefghijklmnopqrstuvwxyz"))
	fmt.Println(caesar(1000, "abcdefghijklmnopqrstuvwxyz"))
}
