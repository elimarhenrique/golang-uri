package main

import "fmt"

func main() {
	var n, par, impar, positivo, negativo int

	for i := 0; i < 5; i++ {
		fmt.Scanln(&n)
		if n%2 == 0 {
			par++
		} else {
			impar++
		}

		if n > 0 {
			positivo++
		} else if n < 0 {
			negativo++
		}
	}

	fmt.Println(par, "valor(es) par(es)")
	fmt.Println(impar, "valor(es) impar(es)")
	fmt.Println(positivo, "valor(es) positivo(s)")
	fmt.Println(negativo, "valor(es) negativo(s)")

}
