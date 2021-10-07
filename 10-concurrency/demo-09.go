package main

import (
	"fmt"
	"time"
)

func main() {
	print("Hello")
	print("World")
	var input string
	fmt.Scanln(&input)
}

func print(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Println(s)
	}
}

/*
Hello
World
Hello
World
.
.
.
*/
