package main

import "fmt"

type Entity struct {
}

func (e Entity) M1() {

}

type Contract interface {
	M1()
}

func main() {
	var x interface{}
	x = 100
	x = "Hello"
	//x = true
	//x = struct{}{}

	if val, ok := x.(int); ok {
		y := val + 100
		fmt.Println(y)
	} else {
		fmt.Println("x is not an int")
	}

	//x = 1000
	//x = "Hello World!"
	//x = true
	//x = struct{}{}
	//x = []int{3, 1, 2, 4}
	//x = map[string]int{"a": 1, "b": 2}
	x = Entity{}

	switch val := x.(type) {
	case int:
		fmt.Println("Double of x is ", val*2)
	case string:
		fmt.Println("Length of x is ", len(val))
	case bool:
		fmt.Println("x is a boolean")
	case struct{}:
		fmt.Println("x is a struct")
	case []int:
		fmt.Println("x is a slice of integers")
	case map[string]int:
		fmt.Println("x is a map of strings to integers")
	case Contract:
		fmt.Println("x is an object implementing the 'Contract' interface")
	default:
		fmt.Println("Unknown type")
	}
}
