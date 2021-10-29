package main

import "fmt"

func main() {
	var n, a, b int

	fmt.Scanln(&n)
	fmt.Scanf("%d %d", &a, &b)

	if (a + b) <= n {
		fmt.Println("Farei hoje!")
	} else {
		fmt.Println("Deixa para amanha!")
	}
}
