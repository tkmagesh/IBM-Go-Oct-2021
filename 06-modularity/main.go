package main

import (
	"fmt"
	"modularity-demo/calculator"
	"modularity-demo/calculator/utils"
)

func main() {
	fmt.Println(utils.IsEven(200))
	fmt.Println(calculator.Add(100, 200))
	fmt.Println(calculator.Subtract(100, 200))
	fmt.Println(calculator.GetOpCount())

}
