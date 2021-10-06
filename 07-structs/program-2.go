package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float64
	Units    int
	Category string
}

/*
type PerishableProduct struct {
	Prod   Product
	Expiry string
}

func main() {
	grapes := PerishableProduct{
		Prod:   Product{102, "Grapes", 60, 10, "Food"},
		Expiry: "2 Days",
	}
	fmt.Printf("%#v\n", grapes)
	fmt.Println(grapes.Prod.Category)
	applyDiscount(&grapes.Prod, 0.10)
	fmt.Println("After applying discount :")
	fmt.Printf("%#v\n", grapes)
}
*/

/*
type PerishableProduct struct {
	Product Product
	Expiry  string
}
*/

type PerishableProduct struct {
	Product
	Expiry string
}

func NewPerishableProduct(id int, name string, cost float64, units int, category string, expiry string) PerishableProduct {
	return PerishableProduct{
		Product: Product{id, name, cost, units, category},
		Expiry:  expiry,
	}
}

func main() {
	/*
		grapes := PerishableProduct{
			Product: Product{102, "Grapes", 60, 10, "Food"},
			Expiry:  "2 Days",
		}
	*/
	grapes := NewPerishableProduct(102, "Grapes", 60, 10, "Food", "2 Days")
	fmt.Printf("%#v\n", grapes)
	fmt.Println(grapes.Category)
	applyDiscount(&grapes.Product, 0.10)
	fmt.Println("After applying discount :")
	//fmt.Printf("%#v\n", grapes)
	Print(grapes.Product)
}

func applyDiscount(product *Product, discount float64) {
	//(*product).Cost = (*product).Cost - ((*product).Cost * discount)
	product.Cost = product.Cost * (1 - discount)
}

func Print(product Product) {
	fmt.Printf("Id = %d, Name = %s, Cost = %f, Units = %d, Category = %s\n", product.Id, product.Name, product.Cost, product.Units, product.Category)
}
