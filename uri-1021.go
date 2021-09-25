package main

import (
	"fmt"
	"math"
)

func main() {
	var num float64
	fmt.Scanln(&num)

	fmt.Println("NOTAS:")
	notas_100 := num / 100.0
	fmt.Println(int(notas_100), "nota(s) de R$ 100.00")

	notas_50 := (int(num) % 100) / 50
	fmt.Println(notas_50, "nota(s) de R$ 50.00")

	notas_20 := ((int(num) % 100) % 50) / 20
	fmt.Println(notas_20, "nota(s) de R$ 20.00")

	notas_10 := (((int(num) % 100) % 50) % 20) / 10
	fmt.Println(notas_10, "nota(s) de R$ 10.00")

	notas_5 := ((((int(num) % 100) % 50) % 20) % 10) / 5
	fmt.Println(notas_5, "nota(s) de R$ 5.00")

	notas_2 := (((((int(num) % 100) % 50) % 20) % 10) % 5) / 2
	fmt.Println(notas_2, "nota(s) de R$ 2.00")

	fmt.Println("MOEDAS:")
	moeda_1 := (((((int(num) % 100) % 50) % 20) % 10) % 5) % 2
	fmt.Println(moeda_1, "moeda(s) de R$ 1.00")

	inteira := int(num)
	decimal := math.Round((num - float64(inteira)) * 100)

	moeda_50 := int(decimal) / 50
	fmt.Println(moeda_50, "moeda(s) de R$ 0.50")

	moeda_25 := (int(decimal) % 50) / 25
	fmt.Println(moeda_25, "moeda(s) de R$ 0.25")

	moeda_10 := ((int(decimal) % 50) % 25) / 10
	fmt.Println(moeda_10, "moeda(s) de R$ 0.10")

	moeda_5 := (((int(decimal) % 50) % 25) % 10) / 5
	fmt.Println(moeda_5, "moeda(s) de R$ 0.05")

	moeda_01 := (((int(decimal) % 50) % 25) % 10) % 5
	fmt.Println(moeda_01, "moeda(s) de R$ 0.01")

}
