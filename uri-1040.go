package main

import (
	"fmt"
)

func main() {
	var n1, n2, n3, n4, exame, nota_final float64
	fmt.Scanf("%f", &n1)
	fmt.Scanf("%f", &n2)
	fmt.Scanf("%f", &n3)
	fmt.Scanf("%f", &n4)

	media := (n1*2 + n2*3 + n3*4 + n4*1) / 10.0

	if media >= 5 && media < 7 {
		fmt.Scanf("%f", &exame)
		nota_final = (media + exame) / 2.0
	}

	fmt.Printf("Media: %.1f\n", media)
	if media >= 7 {
		fmt.Println("Aluno aprovado.")
	} else if media < 5 {
		fmt.Println("Aluno reprovado.")
	} else {
		fmt.Println("Aluno em exame.")
		fmt.Printf("Nota do exame: %.1f\n", exame)
		if nota_final >= 5 {
			fmt.Println("Aluno aprovado.")
		} else {
			fmt.Println("Aluno reprovado.")
		}
		fmt.Printf("Media final: %.1f\n", nota_final)
	}
}
