package main

import "fmt"

// reference types (pointers, slices, maps, functions, channels)

//interface types

func main() {
	x := 10

	//Reference to a pointer
	myFirstPointer := &x

	fmt.Println("x is:", x)
	fmt.Println("myFirstPointer is: ", myFirstPointer)

	//Get the value of a pointer and change it
	*myFirstPointer = 15
	fmt.Println("x is now", x)

	changeValueOfPointer(&x)

	fmt.Println("After func, x is now ", x)
}

func changeValueOfPointer(num *int) {
	*num = 25
}
