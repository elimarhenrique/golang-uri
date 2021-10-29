package main

import "fmt"

func main() {
	var n, h, d, g int
	var resp []string

	fmt.Scanln(&n)

	for i := 0; i < n; i++ {
		fmt.Scanf("%d %d %d", &h, &d, &g)

		if (h >= 200 && h <= 300) && (d >= 50) && (g >= 150) {
			resp = append(resp, "Sim")
		} else {
			resp = append(resp, "Nao")
		}
	}

	for _, v := range resp {
		fmt.Println(v)
	}
}
