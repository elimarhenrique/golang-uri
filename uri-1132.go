package main

import (
	"fmt"
)

func main() {
	var x, y, aux, sum int

	fmt.Scanf("%d", &x)
	fmt.Scanf("%d", &y)

	if x > y {
		aux = y
		y = x
		x = aux
	}

	for x <= y {
		if x%13 != 0 {
			sum += x
		}
		x++
	}

	fmt.Println(sum)
}
