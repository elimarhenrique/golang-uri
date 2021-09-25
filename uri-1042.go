package main

import (
	"fmt"
	"sort"
)

func main() {
	var a int64
	var slice, slice_orig []int64

	for i := 0; i < 3; i++ {
		fmt.Scanf("%d", &a)
		slice = append(slice, a)
		slice_orig = append(slice_orig, a)
	}

	sort.Slice(slice, func(i, j int) bool {
		return slice[i] < slice[j]
	})

	for _, v := range slice {
		fmt.Println(v)
	}

	fmt.Println()
	for _, v := range slice_orig {
		fmt.Println(v)
	}

}
