package main

import "fmt"

func main() {
	var slice []string
	var nota, soma float64
	var valida int

	for valida < 2 {
		fmt.Scanln(&nota)
		if nota >= 0 && nota <= 10 {
			soma += nota
			valida++
		} else {
			slice = append(slice, "nota invalida")
		}
	}

	for _, value := range slice {
		fmt.Println(value)
	}
	fmt.Printf("media = %.2f\n", soma/2.0)
}
