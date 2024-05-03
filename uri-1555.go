package main

import (
	"fmt"
	"math"
)

func main() {
	var n, x, y int
	var slice []string

	fmt.Scanf("%d", &n)

	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &x)
		fmt.Scanf("%d", &y)

		r := rafael(x, y)
		b := beto(x, y)
		c := carlos(x, y)

		slice = append(slice, resultado(r, b, c))
	}

	for _, v := range slice {
		fmt.Println(v)
	}
}

func resultado(r, b, c float64) string {
	if r > b && r > c {
		return "Rafael ganhou"
	}

	if b > r && b > c {
		return "Beto ganhou"
	}

	return "Carlos ganhou"
}

func rafael(x, y int) float64 {
	res := math.Pow((3*float64(x)), 2) + math.Pow(float64(y), 2)
	return res
}

func beto(x, y int) float64 {
	res := (2 * math.Pow(float64(x), 2)) + math.Pow((5*float64(y)), 2)
	return res
}

func carlos(x, y int) float64 {
	res := -100*float64(x) + (math.Pow(float64(y), 3))
	return res
}
