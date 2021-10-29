package main

import (
	"fmt"
)

func main() {
	var n, x int
	var msg []int

	fmt.Scanln(&n)

	for i := 0; i < n; i++ {
		fmt.Scanf("%x", &x)
		msg = append(msg, x)
	}

	for _, v := range msg {
		fmt.Printf("%c", v)
	}
	fmt.Println()
}
