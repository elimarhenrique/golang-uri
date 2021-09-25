package main

import "fmt"

func main() {
	var salario, novo_salario, reajuste float64
	var percent int64
	fmt.Scanln(&salario)

	if salario <= 400 {
		percent = 15
		reajuste = salario * 0.15
		novo_salario = salario + reajuste
	} else if salario > 400 && salario <= 800 {
		percent = 12
		reajuste = salario * 0.12
		novo_salario = salario + reajuste
	} else if salario > 800 && salario <= 1200 {
		percent = 10
		reajuste = salario * 0.10
		novo_salario = salario + reajuste
	} else if salario > 1200 && salario <= 2000 {
		percent = 7
		reajuste = salario * 0.07
		novo_salario = salario + reajuste
	} else {
		percent = 4
		reajuste = salario * 0.04
		novo_salario = salario + reajuste
	}

	fmt.Printf("Novo salario: %.2f\n", novo_salario)
	fmt.Printf("Reajuste ganho: %.2f\n", reajuste)
	fmt.Println("Em percentual:", percent, "%")
}
