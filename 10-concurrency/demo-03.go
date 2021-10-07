package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(2)
	//wg.Add(1)
	go print("Hello", wg)

	//wg.Add(1)
	go print("World", wg)
	wg.Wait()
}

func print(str string, wg *sync.WaitGroup) {
	fmt.Println(str)
	wg.Done()
}
