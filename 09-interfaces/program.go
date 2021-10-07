package main

import (
	"fmt"
	"math"
)

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

type Rectangle struct {
	Height float64
	Width  float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Height + r.Width)
}

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

func main() {
	c := Circle{5}
	//fmt.Println("Area : ", c.Area())
	/*
		PrintArea(c)
		PrintPerimeter(c)
	*/
	PrintDimensions(c)

	r := Rectangle{5, 10}
	//fmt.Println("Area : ", r.Area())
	/*
		PrintArea(r)
		PrintPerimeter(r)
	*/
	PrintDimensions(r)
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

/*
circle perimeter = 2πr
rectangle perimeter = 2(h+w)
*/
