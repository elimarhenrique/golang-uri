package main

import "fmt"

func main() {
	var a, b, c, perimetro, area float64

	fmt.Scanf("%f", &a)
	fmt.Scanf("%f", &b)
	fmt.Scanf("%f", &c)

	if ((a + b) > c) && ((a + c) > b) && ((b + c) > a) {
		perimetro = a + b + c
		fmt.Printf("Perimetro = %.1f\n", perimetro)
	} else {
		area = ((a + b) * c) / 2.0
		fmt.Printf("Area = %.1f\n", area)
	}

}
