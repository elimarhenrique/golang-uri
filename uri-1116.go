package main

import "fmt"

func main() {
	var n, a, b int
	var slice []string

	fmt.Scanf("%d", &n)

	for i := 0; i < n; i++ {
		fmt.Scanf("%d %d", &a, &b)

		if b == 0 {
			slice = append(slice, "divisao impossivel")
		} else {
			result := float64(a) / float64(b)
			slice = append(slice, fmt.Sprintf("%.1f", result))
		}
	}

	for _, value := range slice {
		fmt.Println(value)
	}

}
