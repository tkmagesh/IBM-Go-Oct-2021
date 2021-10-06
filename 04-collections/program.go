package main

import "fmt"

func main() {
	//Array
	//var nos [5]int
	var nos [5]int = [5]int{3, 5, 4, 2, 1}
	fmt.Println(nos)

	//Iterate over array (using the indexer)
	fmt.Println("Iterate over array (using the indexer)")
	for idx := 0; idx < len(nos); idx++ {
		fmt.Println(nos[idx])
	}

	//Iterate over array (using the range construct)
	fmt.Println("Iterate over array (using the range construct)")
	for _, val := range nos {
		fmt.Println(val)
	}
}
