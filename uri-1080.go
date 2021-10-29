package main

import "fmt"

func main() {
	var num, p int
	var lista []int

	for i := 0; i < 100; i++ {
		fmt.Scanln(&num)
		lista = append(lista, num)
	}

	var x = lista[0]
	for pos, val := range lista {
		if val > x {
			x = val
			p = pos
		}
	}
	fmt.Println(x)
	fmt.Println(p + 1)
}
