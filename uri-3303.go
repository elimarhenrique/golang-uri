package main

import "fmt"

func main() {
	var palavra string

	fmt.Scanln(&palavra)
	if len(palavra) >= 10 {
		fmt.Println("palavrao")
	} else {
		fmt.Println("palavrinha")
	}
}
