package main

import (
	"fmt"
	"sync"
)

//communicate by sharing memory (NOT ADVISABLE)
type Result struct {
	Value int
	sync.Mutex
}

func (r *Result) SetValue(value int) {
	r.Lock()
	{
		r.Value = value
	}
	r.Unlock()
}

var result Result = Result{}

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go add(100, 200, wg)
	go add(1000, 2000, wg)
	wg.Wait()
	fmt.Println(result.Value)
	fmt.Println("exiting main")
}

func add(x, y int, wg *sync.WaitGroup) {
	result.SetValue(x + y)
	wg.Done()
}
