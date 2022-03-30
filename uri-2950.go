package main

import "fmt"

func main() {
	var distancia, diametro1, diametro2 int

	fmt.Scanf("%d %d %d", &distancia, &diametro1, &diametro2)

	result := float64(distancia) / (float64(diametro1) + float64(diametro2))
	fmt.Printf("%.2f\n", result)

}
