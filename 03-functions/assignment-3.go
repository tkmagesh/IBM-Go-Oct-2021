/*
	Display the following menu until the user choses to exit
		1. Add
		2. Subtract
		3. Multiply
		4. Divide
		5. Exit

	If the user enters any other choice, display "Invalid choice" and display the menu again.
*/

package main

import "fmt"

func main() {
	for {
		var result int
		choice := getUserChoice()
		if choice == 5 {
			return
		}
		if choice < 1 || choice > 5 {
			fmt.Println("Invalid choice")
			continue
		}
		op := getOperation(choice)
		x, y := getOperands()
		result = op(x, y)
		fmt.Println("Result : ", result)
	}
}

func getOperation(choice int) (op func(int, int) int) {
	switch choice {
	case 1:
		op = add
	case 2:
		op = subtract
	case 3:
		op = multiply
	case 4:
		op = divide
	}
	return op
}

func add(x, y int) int {
	return x + y
}

func subtract(x, y int) int {
	return x - y
}

func multiply(x, y int) int {
	return x * y
}
func divide(x, y int) int {
	return x / y
}

func getUserChoice() int {
	fmt.Println("1. Add")
	fmt.Println("2. Subtract")
	fmt.Println("3. Multiply")
	fmt.Println("4. Divide")
	fmt.Println("5. Exit")
	var choice int
	fmt.Scanln(&choice)
	return choice
}

func getOperands() (int, int) {
	fmt.Println("Enter the numbers : ")
	var x, y int
	fmt.Scanf("%d %d\n", &x, &y)
	return x, y
}
