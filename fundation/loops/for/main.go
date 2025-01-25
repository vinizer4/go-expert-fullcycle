package main

func main() {
	// for loop
	for i := 0; i < 10; i++ {
		println(i)
	}

	// slice of numbers names
	number := []string{"one", "two", "three"}

	// for loop with range
	for key, value := range number {
		println(key, value)
	}

	i := 0

	for i < 10 {
		println(i)
		i++
	}

	// infinite loop
	for {
		println("Infinite loop")
	}
}
