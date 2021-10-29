package main

import "fmt"

func main() {
	var x float64
	var array [100]float64

	fmt.Scanf("%f", &x)

	array[0] = x
	for i := 1; i < 100; i++ {
		array[i] = x / 2
		x = array[i]
	}

	for index, value := range array {
		fmt.Printf("N[%d] = %.4f\n", index, value)
	}
}
