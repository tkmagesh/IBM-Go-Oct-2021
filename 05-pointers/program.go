package main

import "fmt"

func main() {
	var no int = 10
	var noPtr *int = &no

	//dereferencing
	var value = *noPtr
	fmt.Println(no, noPtr, value)

	fmt.Println("Before incrementing, no = ", no)
	increment(&no)
	fmt.Println("After incrementing, no = ", no)

	x, y := 10, 20
	fmt.Println("Before swapping, x = ", x, "y = ", y) // => 10,20
	swap(&x, &y)
	fmt.Println("After swapping, x = ", x, "y = ", y) // => 20, 10
}

func increment(n *int) {
	*n++
}

func swap(x, y *int) {
	*x, *y = *y, *x
}
