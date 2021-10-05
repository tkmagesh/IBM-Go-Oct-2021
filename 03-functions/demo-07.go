package main

import (
	"errors"
	"fmt"
)

func main() {
	quotient, remainder, err := divide(100, 5)
	if err != nil {
		fmt.Println(err)
		return
	}
	//result, _ := divide(100, 0)
	fmt.Println("Result : ", quotient, remainder)
}

func divide(x, y int) (int, int, error) {
	if y == 0 {
		return 0, 0, errors.New("Cannot divide by zero")
	}
	return x / y, x % y, nil
}
