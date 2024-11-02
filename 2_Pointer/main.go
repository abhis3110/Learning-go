package main

import "log"

func main() {
	var myString string
	myString = "Green"
	log.Println("My string is set to", myString)
	changeUsingPointer(&myString)
	log.Println("After function call my string is set to", myString)
}

func changeUsingPointer(s *string) {
	log.Println("String pointer", s)
	newValue := "Red"
	*s = newValue
}
 