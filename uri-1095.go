package main

import "fmt"

func main() {
	j := 60
	i := 1

	for j >= 0 {
		fmt.Printf("I=%d J=%d\n", i, j)
		i = i + 3
		j = j - 5
	}
}
