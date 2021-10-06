package calculator

import "fmt"

var opCount int

func init() {
	fmt.Println("Calculator package initialized")
}

func GetOpCount() int {
	return opCount
}

/*
calculator1
calculator2
main.go

	calculator1
main
	calculator2

main -> calculator1 -> calculator2

*/
