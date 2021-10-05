package main

import "fmt"

func main() {

	//functions assigned to variables
	var helloWorld func()
	helloWorld = func() {
		fmt.Println("Hello World")
	}
	helloWorld()

	var add func(int, int) int
	add = func(x, y int) int {
		return x + y
	}
	fmt.Println(add(100, 200))

	//anonymous functions
	func() {
		fmt.Println("Hi there!")
	}()

	//anonymous functions with parameters
	func(str string) {
		fmt.Println(str)
	}("Hello there!")

	func(x, y int) {
		fmt.Println(x + y)
	}(100, 200)

	//anonymous functions with parameters and return values
	result := func(x, y int) int {
		return x * y
	}(10, 20)
	fmt.Println(result)

}
