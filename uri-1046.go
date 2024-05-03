package main

import (
	"fmt"
)

func main() {
	var horaInicio, horaFim, duracaoJogo uint
	fmt.Scanf("%d", &horaInicio)
	fmt.Scanf("%d", &horaFim)

	if horaInicio == horaFim {
		fmt.Println("O JOGO DUROU 24 HORA(S)")
		return
	}

	if horaInicio > horaFim {
		duracaoJogo = (24 - horaInicio) + horaFim
		fmt.Printf("O JOGO DUROU %d HORA(S)\n", duracaoJogo)
		return
	} else {
		duracaoJogo = horaFim - horaInicio
		fmt.Printf("O JOGO DUROU %d HORA(S)\n", duracaoJogo)
		return
	}
}
