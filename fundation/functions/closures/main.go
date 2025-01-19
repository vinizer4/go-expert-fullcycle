package main

func main() {

	// this is an example of an anonymous function that is assigned to a variable
	// this function is a closure because it uses a variable from the outer scope
	total := func() int {
		return sum(1, 2, 3, 4, 5) * 2
	}()

	println(total)
}

func sum(numbes ...int) int {
	total := 0
	for _, number := range numbes {
		total += number
	}
	return total
}
