package main

import (
	"fmt"
)

func main() {

	var whatToSay string // declare
	var i int

	whatToSay = "Goodbye, cruel world" // initialize

	fmt.Println(whatToSay)

	i = 7 // initialize

	fmt.Println("i is set top", i)

	whatWasSaid, otherThingThatWasSaid := saySomething() // declare and initialize by function call

	fmt.Println("The function returned something", whatWasSaid, otherThingThatWasSaid)


	// 
	var x = 9
	fmt.Println(x)
}

func saySomething() (string, string) {
	return "something", "else"
}
