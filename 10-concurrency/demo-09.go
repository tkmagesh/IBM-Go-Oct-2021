package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go print("Hello", ch1, ch2)
	go print("World", ch2, ch1)
	ch1 <- "start"
	var input string
	fmt.Scanln(&input)
}

func print(s string, in, out chan string) {
	for i := 0; i < 5; i++ {
		<-in
		time.Sleep(500 * time.Millisecond)
		fmt.Println(s)
		out <- "done"
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
