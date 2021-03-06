package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go fibonacci(ch)
	for no := range ch {
		fmt.Println(no)
	}
}

func fibonacci(ch chan int) {
	defer close(ch)
	x, y := 0, 1
	done := time.After(15 * time.Second)
	/* done := func() chan string {
		ch := make(chan string)
		go func() {
			time.Sleep(15 * time.Second)
			ch <- "done"
		}()
		return ch
	}() */
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
