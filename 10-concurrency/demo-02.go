package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup = sync.WaitGroup{}

func main() {
	wg.Add(1)
	go f1()

	f2()
	wg.Wait()
	fmt.Println("exiting main")
}

func f1() {
	fmt.Println("f1 started")
	time.Sleep(5 * time.Second)
	fmt.Println("f1 completed")
	wg.Done()
}

func f2() {
	fmt.Println("f2 invoked")
}
