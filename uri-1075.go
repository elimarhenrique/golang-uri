package main

import "fmt"

func main() {
	var num int
	fmt.Scanln(&num)

	for i := 1; i <= 10000; i++ {
		if i%num == 2 {
			fmt.Println(i)
		}
	}
}
