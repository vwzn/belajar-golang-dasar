package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("Perulangan ke", counter)
		counter++
	}
	fmt.Println("Selesai")

	//for dengan statement
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke", counter)
	}
	fmt.Println("Selesai")

	names := []string{"Ahmad", "Fauzan", "Arra"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	for index, name := range names {
		fmt.Println("Index", index, "=", name)
	}

	for _, name := range names {
		fmt.Println( name)
	}
}
