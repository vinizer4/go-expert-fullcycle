package main

import (
	"fmt"
	"sync"
	"time"
)

func task(name string, waitGroup *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d: Task %s is running\n", i, name)
		time.Sleep(1 * time.Second)
		waitGroup.Done() // Decrementa o contador do wait group
	}
}

// Thread 1
func main() {
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(25) // Adiciona duas goroutines ao wait group
	// Thread 2
	go task("A", &waitGroup)
	// Thread 3
	go task("B", &waitGroup)
	// Nada aqui.
	// Sai do programa

	// temos que manter a thread principal viva

	// Exemplo de goroutine anônima
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Printf("%d: Task %s is running\n", i, "Anonymous")
			time.Sleep(1 * time.Second)
			waitGroup.Done() // Decrementa o contador do wait group
		}
	}()

	waitGroup.Wait() // Espera as goroutines terminarem
}
