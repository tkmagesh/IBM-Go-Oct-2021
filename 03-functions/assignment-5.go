package main

import "fmt"

func main() {
	increment := getIncrement()
	fmt.Println(increment()) //=> 1
	fmt.Println(increment()) //=> 2
	fmt.Println(increment()) //=> 3
	fmt.Println(increment()) //=> 4

	//counter = 10000 //influencing the outcome of the increment function

	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
}

func getIncrement() func() int {
	var counter int
	return func() int {
		counter++
		return counter
	}
}
