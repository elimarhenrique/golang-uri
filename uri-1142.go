package main

import "fmt"

func main() {
	var num, j int

	fmt.Scanln(&num)

	for i := 1; i <= num; i++ {
		for k := 1; k < 4; k++ {
			j++
			fmt.Printf("%d ", j)
		}
		fmt.Printf("PUM\n")
		j++
	}

}
