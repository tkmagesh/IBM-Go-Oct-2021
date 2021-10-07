package utils

import "fmt"

type ShapeWithArea interface {
	Area() float64
}

type ShapeWithPerimeter interface {
	Perimeter() float64
}

//interface composition
type ShapeWithDimensions interface {
	ShapeWithArea
	ShapeWithPerimeter
}

func PrintArea(x ShapeWithArea) {
	fmt.Println("Area : ", x.Area())
}

func PrintPerimeter(s ShapeWithPerimeter) {
	fmt.Println("Perimeter : ", s.Perimeter())
}

func PrintDimensions(x ShapeWithDimensions) {
	fmt.Println("Dimensions")
	PrintArea(x)
	PrintPerimeter(x)
}
