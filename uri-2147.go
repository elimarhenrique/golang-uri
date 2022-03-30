package main

import "fmt"

func main() {
	var slice []string
	var result []float32
	var galopeira string
	var n int
	var sum float32

	fmt.Scanln(&n)

	for i := 0; i < n; i++ {
		fmt.Scanln(&galopeira)
		slice = append(slice, galopeira)
		galopeira = ""
	}

	for _, value := range slice {
		for i := 0; i < len(value); i++ {
			sum += 0.01
		}
		result = append(result, sum)
		sum = 0
	}

	for _, value := range result {
		fmt.Printf("%.2f\n", value)
	}

}
