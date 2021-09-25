package main

import "fmt"

func main() {
	var number, horas int64
	var valor float64
	fmt.Scanln(&number)
	fmt.Scanln(&horas)
	fmt.Scanln(&valor)

	fmt.Printf("NUMBER = %d\n", number)
	fmt.Printf("SALARY = U$ %.2f\n", float64(horas)*valor)
}
