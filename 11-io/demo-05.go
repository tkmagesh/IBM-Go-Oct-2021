package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	data, err := ioutil.ReadFile("data1.txt")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(data))
}
