package main

import "fmt"

func main() {
	var num, n int
	var slice []int
	fmt.Scanln(&num)

	for i := 0; i < num; i++ {
		fmt.Scanln(&n)
		slice = append(slice, n)
	}

	for _, v := range slice {
		if v == 0 {
			fmt.Println("NULL")
		} else {
			if v%2 == 0 {
				if v > 0 {
					fmt.Println("EVEN POSITIVE")
				} else {
					fmt.Println("EVEN NEGATIVE")
				}
			} else {
				if v > 0 {
					fmt.Println("ODD POSITIVE")
				} else {
					fmt.Println("ODD NEGATIVE")
				}
			}
		}

	}
}
