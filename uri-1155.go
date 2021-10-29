package main

import (
	"fmt"
	"math"
)

func main() {
	s := 1.0

	for i := 1; i <= 100; i++ {
		s += 1.0 / (float64(i) + 1)
	}

	// round floating-point value
	fmt.Println(math.Floor(s*100) / 100)
}
