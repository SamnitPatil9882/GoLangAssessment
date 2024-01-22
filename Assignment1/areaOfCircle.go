package main

import (
	"fmt"
	"math"
)

func main() {
	var radius float64
	fmt.Print("Enter radius: ")
	fmt.Scanln(&radius)
	var area float64
	area = math.Pi * math.Pow(radius, 2)
	fmt.Printf("area of circle : %.2f", area)

}
