package main

import (
	"fmt"
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

	fmt.Println("x:", x, " y:", y)
	x++
	for x < y {
		if (x%5 == 2) || (x%5 == 3) {
			fmt.Println(x)
		}
		x++
	}

}
