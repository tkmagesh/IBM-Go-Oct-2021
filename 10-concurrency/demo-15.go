package main

import (
	"fmt"
	"time"
)

/*
Modify the program in such a way that the program keeps generating the fibonacci numbers until the user hits "ENTER" key
*/

func main() {
	ch := make(chan int)
	done := make(chan string)
	go fibonacci(ch, done)
	go func() {
		var input string
		fmt.Scanln(&input)
		done <- "done"
	}()
	for no := range ch {
		fmt.Println(no)
	}
}

func fibonacci(ch chan int, done chan string) {
	defer close(ch)
	x, y := 0, 1
	for {
		select {
		case ch <- x:
			time.Sleep(500 * time.Millisecond)
			x, y = y, x+y
		case <-done:
			return
		}
	}
}
