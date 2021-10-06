package utils

import "fmt"

func init() {
	fmt.Println("utils initialized")
}

func IsEven(n int) bool {
	return n%2 == 0
}
