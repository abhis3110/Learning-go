package main

import (
	"fmt"
)

func main() {

	var whatToSay string // declare a variable of type string
	var i int

	whatToSay = "Goodbye, cruel world" // initialize

	fmt.Println(whatToSay)

	i = 7 // initialize

	fmt.Println("i is set top", i)

	whatWasSaid, otherThingThatWasSaid := saySomething() // declare and initialize by function call

	fmt.Println("The function returned something", whatWasSaid, otherThingThatWasSaid)

	// if initial value is known already, then type can be ommited
	var x = 9
	fmt.Println(x)
}

func saySomething() (string, string) {
	return "something", "else"
}
