package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func Reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

func criptografia(slice []string) []string {
	var retorno []string
	var aux1, aux2, aux3 string
	for _, value := range slice {
		for _, letra := range value {
			if unicode.IsLower(rune(letra)) || unicode.IsUpper(rune(letra)) {
				aux1 += string(letra + 3)
			} else {
				aux1 += string(letra)
			}
		}

		aux2 = Reverse(aux1)

		for index, letra := range aux2 {
			if index > (len(aux2)-1)/2 {
				aux3 += string(letra - 1)
			} else {
				aux3 += string(letra)
			}
		}
		retorno = append(retorno, aux3)
		aux1, aux2, aux3 = "", "", ""
	}

	return retorno
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

	for _, value := range criptografia(slice) {
		fmt.Println(strings.Replace(value, "\n", "", 1))
	}
}
