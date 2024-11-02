package main

import "fmt"

func calSquare(number int, squareChan chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit
		number /= 10
	}
	squareChan <- sum
}

func calCube(number int, cubeChan chan int) {
	fmt.Println(cubeChan)
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit * digit
		number /= 10
	}
	cubeChan <- sum
}

func main() {
	number := 123
	squareChan := make(chan int)
	cubeChan := make(chan int)
	fmt.Println(cubeChan)
	go calSquare(number, squareChan)
	go calCube(number, cubeChan)

	squareNum, cubeNum := <-squareChan, <-cubeChan
	fmt.Println("Sum of sqaure of digit and cube of digit is", squareNum + cubeNum)
}
