package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	var pares []int
	var impares []int

	fmt.Scanln(&n)

	for i := 0; i < n; i++ {
		var num int
		fmt.Scanln(&num)

		if num%2 == 0 {
			pares = append(pares, num)
		} else {
			impares = append(impares, num)
		}
	}

	sort.Ints(pares)
	sort.Sort(sort.Reverse(sort.IntSlice(impares)))
	for _, v := range pares {
		fmt.Println(v)
	}

	for _, v := range impares {
		fmt.Println(v)
	}
}
