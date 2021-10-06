package main

import (
	"errors"
	"fmt"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("[@main] Error:", err)
		}
	}()
	result, err := divide(100, 0)
	if err != nil {
		fmt.Println("Something went wrong : ", err)
		return
	}
	fmt.Println(result)
}

func divide(x, y int) (result int, err error) {
	defer func() {
		if er := recover(); er != nil {
			fmt.Println("[@divide] Error:", er)
			err = errors.New("Cannot divide by zero!")
			//can i panic from here?
		}
	}()
	if y == 0 {
		panic("Cannot divide by zero!")
	}
	result, err = x/y, nil
	return
}
