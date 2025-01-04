package main

import (
	"log"
	"time"
)

var s = "Seven"

var firstName string
var secondName string
var phoneNumber string
var age int
var dob time.Time

func main() {
	var s2 = "Six"

	log.Println("s is ", s)
	log.Println("s2 is", s2)

	saySomething("xxx")
}

// variable shadowing or overriding
// first priority to local variable
func saySomething(s string) (string, string) {
	log.Println("s from saySomething func", s)
	return s, "World"
}
