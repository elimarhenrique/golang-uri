package main

import "fmt"

func main() {
	var n int

	fmt.Scanln(&n)

	fmt.Printf("Feliz nat")
	for i := 0; i < n; i++ {
		fmt.Printf("a")
	}
	fmt.Printf("l!\n")
}
