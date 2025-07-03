package main

import "fmt"

func main() {
	name := "Fauzan"

	if name == "Fauzan" {
		fmt.Println("Hello Fauzan")
	} else if name == "Arra" {
		fmt.Println("Hello Arra")
	} else {
		fmt.Println("Hai boleh kenalan?")
	}

	length := len(name)
	if length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}

	//if short statment
	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
