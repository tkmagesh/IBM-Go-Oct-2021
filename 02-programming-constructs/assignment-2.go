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
		x, y := getOperands()
		switch choice {
		case 1:
			result = add(x, y)
			fmt.Println(result)
		case 2:
			result = subtract(x, y)
			fmt.Println(result)
		case 3:
			result = multiply(x, y)
			fmt.Println(result)
		case 4:
			result = divide(x, y)
			fmt.Println(result)
		}
	}
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
