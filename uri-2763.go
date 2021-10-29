package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	cpf, _ := reader.ReadString('\n')

	for _, v := range cpf {
		if string(v) == "." ||
			string(v) == "-" {
			fmt.Printf("\n")
		} else {
			fmt.Printf("%s", string(v))
		}
	}
}
