package main

import "fmt"

func main() {
	var dog Animal
	dog.Name = "dog"
	dog.NumberOfLegs = 4
	dog.Sound = "woof"
	dog.Says()

	cat := Animal{
		Name:         "cat",
		Sound:        "meow",
		NumberOfLegs: 4,
	}
	cat.Says()
	cat.HowManyLegs()

	fmt.Println(sumTotal(1, 2, 3, 4, 5, 6, 7))
}

type Animal struct {
	Name         string
	Sound        string
	NumberOfLegs int
}

func (a *Animal) Says() {
	fmt.Printf("A %s says %s", a.Name, a.Sound)
	fmt.Println()
}

func (a *Animal) HowManyLegs() {
	fmt.Printf("A %s has %d legs", a.Name, a.NumberOfLegs)
	fmt.Println()
}

func sumTotal(nums ...int) int {
	total := 0

	for _, x := range nums {
		total = total + x
	}
	return total

}
