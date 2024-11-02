package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	var b = [2]string{"Hello", "Coding"}
	fmt.Println(b[0], b[1])


	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}