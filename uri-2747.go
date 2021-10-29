package main

import "fmt"

func main() {

	for i := 0; i < 7; i++ {
		for j := 0; j < 39; j++ {
			if i == 0 || i == 7-1 {
				fmt.Printf("-")
			} else {
				if j == 0 || j == 39-1 {
					fmt.Printf("|")
				} else {
					fmt.Printf(" ")
				}
			}
		}
		fmt.Println()
	}

}
