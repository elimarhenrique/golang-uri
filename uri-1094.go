package main

import "fmt"

func main() {
	cobaias := make(map[string]int)
	var n, total int

	fmt.Scanf("%d", &n)

	for i := 0; i < n; i++ {
		var qtd int
		var c string
		fmt.Scanf("%d %s", &qtd, &c)

		cobaias[c] += qtd
	}

	for _, value := range cobaias {
		total += value
	}

	fmt.Println("Total:", total, "cobaias")
	fmt.Println("Total de coelhos:", cobaias["C"])
	fmt.Println("Total de ratos:", cobaias["R"])
	fmt.Println("Total de sapos:", cobaias["S"])

	var percentualC = float64(cobaias["C"]) * 100.0 / float64(total)
	fmt.Printf("Percentual de coelhos: %.2f %%\n", percentualC)
	var percentualR = float64(cobaias["R"]) * 100.0 / float64(total)
	fmt.Printf("Percentual de ratos: %.2f %%\n", percentualR)
	var percentualS = float64(cobaias["S"]) * 100.0 / float64(total)
	fmt.Printf("Percentual de sapos: %.2f %%\n", percentualS)

}
