package main

import "fmt"

func main() {

	// it statement
	/*
		no := 21
		if no%2 == 0 {
			fmt.Println(no, " is even")
		} else {
			fmt.Println(no, " is odd")
		}
	*/

	if no := 21; no%2 == 0 {
		fmt.Println(no, " is even")
	} else {
		fmt.Println(no, " is odd")
	}

	//for statement

	//ver 1.0
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	//ver 2.0 (simulating a while statement)
	numSum := 1
	for numSum < 100 {
		numSum += numSum
	}
	fmt.Println(numSum)

	//ver 3.0 (simulating an indefinite loop)
	x := 1
	for {
		if x > 100 {
			break
		}
		x += x
	}
	fmt.Println(x)

}
