package main

import "fmt"

func main() {
	logAdd := loggedOp(add)
	logAdd(10, 20)

	logSubtract := loggedOp(subtract)
	logSubtract(10, 20)
}

func add(x, y int) int {
	return x + y
}

func subtract(x, y int) int {
	return x - y
}

func loggedOp(op func(int, int) int) func(int, int) {
	return func(x, y int) {
		fmt.Println("Before invocation")
		fmt.Println("Result : ", op(x, y))
		fmt.Println("After invocation")
	}
}
