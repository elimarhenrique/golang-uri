package main

import "fmt"

func main() {
	var n, x, y int
	var slice []int

	fmt.Scanf("%d", &n)

	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &x)
		fmt.Scanf("%d", &y)
		slice = append(slice, somaImparesConsecutivos(x, y))
	}

	for _, v := range slice {
		fmt.Println(v)
	}
}

func somaImparesConsecutivos(x, y int) int {
	var aux, soma int
	if x > y {
		aux = x
		x = y
		y = aux
	}

	for x = x + 1; x < y; x++ {
		if x%2 != 0 {
			soma += x
		}
	}
	return soma
}
