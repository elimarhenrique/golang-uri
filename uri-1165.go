package main

import "fmt"

func main() {
	var n, num int
	var slice []int

	fmt.Scanln(&n)

	for i := 0; i < n; i++ {
		fmt.Scanln(&num)
		slice = append(slice, num)
	}

	for _, v := range slice {
		primo(v)
	}
}

func primo(n int) {
	var resultado int

	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			resultado++
			break
		}
	}

	if resultado == 0 {
		fmt.Printf("%d eh primo\n", n)
	} else {
		fmt.Printf("%d nao eh primo\n", n)
	}
}
