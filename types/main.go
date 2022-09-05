package main

import (
	"fmt"
	"log"
)

//basic types (numbers, strings, booleans)
var myInt int
var myUint uint // only 0-positive values
var myFloat float32
var myFloat64 float64

func main() {
	myInt = 10
	myUint = 20
	myFloat = 10.1
	myFloat64 = 100.1

	log.Println(myInt, myUint, myFloat, myFloat64)

	myString := "Hi"
	log.Println(myString)
	myString = "Bye" //Strings are immutable in go

	var myBool = true
	myBool = false
	log.Println(myBool)

	//Aggregate types (array, struct)

	var myStrings [3]string
	myStrings[0] = "cat"
	myStrings[1] = "dog"
	myStrings[2] = "fish"

	fmt.Println("First element of array is: ", myStrings[0])

	type Car struct {
		NumberOfTires int
		Luxury        bool
		BucketSeats   bool
		Make          string
		Model         string
		Year          int
	}

	myCar := Car{
		NumberOfTires: 4,
		Luxury:        false,
		Make:          "Ford",
		BucketSeats:   false,
		Model:         "Fiesta",
		Year:          2022,
	}
	fmt.Printf("My car is a %d %s %s", myCar.Year, myCar.Make, myCar.Model)

}
