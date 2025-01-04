package main

import (
	"fmt"
	"reflect"
)

func main() {
	//1. checking type of different types of value
	fmt.Println(reflect.TypeOf(42))
	fmt.Println(reflect.TypeOf(3.1415))
	fmt.Println(reflect.TypeOf(true))
	fmt.Println(reflect.TypeOf("Hello, Go!"))

	//2. variable declaration in one way
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

	//3. variable declarion in one way
	var x = 5
	var l = 1.2
	var w = 2.4
	var c = "david"

	fmt.Println("Name of Custome", c)
	fmt.Println("has ordered", x, "sheets")
	fmt.Println("length :", l, "Breadth :", w)
	fmt.Println("each with an area of")
	fmt.Println("Area ", l*w)

	// 4. variable declarion in another way
	q := 5
	len := 1.2
	wid := 2.4
	cust := "mehak"

	fmt.Println("Name of Custome", cust)
	fmt.Println("has ordered", q, "sheets")
	fmt.Println("length :", len, "Breadth :", wid)
	fmt.Println("each with an area of")
	fmt.Println("Area ", len*wid)

	fmt.Println(reflect.TypeOf(q))
	fmt.Println(reflect.TypeOf(len))
	fmt.Println(reflect.TypeOf(cust))
}
