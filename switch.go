package main

import "fmt"

func main() {
	name := "Fauzan"

	switch name {
	case "Fauzan":
		fmt.Println("Hallo Fauzan")
	case "Arra":
		fmt.Println("Hallo Arra")
	default:
		fmt.Println("Siapa kamu ?")
	}

	//switch short statement
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama kepanjangan")
	case false:
		fmt.Println("Nama sudah benar")
	}

	//switch tanpa kondisi
	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah benar")
	}
}
