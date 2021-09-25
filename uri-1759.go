package main

import "fmt"

func main() {
	var num int

	fmt.Scanln(&num)

	for i := 0; i < num-1; i++ {
		fmt.Printf("Ho ")
	}
	fmt.Println("Ho!")
}
