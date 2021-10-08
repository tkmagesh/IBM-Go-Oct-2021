package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	ch1 := make(chan string, 1)
	ch2 := make(chan string)
	wg.Add(2)
	go print("Hello", ch1, ch2, wg)
	go print("World", ch2, ch1, wg)
	ch1 <- "start"
	wg.Wait()
}

func print(s string, in, out chan string, wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		<-in
		time.Sleep(500 * time.Millisecond)
		fmt.Println(s)
		out <- "done" /* TODO : Fix this */
	}
	wg.Done()
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
