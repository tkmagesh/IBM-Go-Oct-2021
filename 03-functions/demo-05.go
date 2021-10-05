package main

import "fmt"

func main() {
	exec(helloWorld)
	processOp(add, 100, 200)
	processOp(subtract, 100, 200)
}

func helloWorld() {
	fmt.Println("Hello World!")
}

func exec(fn func()) {
	fn()
}

func add(x, y int) int {
	return x + y
}

func subtract(x, y int) int {
	return x - y
}

func processOp(op func(int, int) int, x, y int) {
	fmt.Println("Before invocation")
	fmt.Println("Result : ", op(x, y))
	fmt.Println("After invocation")
}
