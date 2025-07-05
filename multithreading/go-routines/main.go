package main

import (
	"fmt"
	"time"
)

func task(name string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d: Task %s is running\n", i, name)
		time.Sleep(1 * time.Second)
	}
}

// Thread 1
func main() {
	// Thread 2
	go task("A")
	// Thread 3
	go task("B")
	// Nada aqui.
	// Sai do programa

	// temos que manter a thread principal viva

	// Exemplo de goroutine anônima
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Printf("%d: Task %s is running\n", i, "Anonymous")
			time.Sleep(1 * time.Second)
		}
	}()

	time.Sleep(15 * time.Second) // Espera as goroutines terminarem
}
