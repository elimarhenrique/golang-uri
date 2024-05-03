package main

import "fmt"

func main() {
	var h, z, l int

	fmt.Scanf("%d", &h)
	fmt.Scanf("%d", &z)
	fmt.Scanf("%d", &l)

	if (h < l && h > z) || (h < z && h > l) {
		fmt.Println("huguinho")
		return
	}

	if (z < l && z > h) || (z < h && z > l) {
		fmt.Println("zezinho")
		return
	}

	fmt.Println("luisinho")
}
