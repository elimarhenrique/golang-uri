package main

import (
	"fmt"
	"math"
)

func main() {
	const n = 3.14159
	var a float64
	fmt.Scanf("%f", &a)
	fmt.Printf("A=%.4f\n", n*(math.Pow(a, 2)))
}
