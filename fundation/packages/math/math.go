package math

// If you want to make a function public, the first letter of the function must be uppercase

// Sum is a function that receives two parameters and returns the sum of them
func Sum[T int | float64](a, b T) T {
	return a + b
}

// if the first letter of the function tiny "minuscule" this function is private
func sum[T int | float64](a, b T) T {
	return a + b
}

// in a struct is the same case the first letter of the struct must be uppercase to make it public
// example

// This is a public struct, is possible to import this struct in another package

// Car is a struct that represents a car
type Car struct {
	Brand string
}

// This is a private struct, is not possible to import this struct in another package

type car struct {
	brand string
}

func (c Car) Run() string {
	return "The car is running"
}
