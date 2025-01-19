package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(sum(1, 2))        // 3
	fmt.Println(sumAndBool(1, 2)) // 3, true

	value, err := sumWithError(1, 2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(value)
	}

	value, err = sumWithError(35, 25)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(value)
	}
}

func sum(a int, b int) int {
	return a + b
}

// another way to declare the function signature
//func sum(a, b int) int {
//	return a + b
//}

// in go, you can return multiple values from a function
// in go is common to return an error as the last return value
func sumAndBool(a, b int) (int, bool) {
	return a + b, true
}

// this is an example of how to return an error
// is a common pattern in go to return an error as the last return value
func sumWithError(a, b int) (int, error) {
	if a+b >= 50 {
		return 0, errors.New("The sum is greater than 50")
	}
	return a + b, nil
}
