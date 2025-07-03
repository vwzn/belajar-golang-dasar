package main

import "fmt"

type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	filteredName := filter(name)
	fmt.Println("Hi", filteredName)
}
func spamFilter(name string) string {
	if name == "Anjing" {
		return "***"
	} else {
		return name
	}

}
func main() {
	sayHelloWithFilter("Fauzan", spamFilter)

	filter := spamFilter
	sayHelloWithFilter("Anjing", filter)
}
