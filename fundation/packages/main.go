package main

import (
	"fmt"
	"go-course/math"
)

func main() {
	sum := math.Sum(10, 20)

	fmt.Println(sum)

	car := math.Car{Brand: "BMW"}
	fmt.Println(car.Run())
}
