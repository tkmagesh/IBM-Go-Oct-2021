package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float64
	Units    int
	Category string
}

func (product Product) Format() string {
	return fmt.Sprintf("Id = %d, Name = %s, Cost = %f, Units = %d, Category = %s", product.Id, product.Name, product.Cost, product.Units, product.Category)
}

func (product *Product) applyDiscount(discount float64) {
	//(*product).Cost = (*product).Cost - ((*product).Cost * discount)
	product.Cost = product.Cost * (1 - discount)
}

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

func (pp PerishableProduct) Format() string {
	return fmt.Sprintf("%s, Expiry = %s", pp.Product.Format(), pp.Expiry)
}

func main() {

	product := &Product{Id: 100, Cost: 5, Units: 10, Name: "Pen", Category: "Stationary"}
	fmt.Println(product.Format())
	fmt.Println("After applying 10% discount")
	//applyDiscount(&product, 0.1)
	product.applyDiscount(0.1)
	fmt.Println(product.Format())

	grapes := NewPerishableProduct(101, "Grapes", 10, 10, "Fruit", "20/10/2020")
	//applyDiscount(&grapes.Product, 0.1)
	grapes.applyDiscount(0.1)
	fmt.Println(grapes.Format())
}
