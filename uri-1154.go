package main

import "fmt"

func main() {
	var n, idades, idade int

	for {
		fmt.Scan(&idade)
		if idade < 0 {
			break
		}
		idades += idade
		n = n + 1
	}

	fmt.Printf("%.2f\n", float64(idades)/float64(n))
}
