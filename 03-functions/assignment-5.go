package main

import "fmt"

func main() {
	increment, decrement := getCounter()
	fmt.Println(increment()) //=> 1
	fmt.Println(increment()) //=> 2
	fmt.Println(increment()) //=> 3
	fmt.Println(increment()) //=> 4

	//counter = 10000 //influencing the outcome of the increment function

	fmt.Println(decrement()) // => 3
	fmt.Println(decrement()) // => 2
	fmt.Println(decrement()) // => 1
	fmt.Println(decrement()) // => 0
}

func getCounter() (func() int, func() int) {

	var counter int
	increment := func() int {
		counter++
		return counter
	}
	decrement := func() int {
		counter--
		return counter
	}
	return increment, decrement
}
