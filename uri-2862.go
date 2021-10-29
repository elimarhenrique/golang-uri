package main

import "fmt"

func main() {
	var x, c int
	var n []int

	fmt.Scanln(&x)
	for i := 0; i < x; i++ {
		fmt.Scanln(&c)
		n = append(n, c)
	}

	for _, value := range n {
		if value > 8000 {
			fmt.Println("Mais de 8000!")
		} else {
			fmt.Println("Inseto!")
		}
	}

}
