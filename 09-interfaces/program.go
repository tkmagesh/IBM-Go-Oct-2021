package main

import (
	"interfaces-demo/shapes"
	"interfaces-demo/utils"
)

func main() {
	c := shapes.Circle{5}
	//fmt.Println("Area : ", c.Area())
	/*
		PrintArea(c)
		PrintPerimeter(c)
	*/
	utils.PrintDimensions(c)

	r := shapes.Rectangle{5, 10}
	//fmt.Println("Area : ", r.Area())
	/*
		PrintArea(r)
		PrintPerimeter(r)
	*/
	utils.PrintDimensions(r)
}

/*
circle perimeter = 2πr
rectangle perimeter = 2(h+w)
*/
