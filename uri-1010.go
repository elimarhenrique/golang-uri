package main

import "fmt"

func main() {
	var codigo1, peca1, codigo2, peca2 int
	var valor1, valor2 float64
	fmt.Scanf("%d", &codigo1)
	fmt.Scanf("%d", &peca1)
	fmt.Scanf("%f", &valor1)
	fmt.Scanf("%d", &codigo2)
	fmt.Scanf("%d", &peca2)
	fmt.Scanf("%f", &valor2)

	valor := float64(peca1)*valor1 + float64(peca2)*valor2

	fmt.Printf("VALOR A PAGAR: R$ %.2f\n", valor)
}
