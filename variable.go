package main

import "fmt"

func main() {

	// way 1
	var name string

	name = "Fauzan"
	fmt.Println("Nama =", name)

	// way 2
	var fullName = "Ahmad Fauzan Abdurrohman"
	fmt.Println("Nama Lengkap =", fullName)

	// way 3
	email := "zan@gmail.com" //initial declare
	fmt.Println("Email =", email)

	email = "afauzan@gmail.com" //declared at initial declare
	fmt.Println("Second Email =", email)

	// way 4
	var (
		hobby    = "writing code"
		language = "Go Language"
	)
	fmt.Println("My Hobby is", hobby, "and now i learning", language)
}
