package main

import "fmt"

func shiftLetter(char rune) rune {
	// Tentukan batas atas dari rentang huruf besar dan huruf kecil
	var base rune
	if char >= 'a' && char <= 'z' {
		base = 'a'
	} else if char >= 'A' && char <= 'Z' {
		base = 'A'
	} else {
		// Jika bukan huruf, kembalikan karakter aslinya tanpa menggeser
		return char
	}

	// Geser nilai ASCII
	shifted := (char-base+10)%26 + base
	return shifted
}

func shiftString(sentence string) string {
	var result string

	for _, char := range sentence {
		// Geser setiap huruf dalam string
		shifted := shiftLetter(char)
		result += string(shifted)
	}

	return result
}

func main() {
	fmt.Println(shiftString("SEPULSA OKE"))
	fmt.Println(shiftString("ALTERRA ACADEMY"))
	fmt.Println(shiftString("INDONESIA"))
	fmt.Println(shiftString("GOLANG"))
	fmt.Println(shiftString("PROGRAMMER"))
}
