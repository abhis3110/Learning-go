package main

import (
	"fmt"
	"reflect"
)

func main() {
	// checking type of different types of value
	fmt.Println(reflect.TypeOf(42))
	fmt.Println(reflect.TypeOf(3.1415))
	fmt.Println(reflect.TypeOf(true))
	fmt.Println(reflect.TypeOf("Hello, Go!"))

	// how to use fmt
	var quantity int
	var length, width float64
	var customerName string

	quantity = 4
	length, width = 1.2, 2.4
	customerName = "damon Cole"

	fmt.Println("Name of Custome", customerName)
	fmt.Println("has ordered", quantity, "sheets")
	fmt.Println("length :", length, "Breadth :", width)
	fmt.Println("each with an area of")
	fmt.Println("Area ", length*width)
}
