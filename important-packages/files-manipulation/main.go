package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("file.txt")
	if err != nil {
		panic(err)
	}

	size, err := f.Write([]byte("Writing data to the file"))

	// This is another way to write data to the file
	// size, err := f.WriteString("Hello, World!")
	if err != nil {
		panic(err)
	}
	fmt.Printf("File created with size: %d bytes\n", size)
	f.Close()

	// Open the file for reading
	file, err := os.ReadFile("file.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("File content: %s\n", string(file))

	// Reading the file in parts using buffer with package bufio
	file2, err := os.Open("file.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(file2)
	buffer := make([]byte, 5)
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Print(string(buffer[:n]))
	}
}
