package main

import "fmt"

func main() {
	fmt.Println(isPrime(11))
	fmt.Println(isPrime(12))
	fmt.Println(add(100, 200))
	//fmt.Println(divide(100, 7))
	/*
		quotient, remainder := divide(100, 7)
		fmt.Println(quotient, remainder)
	*/

	quotient, _ := divide(100, 7)
	fmt.Println(quotient)
}

/*
func isPrime(no int) bool {
	for j := 2; j <= (no / 2); j++ {
		if no%j == 0 {
			return false
		}
	}
	return true
}
*/

func isPrime(no int) (result bool) {
	for j := 2; j <= (no / 2); j++ {
		if no%j == 0 {
			result = false
			return
		}
	}
	result = true
	return
}

/*
func add(x int, y int) (result int) {
	result = x + y
	return
}
*/

func add(x, y int) (result int) {
	result = x + y
	return
}

//returning more than one value

/*
func divide(x, y int) (int, int) {
	quotient := x / y
	remainder := x % y
	return quotient, remainder
}
*/

func divide(x, y int) (quotient int, remainder int) {
	quotient = x / y
	remainder = x % y
	return
}
