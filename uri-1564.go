package main

import (
	"fmt"
	"io"
)

func main() {
	var n int
	var slice []string

	for {
		_, err := fmt.Scanf("%d\n", &n)

		if err == io.EOF {
			break
		}

		if n == 0 {
			slice = append(slice, "vai ter copa!")
		} else {
			slice = append(slice, "vai ter duas!")
		}
	}

	for _, v := range slice {
		fmt.Println(v)
	}
}
