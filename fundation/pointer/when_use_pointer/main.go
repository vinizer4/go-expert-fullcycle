package main

// in this example the function soma receiving a copy of the value of a and b
//func soma(a, b int) int {
//	a = 50
//	return a + b
//}

// in this example the function soma receiving a pointer to the memory address of a and b
func soma(a, b *int) int {
	*a = 50
	return *a + *b
}

func main() {
	myVar1 := 10
	myVar2 := 20

	soma(&myVar1, &myVar2)
	println(myVar1)
}
