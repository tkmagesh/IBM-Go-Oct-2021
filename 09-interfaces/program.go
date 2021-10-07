package main

import (
	"fmt"
	"interfaces-demo/entities"
	"interfaces-demo/shapes"
	"interfaces-demo/utils"
)

func main() {
	c := shapes.Circle{5}
	utils.PrintDimensions(c)

	r := shapes.Rectangle{5, 10}
	utils.PrintDimensions(r)

	products := entities.Products{
		entities.Product{105, "Pen", 5, 50, "Stationary"},
		entities.Product{107, "Pencil", 2, 100, "Stationary"},
		entities.Product{103, "Marker", 50, 20, "Utencil"},
		entities.Product{102, "Stove", 5000, 5, "Utencil"},
		entities.Product{101, "Kettle", 2500, 10, "Utencil"},
		entities.Product{104, "Scribble Pad", 20, 20, "Stationary"},
	}

	fmt.Println("Sorting")
	fmt.Println("Default Sort (by id)")
	products.Sort("id")
	fmt.Println(products)

	fmt.Println()
	fmt.Println("Sort by name")
	//sort.Sort(ByName{products})
	products.Sort("name")
	fmt.Println(products)

	fmt.Println()
	fmt.Println("Sort by cost")
	//sort.Sort(ByCost{products})
	products.Sort("cost")
	fmt.Println(products)
}

/*
circle perimeter = 2πr
rectangle perimeter = 2(h+w)
*/
