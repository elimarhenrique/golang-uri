package main

import "fmt"

func main() {
	var x int
	var slice []int

	fmt.Scanf("%d", &x)

	for x != 0 {
		slice = append(slice, somaParesConsecutivos(x))
		fmt.Scanf("%d", &x)
	}

	for _, v := range slice {
		fmt.Println(v)
	}
}

func somaParesConsecutivos(x int) int {
	var soma, i int
	for i < 5 {
		if x%2 == 0 {
			soma += x
			i++
		}
		x++
	}
	return soma
}
