package main

import "fmt"

func main() {
	var refeicoes [3]int
	var pedidos [3]int

    for i:=0; i<3; i++ {
        fmt.Scanf("%d", &refeicoes[i])
    }

    for i:=0; i<3; i++ {
        fmt.Scanf("%d", &pedidos[i])
    }


    var soma int
    for i, pedidos := range pedidos {
        if pedidos >= refeicoes[i] {
            soma += (pedidos - refeicoes[i])
        }
    }

    fmt.Println(soma)

}
