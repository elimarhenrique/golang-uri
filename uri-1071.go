package main

import "fmt"

func main() {
	var a, b, count, i, x int

	fmt.Scanln(&a)
	fmt.Scanln(&b)

	if a < b {
		i = a + 1
		x = b - 1
	} else {
		i = b + 1
		x = a - 1
	}

	for i <= x {
		if i%2 != 0 {
			count += i
		}
		i++
	}

	fmt.Println(count)
}
