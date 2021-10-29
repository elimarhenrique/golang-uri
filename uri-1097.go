package main

import "fmt"

func main() {
	i := 1
	j := 6

	for i <= 9 {
		j += i
		for n := 0; n < 3; n++ {
			fmt.Printf("I=%d J=%d\n", i, j)
			j--
		}
		i += 2
		j = 6
	}
}
