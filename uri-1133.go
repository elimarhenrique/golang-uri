package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y, aux int

	fmt.Scanf("%d", &x)
	fmt.Scanf("%d", &y)

	if x > y {
		aux = y
		y = x
		x = aux
	}

	for x < y {
		res := math.Mod(float64(x), 5)
		if res == 2 || res == 3 {
			fmt.Println(x)
		}
		x++
	}

}
