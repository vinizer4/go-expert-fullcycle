package main

import (
	"fmt"
	"github.com/devfullcycle/goexpert/7-Packaging/3/math"
	"github.com/google/uuid"
)

func main() {
	m := math.NewMath(1, 2)
	fmt.Println(m.Add())
	fmt.Println(uuid.New())
}
