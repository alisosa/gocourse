package main

import "fmt"

func main() {

	//create a map that the key is string and contains integers
	//they are allways passed by reference (no need of pointers)
	intMap := make(map[string]int)

	intMap["one"] = 1
	intMap["two"] = 2
	intMap["three"] = 3
	intMap["four"] = 4
	intMap["five"] = 5

	//map are never sorted
	for key, value := range intMap {
		fmt.Println(key, value)
	}

	deleteFromMap(intMap, "two")

}

func deleteFromMap(m map[string]int, key string) {

	delete(m, key)
	el, ok := m[key]
	if ok {
		fmt.Println(el, "in in map")
	} else {
		fmt.Println(el, "is not in the map")
	}

}
