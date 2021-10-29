package main

import "fmt"

func main() {
	var x, y int
	var slice []string

	fmt.Scanf("%d", &x)
	fmt.Scanf("%d", &y)

	for x != y {
		if x > y {
			slice = append(slice, "Decrescente")
		} else if x < y {
			slice = append(slice, "Crescente")
		} else if x == y {
			break
		}

		fmt.Scanf("%d", &x)
		fmt.Scanf("%d", &y)
	}

	for _, v := range slice {
		fmt.Println(v)
	}
}
