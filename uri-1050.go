package main

import "fmt"

func main() {
	var ddd int
	fmt.Scanln(&ddd)

	if ddd == 61 {
		fmt.Println("Brasilia")
	} else if ddd == 71 {
		fmt.Println("Salvador")
	} else if ddd == 11 {
		fmt.Println("Sao Paulo")
	} else if ddd == 21 {
		fmt.Println("Rio de Janeiro")
	} else if ddd == 32 {
		fmt.Println("Juiz de Fora")
	} else if ddd == 19 {
		fmt.Println("Campinas")
	} else if ddd == 27 {
		fmt.Println("Vitoria")
	} else if ddd == 31 {
		fmt.Println("Belo Horizonte")
	} else {
		fmt.Println("DDD nao cadastrado")
	}

}
