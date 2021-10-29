package main

import "fmt"

func main() {
	var n int

	fmt.Scanln(&n)
	if n%2 == 0 {
		fmt.Printf("%d casas brancas e %d casas pretas\n", n*n/2, n*n/2)
	} else {
		fmt.Printf("%d casas brancas e %d casas pretas\n", (n*n/2)+1, n*n/2)
	}
}
