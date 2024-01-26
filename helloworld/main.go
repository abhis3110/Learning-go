package main

import "fmt"

const s string = "constant"

func main() {
	fmt.Println("Hi there!")

	// values in golang
	fmt.Println("go" + "lang")
	fmt.Println("1+1", 1+1)
	fmt.Println("7.0/3.0", 7.0/3.0)
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

	// Variables in golang
	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)
	//fmt.Println(c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "apple"
	fmt.Println(f)

	// constants
	fmt.Println(s)

	const n = 500000000
	const g = 3e20 / n
	fmt.Println(g)
	fmt.Println(int64(g))

	// arrays
	var arr [5]int

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

}
