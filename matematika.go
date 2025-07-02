package main

import "fmt"

func main() {
	var a = 10
	var b = 20
	var c = a + b
	var d = a - b + c
	var e = a*b + c - d
	var f = a/b + e - c%a
	var g = a%b - f*e + a - a + c

	fmt.Println(c, d, e, f, g)

	var i = 10
	i += i //i = i + 10
	fmt.Println(i)

	i += 5 //i= i + 5
	fmt.Println(i)

	var j = 1
	j++
	j++
	j++
	j++
	j++
	j++
	j++
	j++
	j--
	j--
	j--
	j--
	j--
	j--
	j--
	j--
	j--
	fmt.Println(j)
}
