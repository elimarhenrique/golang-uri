package main

import "fmt"

func main() {
	j := 7
	i := 1
	for i <= 9 {
		for j >= 5 {
			fmt.Printf("I=%d J=%d\n", i, j)
			j--
		}
		j = 7
		i += 2
	}
}
