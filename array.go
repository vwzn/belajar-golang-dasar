package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Ahmad"
	names[1] = "Fauzan"
	names[2] = "Abdurrohman"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var v =[...]int {
		90, 80, 85,
	}
	fmt.Println(v)
	fmt.Println(v[0])
	fmt.Println(v[1])
	fmt.Println(v[2])
	fmt.Println(len(v))
	v[0]= 199
	fmt.Println(v)
}