package main

import "fmt"

func main() {
	var n, x, in, out int
	var slice []int

	fmt.Scanln(&n)
	for i := 0; i < n; i++ {
		fmt.Scanln(&x)
		slice = append(slice, x)
	}

	for _, v := range slice {
		if v >= 10 && v <= 20 {
			in++
		} else {
			out++
		}
	}

	fmt.Println(in, "in")
	fmt.Println(out, "out")
}
