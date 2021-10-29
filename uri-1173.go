package main

import "fmt"

func main() {
	var n int
	var array [10]int

	fmt.Scanln(&n)
	array[0] = n

	for i := 1; i < 10; i++ {
		array[i] = n * 2
		n = array[i]
	}

	for index, value := range array {
		fmt.Printf("N[%d] = %d\n", index, value)
	}
}
