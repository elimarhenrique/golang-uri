package main

import (
	"fmt"
	"sort"
)

func main() {
	var dm, fa, fb int
	var slice []int

	fmt.Scanln(&dm)
	slice = append(slice, dm)
	fmt.Scanln(&fa)
	slice = append(slice, fa)
	fmt.Scanln(&fb)
	slice = append(slice, fb)

	fc := dm - (fa + fb)
	slice = append(slice, fc)

	sort.Ints(slice)

	fmt.Println(slice[2])

}
