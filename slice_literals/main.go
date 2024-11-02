package main 


import "fmt"

func main() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q, len(q))
	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct{
		i int 
		b bool
		}{
			{2, true}, 
			{3, false}, 
			{4, false}, 
			{1, true},
		}
	fmt.Println(s)
} 