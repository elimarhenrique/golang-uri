package main

import (
	"fmt"
)

func main() {
	var n, x, min, pos int
	var slice []int

	fmt.Scanln(&n)

	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &x)
		slice = append(slice, x)
	}

	min = slice[0]
	for i := 1; i < len(slice); i++ {
		if slice[i] < min {
			min = slice[i]
			pos = i
		}
	}

	fmt.Println("Menor valor:", min)
	fmt.Println("Posicao:", pos)

}
