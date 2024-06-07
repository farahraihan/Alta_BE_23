package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var s string

	fmt.Print("Masukkan sebuah kata atau kalimat: ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		s = scanner.Text()
	}

	s = strings.ToLower(strings.ReplaceAll(s, " ", ""))
	length := len(s)
	isPalindrome := true

	for i := 0; i < length/2; i++ {
		if s[i] != s[length-1-i] {
			isPalindrome = false
			break
		}
	}

	if isPalindrome {
		fmt.Println("true (palindrome)")
	} else {
		fmt.Println("false (bukan palindrome)")
	}
}
