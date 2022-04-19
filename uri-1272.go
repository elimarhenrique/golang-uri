package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func mensagemOculta(slice []string) []string {
	var menssagem []string
	var aux string

	for _, value := range slice {
		array := strings.Split(value, " ")
		for _, letra := range array {
			if letra != "" {
				aux += strings.TrimSpace(string(letra[0]))
			}
		}
		menssagem = append(menssagem, aux)
		aux = ""
	}

	return menssagem
}

func main() {
	var slice []string

	var count int
	fmt.Scan(&count)

	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < count; i++ {
		text, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		slice = append(slice, text)
	}

	for _, value := range mensagemOculta(slice) {
		fmt.Println(value)
	}
}
