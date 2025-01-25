package main

import (
	"fmt"
	"github.com/google/uuid"
	"go-course/math"
)

func main() {
	sum := math.Sum(10, 20)

	fmt.Println(sum)

	car := math.Car{Brand: "BMW"}
	fmt.Println(car.Run())

	fmt.Println(uuid.New())
}
