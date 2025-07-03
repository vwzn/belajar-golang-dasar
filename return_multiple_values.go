package main

import "fmt"

func getFullname() (string, string) {
	return "Ahmad", "Fauzan"

}
func main() {
	// firstName, lastName := getFullname()
	// fmt.Println(firstName,lastName)

	firstName, _ := getFullname()
	fmt.Println(firstName)
}
