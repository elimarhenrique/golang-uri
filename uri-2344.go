package main

import "fmt"

func main() {
	var nota int

	fmt.Scanln(&nota)

	if nota >= 86 {
		fmt.Println("A")
	} else if nota >= 61 {
		fmt.Println("B")
	} else if nota >= 36 {
		fmt.Println("C")
	} else if nota >= 1 {
		fmt.Println("D")
	} else if nota < 1 {
		fmt.Println("E")
	}
}
