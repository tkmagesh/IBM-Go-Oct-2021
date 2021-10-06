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
		fmt.Println(&nos[idx], nos[idx])
	}

	//Iterate over array (using the range construct)
	fmt.Println("Iterate over array (using the range construct)")
	for _, val := range nos {
		fmt.Println(val)
	}

	//creating a copy
	fmt.Println("creating a copy of an array")
	newNos := nos
	fmt.Println(newNos)

	fmt.Println("after changing the copy of the array")
	newNos[0] = 10
	fmt.Printf("nos = %#v, newNos = %#v\n", nos, newNos)

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
	fmt.Printf("randomNos = %v, len() = %d, cap() = %d, &randomNos[0] = %p\n", randomNos, len(randomNos), cap(randomNos), &randomNos[0])
	fmt.Println("After appending new values")
	randomNos = append(randomNos, 11)
	fmt.Printf("randomNos = %v, len() = %d, cap() = %d, &randomNos[0] = %p\n", randomNos, len(randomNos), cap(randomNos), &randomNos[0])
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

	fmt.Println("Copy of a slice")
	//var newRandomNos []int = randomNos
	var newRandomNos = make([]int, len(randomNos))
	copy(newRandomNos, randomNos)
	//var newRandomNos []int = randomNos[:]
	fmt.Printf("randomNos = %v, newRandomNos = %v\n", randomNos, newRandomNos)

	fmt.Println("after changing the copy of the slice")
	newRandomNos[0] = 10
	fmt.Printf("randomNos = %#v, newRandomNos = %#v\n", randomNos, newRandomNos)

	//slicing
	/*
		[lo : hi] => from lo to hi-1
		[lo : ] => from lo to end
		[ : hi] => from 0 to hi-1
		[lo : lo] => empty slice
		[lo : lo+1] => single element slice with [lo]
		[:] => copy of the slice

	*/
	fmt.Println("Slicing")
	fmt.Println("randomNos[2:5] = ", randomNos[2:5])
	fmt.Println("randomNos[2:] = ", randomNos[2:])
	fmt.Println("randomNos[:5] = ", randomNos[:5])

	//Maps
	fmt.Println("Maps")
	var cityRanks map[string]int = map[string]int{
		"Bengaluru": 4,
		"Udupi":     1,
		"Mangaluru": 2,
		"Mysuru":    3,
	}
	fmt.Println("cityRanks = ", cityRanks)
	fmt.Println("cityRanks[\"Bengaluru\"] = ", cityRanks["Bengaluru"])

	//appending a value to a map
	fmt.Println("Appending a value to a map")
	cityRanks["Pune"] = 5
	fmt.Println("cityRanks = ", cityRanks)

	//checking if a key exists
	fmt.Println("Checking if a key exists")
	if value, exists := cityRanks["Kochi"]; exists {
		fmt.Println("cityRanks[\"Kochi\"] = ", value)
	} else {
		fmt.Println("cityRanks[\"Kochi\"] does not exist")
	}

	//removing a key from a map
	fmt.Println("Removing a key from a map")
	delete(cityRanks, "Pune")
	fmt.Println("cityRanks = ", cityRanks)

	//Iterating over a map
	fmt.Println("Iterating over a map")
	for city, rank := range cityRanks {
		fmt.Printf("city = %s, rank = %d\n", city, rank)
	}

}
