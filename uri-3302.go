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

	for i, v := range slice {
		fmt.Printf("resposta %d: %d\n", i+1, v)
	}
}
