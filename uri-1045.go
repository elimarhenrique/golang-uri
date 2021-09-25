package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	var a float64
	var slice, slice_orig []float64

	for i := 0; i < 3; i++ {
		fmt.Scanf("%f", &a)
		slice = append(slice, a)
		slice_orig = append(slice_orig, a)
	}

	sort.Slice(slice, func(i, j int) bool {
		return slice[i] > slice[j]
	})

	if slice[0] >= (slice[1] + slice[2]) {
		fmt.Println("NAO FORMA TRIANGULO")
	} else {
		if math.Pow(slice[0], 2) == (math.Pow(slice[1], 2) + math.Pow(slice[2], 2)) {
			fmt.Println("TRIANGULO RETANGULO")
		}

		if math.Pow(slice[0], 2) > (math.Pow(slice[1], 2) + math.Pow(slice[2], 2)) {
			fmt.Println("TRIANGULO OBTUSANGULO")
		}

		if math.Pow(slice[0], 2) < (math.Pow(slice[1], 2) + math.Pow(slice[2], 2)) {
			fmt.Println("TRIANGULO ACUTANGULO")
		}

		if slice[0] == slice[1] && slice[0] == slice[2] {
			fmt.Println("TRIANGULO EQUILATERO")
		} else if slice[0] == slice[1] || slice[1] == slice[2] || slice[0] == slice[2] {
			fmt.Println("TRIANGULO ISOSCELES")
		}

	}
}
