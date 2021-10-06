package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float64
	Units    int
	Category string
}

func main() {
	/*
		var product Product
		product.Id = 101
		product.Name = "Pen"
		product.Cost = 5
		product.Units = 100
		product.Category = "Stationery"
	*/

	/*
		var product Product = Product{101, "Pen", 5, 100, "Stationery"}
	*/
	var product Product = Product{Id: 100, Cost: 5, Units: 10, Name: "Pen"}
	fmt.Printf("%#v\n", product)
}
