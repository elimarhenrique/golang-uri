package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	i := 1

	fmt.Scanln(&n)

	for i <= n {
		fmt.Println(i, math.Pow(float64(i), 2), math.Pow(float64(i), 3))
		i++
	}
}
