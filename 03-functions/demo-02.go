package main

import "fmt"

func main() {
	fmt.Println(sum(10))
	fmt.Println(sum(10, 20))
	fmt.Println(sum(10, 20, 30))
	fmt.Println(sum(10, 20, 30, 40, 50))
}

func sum(x int, nos ...int) int {
	//nos => list of integers
	//len(nos) => number of integers
	//nos[0] => first integer

	result := x
	for idx := 0; idx < len(nos); idx++ {
		result += nos[idx]
	}
	return result
}
