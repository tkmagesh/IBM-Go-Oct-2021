package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("[@main] Deferred")
	}()
	fmt.Println(f1())
	fmt.Println("main invoked")
}

func f1() (result int) {
	defer func() {
		fmt.Println("[@f1] Deferred - 1")
		result = 500
	}()
	/* defer func() {
		fmt.Println("[@f1] Deferred - 2")
	}() */
	defer fmt.Println("[@f1] Deferred - 2")

	defer func() {
		fmt.Println("[@f1] Deferred - 3")
		result = 200
	}()
	fmt.Println("f1 invoked")
	f2()
	result = 100
	return
}

func f2() {
	defer func() {
		fmt.Println("[@f2] Deferred - 1")
	}()
	defer func() {
		fmt.Println("[@f2] Deferred - 2")
	}()
	defer func() {
		fmt.Println("[@f2] Deferred - 3")
	}()
	fmt.Println("f2 invoked")
}
