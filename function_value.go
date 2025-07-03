package main

import "fmt"

func getGoodBye(name string) string {
	return "Good Bye "  + name 

}

func main() {
	contoh := getGoodBye
	fmt.Println(contoh("Fauzan"))
}
