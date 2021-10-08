package main

import (
	"bufio"
	"io"
	"log"
	"os"
)

func main() {
	//reading sentence by sentence
	file, err := os.Open("data1.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	bufferedReader := bufio.NewReader(file)
	counter := 0
	for {
		line, err := bufferedReader.ReadString('.')
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		counter++
		log.Printf("%d: %s", counter, line)
	}
}
