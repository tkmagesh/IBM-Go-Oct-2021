package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
)

func main() {
	dataCh := make(chan int)
	fileWg := &sync.WaitGroup{}
	evenCh := make(chan int)
	oddCh := make(chan int)
	evenSumCh := make(chan int)
	oddSumCh := make(chan int)
	processWg := &sync.WaitGroup{}

	fileWg.Add(2)
	go Source("data1.dat", dataCh, fileWg)
	go Source("data2.dat", dataCh, fileWg)

	processWg.Add(4)
	go Splitter(dataCh, evenCh, oddCh, processWg)
	go Sum(evenCh, evenSumCh, processWg)
	go Sum(oddCh, oddSumCh, processWg)
	go Merger("result.dat", evenSumCh, oddSumCh, processWg)

	fileWg.Wait()
	close(dataCh)

	processWg.Wait()
	fmt.Println("Done")
}

func Source(fileName string, dataCh chan int, wg *sync.WaitGroup) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		file.Close()
		wg.Done()
	}()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data := scanner.Text()
		if no, err := strconv.Atoi(data); err == nil {
			dataCh <- no
		}
	}
}

func Splitter(dataCh, evenCh, oddCh chan int, wg *sync.WaitGroup) {
	defer func() {
		close(evenCh)
		close(oddCh)
		wg.Done()
	}()
	for no := range dataCh {
		if no%2 == 0 {
			evenCh <- no
		} else {
			oddCh <- no
		}
	}
}

func Sum(ch chan int, sumCh chan int, wg *sync.WaitGroup) {
	var sum int
	for no := range ch {
		sum += no
	}
	sumCh <- sum
	wg.Done()
}

func Merger(resultFile string, evenSumCh chan int, oddSumCh chan int, wg *sync.WaitGroup) {
	file, err := os.Create(resultFile)
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		file.Close()
		wg.Done()
	}()
	/*
		file.WriteString(fmt.Sprintf("Even total : %d\n", <-evenSumCh))
		file.WriteString(fmt.Sprintf("Odd total : %d\n", <-oddSumCh))
	*/
	for idx := 0; idx < 2; idx++ {
		select {
		case evenSum := <-evenSumCh:
			file.WriteString(fmt.Sprintf("Even total : %d\n", evenSum))
		case oddSum := <-oddSumCh:
			file.WriteString(fmt.Sprintf("Odd total : %d\n", oddSum))
		}
	}
}
