package main

import "fmt"

func main() {
	slice := make([]float64, 6)
	var num, positivo, soma float64

	for i := 0; i < len(slice); i++ {
		fmt.Scanln(&num)
		slice[i] = num

		if num >= 0 {
			positivo++
			soma += num
		}
	}

	fmt.Println(positivo, "valores positivos")
	fmt.Printf("%.1f\n", soma/positivo)
}
