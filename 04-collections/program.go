package main

import "fmt"

func main() {
	//Array
	fmt.Println("Array")
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

	//Slices
	fmt.Println("Slice")
	var names []string
	names = append(names, "Magesh")
	names = append(names, "Suresh", "Ganesh")

	friends := []string{"Ramesh", "Rajesh"} //slice intialized with values
	names = append(names, friends...)
	fmt.Println(names)

	//Memory allocation
	fmt.Println("Memory allocation")
	//var randomNos []int
	//var randomNos []int = []int{9, 6, 3, 7, 2}
	/*
		var randomNos = make([]int, 0, 10)
		randomNos = append(randomNos, 9, 6, 3, 7, 2)
	*/
	var randomNos = make([]int, 5, 10)
	randomNos[0] = 9
	randomNos[1] = 6
	randomNos[2] = 3
	randomNos[3] = 7
	randomNos[4] = 2
	fmt.Printf("randomNos = %v, len() = %d, cap() = %d\n", randomNos, len(randomNos), cap(randomNos))
	fmt.Println("After appending new values")
	randomNos = append(randomNos, 11)
	fmt.Printf("randomNos = %v, len() = %d, cap() = %d\n", randomNos, len(randomNos), cap(randomNos))
	randomNos = append(randomNos, 22)
	fmt.Printf("randomNos = %v, len() = %d, cap() = %d\n", randomNos, len(randomNos), cap(randomNos))
	randomNos = append(randomNos, 33)
	fmt.Printf("randomNos = %v, len() = %d, cap() = %d\n", randomNos, len(randomNos), cap(randomNos))
	randomNos = append(randomNos, 44)
	fmt.Printf("randomNos = %v, len() = %d, cap() = %d\n", randomNos, len(randomNos), cap(randomNos))
	randomNos = append(randomNos, 55)
	fmt.Printf("randomNos = %v, len() = %d, cap() = %d\n", randomNos, len(randomNos), cap(randomNos))
	randomNos = append(randomNos, 66)
	fmt.Printf("randomNos = %v, len() = %d, cap() = %d\n", randomNos, len(randomNos), cap(randomNos))
}
