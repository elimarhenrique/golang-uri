package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var slice []string
	for {
		n, m := lerValores()
		if n == 0 && m == 0 {
			break
		}
		slice = append(slice, soma(n, m))
	}

	for _, v := range slice {
		fmt.Println(v)
	}
}

func lerValores() (int, int) {
	var n, m int

	fmt.Scanf("%d", &n)
	fmt.Scanf("%d", &m)

	return n, m
}

func soma(n, m int) string {
	soma := n + m
	return strings.ReplaceAll(strconv.Itoa(soma), "0", "")
}
