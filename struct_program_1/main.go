package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})

	v := Vertex{} // creating new struct
	v.X = 4
	fmt.Println(v)

	// struct and pointer
	s := Vertex{5, 6}
	p := &s
	(*p).X = 1e9 // p.X is aslo valid in go for same operation
	fmt.Println(s)
}
