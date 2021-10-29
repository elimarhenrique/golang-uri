package main

import "fmt"

func main() {
	var c [4]int

	fmt.Scanf("%d %d %d %d", &c[0], &c[1], &c[2], &c[3])

	for index, value := range c {
		if value == 1 {
			fmt.Println(index + 1)
		}
	}
}
