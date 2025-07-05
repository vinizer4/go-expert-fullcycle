package main

import (
	"fmt"
	"sync"
)

// Thread 1
func main() {
	ch := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(10)

	go publishWG(ch)
	go readerWG(ch, &wg)

	wg.Wait()
}

func readerWG(ch chan int, wg *sync.WaitGroup) {
	for x := range ch {
		fmt.Printf("Received %d\n", x)
		wg.Done()
	}
}

func publishWG(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
}
