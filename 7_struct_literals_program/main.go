package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}
	v2 = Vertex{X: 1}
	v3 = Vertex{}
	p  = &Vertex{1, 2} // pointer to new vertex created
)

func main() {
	fmt.Println(v1, (*p).X, v2, v3)
	fmt.Println(*p)  // Output: {1 2} (dereferencing the pointer)
}
