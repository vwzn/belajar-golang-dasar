package main

import "fmt"

func main() {
	const (
		firstName string = "Ahmad"
		lastName         = "Abdurrohman"
	)

	fmt.Println(firstName, lastName)
	
	// Err
	// firstName = "Budi"
	// lastName = "Joko"
}
