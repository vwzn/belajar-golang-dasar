package main

import "fmt"

func getComplete() (firstName , middleName , lastName string) {
	firstName = "Ahmad"
	middleName = "Fauzan"
	lastName = "Abdurrohman"

	return firstName, middleName, lastName
}

func main() {
	a, b, c := getComplete()
	fmt.Println(a, b, c)
}
