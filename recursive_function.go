package main

import "fmt"

func factorialLoop(v int) int {
	r := 1

	for i := v; i > 0; i-- {
		r *= i
	}
	return r

}

func factorialRecursive(v int) int {
	if v == 1 {
		return 1
	} else {
		return v * factorialRecursive(v-1)
	}
}
func main() {
	fmt.Println(factorialLoop(10))
	fmt.Println(factorialRecursive(10))

}
