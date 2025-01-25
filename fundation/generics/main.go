package main

type MyNumber int

func SumInt(m map[string]int) int {
	var sum int
	for _, value := range m {
		sum += value
	}
	return sum
}

func SumFloat(m map[string]float64) float64 {
	var sum float64
	for _, value := range m {
		sum += value
	}
	return sum
}

// Number is a constraint of types that can be used in the Sum function
type Number interface {
	// ~ is a type constraint that allows the use of int or float64 in this example has used ~ to available the use MyNumber int type too
	~int | ~float64
}

// Sum is a generic function that receives a map of int or float64 and returns the sum of the values
func Sum[T Number](m map[string]T) T {
	var sum T
	for _, value := range m {
		sum += value
	}
	return sum
}

// Compare is a generic function that receives two values and returns true if they are equal
func Compare[T comparable](a T, b T) bool {
	// the type comparable is a constraint that allows the use of types that can be compared
	if a == b {
		return true
	}
	return false
}

func main() {
	mapInt := map[string]int{"Wesley": 1000, "João": 2000, "Maria": 3000}
	mapFloat := map[string]float64{"Wesley": 1000.0, "João": 2000.0, "Maria": 3000.0}
	// Two methods with the same logic, but with different types
	println(SumInt(mapInt))
	println(SumFloat(mapFloat))

	// The same logic, but with a method that receives a generic type
	println(Sum(mapInt))
	println(Sum(mapFloat))

	mapIntMyNumber := map[string]MyNumber{"Wesley": 1000, "João": 2000, "Maria": 3000}
	println(Sum(mapIntMyNumber))

	// println(Compare(10, "ADC")) // Error: cannot use "ADC" (type string) as type int in argument to Compare
	println(Compare(10, 10))
	println(Compare("ADC", "ADC"))
}
