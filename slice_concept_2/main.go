package main

import "fmt"

func main() {

	// primes is an array fixed size
	primes := [6]int{2, 3, 5, 7, 11, 13} 
	var x int32=320000000

	fmt.Printf("Type is %T\n",x)


	// s is slice can increase size
	var s []int = primes[1:4] 
	fmt.Println(s)

	// append is not possible in array
	fmt.Println(cap(s), len(s))
	arr1:= append(s, 11, 13, 17, 23) 
	fmt.Println(&s[0], &arr1[0])


	// slices are like refrence to arrays
	names := [4]string {"John", "Paul", "George", "Ringo"}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}
