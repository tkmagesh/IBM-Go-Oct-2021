package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	var ch chan int = make(chan int)
	wg.Add(1)
	go add(100, 200, wg, ch)
	//Reading data from the channel
	//wg.Wait()
	result := <-ch
	fmt.Println(result)
	wg.Wait()
	fmt.Println("exiting main")
}

func add(x, y int, wg *sync.WaitGroup, ch chan int) {
	result := x + y
	//write the result into the channel
	ch <- result
	wg.Done()
}
