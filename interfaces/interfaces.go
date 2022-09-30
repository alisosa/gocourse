package main

import "fmt"

type Animal interface {
	Says() string
	HowManyLegs() int
}

type Dog struct {
	Name         string
	Sound        string
	NumberOfLegs int
}

type Cat struct {
	Name         string
	Sound        string
	NumberOfLegs int
	HasTail      bool
}

//In order to be an Animal, dog and cat has to implement
// says and howmanylegs, and automatically will be animal

func (d *Dog) Says() string {
	return d.Sound
}

func (d *Dog) HowManyLegs() int {
	return d.NumberOfLegs
}

func (c *Cat) Says() string {
	return c.Sound
}

func (c *Cat) HowManyLegs() int {
	return c.NumberOfLegs
}

func main() {
	dog := Dog{
		Name:         "Rocco",
		Sound:        "woof",
		NumberOfLegs: 4,
	}
	cat := Cat{
		Name:         "Mango",
		Sound:        "meow",
		NumberOfLegs: 4,
	}

	Riddle(&dog)
	Riddle(&cat)
}

func Riddle(a Animal) {
	riddle := fmt.Sprintf("This animal says %s and has %d legs, what animal is it?", a.Says(), a.HowManyLegs())
	fmt.Println(riddle)
}
