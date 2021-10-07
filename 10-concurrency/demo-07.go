package main

import (
	"fmt"
	"time"
)

func main() {

	var ch chan int = make(chan int)
	go add(100, 200, ch)

	//channel read is a blocking operation

	fmt.Println("[@main] Attempting to read data from the channel")
	result := <-ch
	fmt.Println("[@main] Completed reading data from the channel")

	fmt.Println(result)
	//time.Sleep(10 * time.Second)
	fmt.Println("exiting main")
}

func add(x, y int, ch chan int) {
	time.Sleep(4 * time.Second)
	result := x + y
	//write the result into the channel
	fmt.Println("[@add] Attempting to write data into the channel")
	/* The channel write will succeed ONLY when a read is already initiated */
	ch <- result
	fmt.Println("[@add] Completed writing data into the channel.")

}
