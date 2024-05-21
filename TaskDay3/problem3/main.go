package main

import "fmt"

func main() {
	var nama string
    
	fmt.Print("Masukkan sebuah nama : ")
	fmt.Scanln(&nama)
    fmt.Printf("Hello %s ! Saya Golang, bahasa yang sangat menyenangkan.", nama)

}
