package main

import "fmt"

type Result struct {
	opId  string
	value int
}

func main() {
	ch := make(chan Result)
	go add(1, 2, ch, "op1")
	go add(10, 20, ch, "op2")
	go add(100, 200, ch, "op3")
	go add(1000, 2000, ch, "op4")

	for i := 0; i < 4; i++ {
		result := <-ch
		fmt.Println(result.opId, ":", result.value)
	}
	fmt.Println("exiting main")
}

func add(x, y int, ch chan Result, opId string) {
	result := Result{
		opId:  opId,
		value: x + y,
	}
	ch <- result
}
