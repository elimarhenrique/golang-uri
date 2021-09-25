package main

import "fmt"

func main() {
	slice := make([]int, 5)
	var num, par int

	for i := 0; i < 5; i++ {
		fmt.Scanln(&num)
		slice[i] = num

		if num%2 == 0 {
			par++
		}
	}

	fmt.Println(par, "valores pares")
}
