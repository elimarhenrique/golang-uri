package main

import "fmt"

func main() {
	var m, n, aux, sum, index int
	var slice []int
	mapa := make(map[int][]int)

	for {

		fmt.Scanf("%d %d", &m, &n)

		if m <= 0 || n <= 0 {
			break
		}

		if m > n {
			aux = m
			m = n
			n = aux
		}

		for i := m; i <= n; i++ {
			slice = append(slice, i)
		}
		mapa[index] = slice
		slice = nil
		index++
	}

	for _, value := range mapa {
		for _, value2 := range value {
			sum += value2
			fmt.Printf("%d ", value2)
		}
		fmt.Printf("Sum=%d\n", sum)
		sum = 0
	}
}
