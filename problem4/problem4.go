package main

import (
	"fmt"
)

type student struct {
	name       string
	nameEncode string
}

type Chiper interface {
	Encode() string
	Decode() string
}

func (s *student) Encode() string {
	var nameEncode = ""
	for _, char := range s.name {

		if char >= 'a' && char <= 'z' {
			nameEncode += string((char-'a'+3)%26 + 'a')
		} else if char >= 'A' && char <= 'Z' {
			nameEncode += string((char-'A'+3)%26 + 'A')
		} else {
			nameEncode += string(char)
		}
	}
	s.nameEncode = nameEncode
	return nameEncode
}

func (s *student) Decode() string {
	var nameDecode = ""
	for _, char := range s.nameEncode {
		// Check if character is a letter
		if char >= 'a' && char <= 'z' {
			nameDecode += string((char-'a'-3+26)%26 + 'a')
		} else if char >= 'A' && char <= 'Z' {
			nameDecode += string((char-'A'-3+26)%26 + 'A')
		} else {
			nameDecode += string(char)
		}
	}
	return nameDecode
}

func main() {
	var menu int
	var a = student{}
	var c Chiper = &a
	fmt.Print("[1] Encrypt \n[2] Decrypt \nChoose your menu ? ")
	fmt.Scan(&menu)
	if menu == 1 {
		fmt.Print("\nInput Studet's Name : ")
		fmt.Scan(&a.name)
		fmt.Print("\nEncode pf Student's Name " + a.name + " is : " + c.Encode())
	} else if menu == 2 {
		fmt.Print("\nInput Student's Encode Name : ")
		fmt.Scan(&a.nameEncode)
		fmt.Print("\nDecode of Students's Name " + a.nameEncode + " is : " + c.Decode())
	} else {
		fmt.Println("Wrong input name menu")
	}
}
