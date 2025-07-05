package main

func main() {
	// 2 is the buffer size
	ch := make(chan string, 2)
	ch <- "Hello"
	ch <- "World"

	println(<-ch)
	println(<-ch)
}
