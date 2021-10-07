package main

import (
	"fmt"
	"sync"
)

//communicate by sharing memory (NOT ADVISABLE)
var result int
var mutex = sync.Mutex{}

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go add(100, 200, wg)
	go add(1000, 2000, wg)
	wg.Wait()
	fmt.Println(result)
	fmt.Println("exiting main")
}

func add(x, y int, wg *sync.WaitGroup) {
	mutex.Lock()
	{
		result = x + y
	}
	mutex.Unlock()
	wg.Done()
}
