package main

import (
	"fmt"
	"strconv"
)

func main() {
	leds := make(map[int]int)
	leds[0] = 6
	leds[1] = 2
	leds[2] = 5
	leds[3] = 5
	leds[4] = 4
	leds[5] = 5
	leds[6] = 6
	leds[7] = 3
	leds[8] = 7
	leds[9] = 6

	var n int
	fmt.Scanf("%d", &n)

	var slice []string
	for i := 0; i < n; i++ {
		var num string
		fmt.Scanf("%s", &num)
		slice = append(slice, num)
	}

	for _, value := range slice {
		sum := 0
		for _, value2 := range value {
			x, _ := strconv.Atoi(string(value2))
			sum += leds[x]
		}
		fmt.Println(sum, "leds")
	}
}
