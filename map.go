package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Fauzan",
		"country": "Indonesia",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["country"])

	book := map[string]string{
		"author": "Fauzan",
		"title":  "buku Golang",
		"wrong":  "ups",
	}

	fmt.Println(book)
	fmt.Println(book["author"])
	fmt.Println(book["title"])
	delete(book, "wrong")
	fmt.Println(book)
}
