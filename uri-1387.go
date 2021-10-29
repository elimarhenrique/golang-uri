package main

import (
	"fmt"
)

func main() {
	var l, r int
	var slice []int

	fmt.Scanf("%d %d", &l, &r)

	for l != 0 && r != 0 {
		slice = append(slice, (l + r))
		fmt.Scanf("%d %d", &l, &r)
	}

	for _, v := range slice {
		fmt.Println(v)
	}

}
