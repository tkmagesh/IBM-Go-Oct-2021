package main

import (
	"fmt"
	"modularity-demo/calculator"
	"modularity-demo/calculator/utils"

	"github.com/fatih/color"
)

func main() {

	color.Red(fmt.Sprintf("%t\n", utils.IsEven(200)))
	color.Blue(fmt.Sprintf("%d\n", calculator.Add(100, 200)))
	color.Green(fmt.Sprintf("%d\n", calculator.Subtract(100, 200)))
	color.Yellow(fmt.Sprintf("%d\n", calculator.GetOpCount()))

	fmt.Println(utils.IsEven(200))
	fmt.Println(calculator.Add(100, 200))

}
