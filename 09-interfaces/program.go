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

type Rectangle struct {
	Height float64
	Width  float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

type ShapeWithArea interface {
	Area() float64
}

func main() {
	c := Circle{5}
	//fmt.Println("Area : ", c.Area())
	PrintArea(c)
	PrintPerimeter(c)

	r := Rectangle{5, 10}
	//fmt.Println("Area : ", r.Area())
	PrintArea(r)
	PrintPerimeter(r)
}

func PrintArea(x ShapeWithArea) {
	fmt.Println("Area : ", x.Area())
}

func PrintPerimeter(x /*  */) {
	fmt.Println("Perimeter : ", x.Perimeter())
}

/*
circle perimeter = 2πr
rectangle perimeter = 2(h+w)
*/
