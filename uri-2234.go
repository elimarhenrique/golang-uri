package main

import "fmt"

func main() {

	var h, p int

	fmt.Scanf("%d %d", &h, &p)

	fmt.Printf("%.2f\n", float64(h)/float64(p))

}
