package main

import "fmt"

func main() {
	var num, a, b int
	var resp []int
	fmt.Scanln(&num)

	for i := 0; i < num; i++ {
		fmt.Scanf("%d", &a)
		fmt.Scanf("%d", &b)
		resp = append(resp, a+b)
	}

	for _, v := range resp {
		fmt.Println(v)
	}
}
