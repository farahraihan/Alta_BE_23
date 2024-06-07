package main

import (
	"fmt"
)

func compare(a, b string) string {
	var result string

	for i := -0; i < len(a) && i < len(b); i++ {
		if a[i] == b[i] {
			result += string(a[i])
		} else {
			break
		}
	}

	return result
}

func main() {
	fmt.Println(compare("AKA", "AKASHI"))     //AKA
	fmt.Println(compare("KANGOORO", "KANG"))  //KANG
	fmt.Println(compare("KI", "KIJANG"))      //KI
	fmt.Println(compare("KUPU-KUPU", "KUPU")) //KUPU
	fmt.Println(compare("ILALANG", "ILA"))    //ILA
}
