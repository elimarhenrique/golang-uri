package main

import "fmt"

func main() {
	var num, count int

	fmt.Scanln(&num)

	for {
		if num%2 != 0 {
			fmt.Printf("%d\n", num)
			count++
		}
		if count >= 6 {
			break
		}
		num++
	}
}
