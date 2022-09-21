package main

import (
	"fmt"
	"sort"
)

func main() {

	var animals []string
	animals = append(animals, "dog")
	animals = append(animals, "fish")
	animals = append(animals, "cat")
	animals = append(animals, "horse")

	fmt.Println(animals)

	//for (index) (element) in range animals
	for i, x := range animals {

		fmt.Println(i, x)
	}

	fmt.Println("Element 0 is", animals[0])

	fmt.Println("First two elements are", animals[0:2])

	fmt.Println("the slice is", len(animals), "elements long")

	fmt.Println("is it sorted?", sort.StringsAreSorted(animals))

	if !sort.StringsAreSorted(animals) {
		sort.Strings(animals)
	}

	fmt.Println("is it sorted now?", sort.StringsAreSorted(animals))
	fmt.Println(animals)

	fmt.Println("dellete cat")
	animals = deleteFromSlice(animals, 0)
	fmt.Println(animals)
}

func deleteFromSlice(a []string, i int) []string {
	a[i] = a[len(a)-1]
	a[len(a)-1] = ""
	a = a[:len(a)-1]
	return a
}
