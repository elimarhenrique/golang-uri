package main

import "fmt"

func main() {
	var n, x, aux int
	var slice []int
	fmt.Scanln(&n)

	for i := 0; i < n; i++ {
		fmt.Scanln(&x)
		for i := 0; i < x; i++ {
			if i%2 == 0 {
				aux += 1
			} else {
				aux -= 1
			}
		}
		slice = append(slice, aux)
		x, aux = 0, 0
	}

	for _, value := range slice {
		fmt.Println(value)
	}
}
