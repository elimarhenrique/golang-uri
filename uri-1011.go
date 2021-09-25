package main

import (
	"fmt"
	"math"
)

func main() {
	pi := 3.14159
	var raio int
	fmt.Scanf("%d", &raio)

	volume := (4 / 3.0) * pi * math.Pow(float64(raio), 3)
	fmt.Printf("VOLUME = %.3f\n", volume)
}
