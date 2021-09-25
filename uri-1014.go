package main

import (
	"fmt"
)

func main() {
	var x int64
	var y float64
	fmt.Scanf("%d", &x)
	fmt.Scanf("%f", &y)

	saida := float64(x) / y
	fmt.Printf("%.3f km/l\n", saida)
}
