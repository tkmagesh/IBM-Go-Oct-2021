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
	/* var product Product = Product{Id: 100, Cost: 5, Units: 10, Name: "Pen"} */
	product := Product{Id: 100, Cost: 5, Units: 10, Name: "Pen"}
	fmt.Printf("%#v\n", product)
	fmt.Println("After applying 10% discount")
	applyDiscount(&product, 0.1)
	fmt.Printf("%#v\n", product)

	var ptrProduct = &product
	fmt.Println(ptrProduct.Name)
}

func applyDiscount(product *Product, discount float64) {
	//(*product).Cost = (*product).Cost - ((*product).Cost * discount)
	product.Cost = product.Cost * (1 - discount)
}
