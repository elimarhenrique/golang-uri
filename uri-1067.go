package main

import "fmt"

func main() {
	var num int
	fmt.Scanln(&num)

	for i := 1; i <= num; i++ {
		if i%2 != 0 {
			fmt.Println(i)
		}
	}
}
