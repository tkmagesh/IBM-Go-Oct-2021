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
	marker := Product{103, "Marker", 50, 20, "Utencil"}
	products := []Product{
		Product{105, "Pen", 5, 50, "Stationary"},
		Product{107, "Pencil", 2, 100, "Stationary"},
		Product{103, "Marker", 50, 20, "Utencil"},
		Product{102, "Stove", 5000, 5, "Utencil"},
		Product{101, "Kettle", 2500, 10, "Utencil"},
		Product{104, "Scribble Pad", 20, 20, "Stationary"},
	}

	fmt.Println(marker.Format())

	fmt.Println("Initial List")
	fmt.Println(FormatProducts(products))
	fmt.Println("Index of Marker = ", IndexOf(products, marker))

	costlyProductPredicate := func(product Product) bool {
		return product.Cost > 1000
	}

	costlyProducts := Filter(products, costlyProductPredicate)
	fmt.Println("Costly Products")
	fmt.Println(FormatProducts(costlyProducts))
}

func (product Product) Format() string {
	return fmt.Sprintf("Id = %d, Name = %s, Cost = %f, Units = %d, Category = %s", product.Id, product.Name, product.Cost, product.Units, product.Category)
}

func FormatProducts(products []Product) string {
	var result string
	for _, p := range products {
		result += fmt.Sprintf("%v\n", p.Format())
	}
	return result
}

func IndexOf(products []Product, product Product) int {
	for idx, p := range products {
		if p == product {
			return idx
		}
	}
	return -1
}

func Includes(products []Product, product Product) bool {
	for _, p := range products {
		if p == product {
			return true
		}
	}
	return false
}

func Filter(products []Product, predicate func(Product) bool) []Product {
	var result []Product
	for _, p := range products {
		if predicate(p) {
			result = append(result, p)
		}
	}
	return result
}

/*
Write the following function
IndexOf => return the index of the given product in the collection
Includes => return true if the given product is present in the collection else return false
Filter => return products that matches the given criteria
	Use cases:
		1. filter all costly products (cost > 1000)
		2. filter all products that are stationary
All => return true if ALL the products matches the given criteria else return false
	Use cases:
		1. Are all the products costly products (cost > 1000) ?
		2. Are all the products stationary products? (category == "Stationary")

Any => return true if ANY of the products matches the given criteria else returns false
	Use cases:
		1. Are there ANY costly product (cost > 1000) ?
		2. Are there ANY stationary products? (category == "Stationary")

*/
