package main

import "fmt"

func main() {
	passwd := 0
	var slice []string

	for passwd != 2002 {
		fmt.Scanln(&passwd)
		if passwd != 2002 {
			slice = append(slice, "Senha Invalida")
		} else {
			slice = append(slice, "Acesso Permitido")
		}
	}

	for _, v := range slice {
		fmt.Println(v)
	}
}
