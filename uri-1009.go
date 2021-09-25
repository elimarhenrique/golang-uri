package main

import "fmt"

func main() {
	var salario, vendas float64
	var nome string
	fmt.Scanln(&nome)
	fmt.Scanln(&salario)
	fmt.Scanln(&vendas)

	fmt.Printf("TOTAL = R$ %.2f\n", salario+(vendas*0.15))
}
