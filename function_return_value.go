package main

import "fmt"

func getHello(name string) string {
	hello := "Hello" +" "+ name
	return hello
}
func main() {
	result := getHello("Fauzan")
	fmt.Println(result)

	fmt.Println(getHello("Arra"))
	fmt.Println(getHello("Shya"))
}
