package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c int64
	fmt.Scanf("%d", &a)
	fmt.Scanf("%d", &b)
	fmt.Scanf("%d", &c)

	maiorAB := (float64(a) + float64(b) + math.Abs(float64(a-b))) / 2.0
	maiorABC := (maiorAB + float64(c) + math.Abs(maiorAB-float64(c))) / 2.0

	fmt.Println(maiorABC, "eh o maior")
}
