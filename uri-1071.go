package main

import "fmt"

func main() {
	var a, b, count, i, x int

	fmt.Scanln(&a)
	fmt.Scanln(&b)
	fmt.Println(a, b)

	if a < b {
		i = a
		x = b
	} else {
		i = b
		x = a
	}

	for i < x-1 {
		i++
		if i > 0 && i%2 != 0 {
			count += i
		}
	}

	fmt.Println(count)
}
