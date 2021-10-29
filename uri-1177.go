package main

import "fmt"

func main() {
	var n, j int
	var array [1000]int

	fmt.Scanln(&n)

	for i := 0; i < 1000; i++ {
		if j == n {
			j = 0
		}
		array[i] = j
		j++
	}

	for index, value := range array {
		fmt.Printf("N[%d] = %d\n", index, value)
	}
}
