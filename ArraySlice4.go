package main

import (
	"fmt"
	"strconv"
)

func main() {
	//Solicitação de tamanho, valores e soma dos elementos de um slice.
	var tamanho int
	fmt.Print("Informe o tamanho do slice: ")
	fmt.Scanln(&tamanho)

	slice := make([]int, tamanho)
	for i := 0; i < tamanho; i++ {
		fmt.Printf("Informe o valor do elemento %d: ", i+1)
		var pergunta string
		fmt.Scanln(&pergunta)
		value, err := strconv.Atoi(pergunta)
		if err != nil {
			fmt.Println("Valor inválido. Tente novamente.")
			i--
			continue
		}
		slice[i] = value
	}
	fmt.Println("Slice:", slice)
	som := 0
	for _, value := range slice {
		som += value
	}
	fmt.Println("Soma dos valores:", som)
}
