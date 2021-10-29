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
		for j := 0; j < 2; j++ {
			fmt.Println(i, math.Pow(float64(i), 2)+float64(j), int(math.Pow(float64(i), 3)+float64(j)))
		}
		i++
	}
}
