package main

import "fmt"

func main() {
	var n, alcool, gasolina, diesel int

	fmt.Scanln(&n)
	for n != 4 {

		if n == 1 {
			alcool++
		} else if n == 2 {
			gasolina++
		} else if n == 3 {
			diesel++
		}
		fmt.Scanln(&n)
	}

	fmt.Println("MUITO OBRIGADO")
	fmt.Println("Alcool:", alcool)
	fmt.Println("Gasolina:", gasolina)
	fmt.Println("Diesel:", diesel)
}
