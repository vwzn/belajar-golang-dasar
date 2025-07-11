package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total

}
func main() {
	fmt.Println(sumAll(10, 10, 10))
	fmt.Println(sumAll(10, 10, 10, 20, 20, 90, 80, 70, 60))

	numbers := []int{10, 10, 10, 10}
	fmt.Println(sumAll(numbers...))
}
