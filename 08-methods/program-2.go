package main

import "fmt"

type MyString string //alias for string

func (s MyString) Length() int {
	return len(s)
}

func main() {
	var s MyString = MyString("Cillum do eu esse labore excepteur sunt sint incididunt minim id aute. Tempor ullamco nulla id cupidatat ullamco fugiat.")
	fmt.Println(s.Length())
}
