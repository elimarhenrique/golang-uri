package main

import "fmt"

func main() {
	var num int
	var slice []float64
	var a, b, c float64
	fmt.Scanln(&num)

	for i := 0; i < num; i++ {
		fmt.Scanf("%f", &a)
		fmt.Scanf("%f", &b)
		fmt.Scanf("%f", &c)

		media := (a*2 + b*3 + c*5) / 10.0
		slice = append(slice, media)
	}

	for _, v := range slice {
		fmt.Printf("%.1f\n", v)
	}
}
