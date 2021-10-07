package main

import "fmt"

//package level unused variable - allowed
var z int = 200

type X struct {
	y int
}

func main() {

	const o = X{y: 100}
	/*
		var msg string
		msg = "Hello, World!"
	*/

	/*
		var msg string = "Hello, World!"
	*/

	// Type Inference

	/*
		var msg = "Hello, World!"
	*/
	msg := "Hello, World!" // := is applicable only in a function
	fmt.Println(msg)

	//Handling multiple variables
	/*
		var x int
		var y int
		var result int
		var message string

		x = 100
		y = 200
		message = "Result : "
		result = x + y
	*/
	/*
		var x int = 100
		var y int = 200
		var result int = x + y
		var message string = "Result :"
	*/

	/*
		var x, y, result int
		var message string = "Result :"
		x = 100
		y = 200
		result = x + y
	*/

	/*
		var (
			x, y, result int
			message      string
		)
		x = 100
		y = 200
		result = x + y
		message = "Result : "
	*/

	/*
		var (
			x, y, result int = 100, 200, 0
			message      string
		)
	*/
	/*
		x, y, result := 100, 200, 0
		message := "Result :"
	*/
	/*
		x, y, result, message := 100, 200, 0, "Result :"
		result = x + y
	*/

	x, y, message := 100, 200, "Result :"
	result := x + y

	fmt.Println(message, result)

	//Constants
	const z = 100
	//z = 200
	fmt.Println(z)

	//iota (represents integer values from 0)
	/*
		const (
			red   = iota
			green = iota
			blue  = iota
		)

		const (
			pen    = iota
			pencil = iota
			marker = iota
		)
		fmt.Println(red, green, blue)
		fmt.Println(pen, pencil, marker)
	*/

	/*
		const (
			red = iota
			green
			blue
		)
	*/
	/*
		const (
			red = iota + 1
			green
			blue
		)
	*/

	const (
		red = iota + 1
		green
		_
		blue
	)

	fmt.Println(red, green, blue)

}
