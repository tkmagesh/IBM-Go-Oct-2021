package entities

import (
	"fmt"
	"sort"
)

type Products []Product

//"Stringer" interface implementation
func (products Products) String() string {
	var result string
	for _, p := range products {
		result += fmt.Sprintf("%v\n", p)
	}
	return result
}

//'Interface' interface implementation

func (p Products) Len() int {
	return len(p)
}

func (p Products) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p Products) Less(i, j int) bool {
	return p[i].Id < p[j].Id
}

type byName struct {
	Products
}

func (p byName) Less(i, j int) bool {
	return p.Products[i].Name < p.Products[j].Name
}

type byCost struct {
	Products
}

func (p byCost) Less(i, j int) bool {
	return p.Products[i].Cost < p.Products[j].Cost
}

func (products Products) Sort(attrName string) {
	switch attrName {
	case "name":
		sort.Sort(byName{products})
	case "cost":
		sort.Sort(byCost{products})
	case "id":
		sort.Sort(products)
	default:
		sort.Sort(products)
	}
}
