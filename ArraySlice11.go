package main

import (
	"fmt"
)

func main() {
	//Solitação de linha e coluna em um array.
	array := [2][3]int{}
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("Informe o valor do elemento na posição [%d][%d]: ", i, j)
			fmt.Scanln(&array[i][j])
		}
	}
	var linha, coluna int
	fmt.Print("Informe o índice da linha: ")
	fmt.Scanln(&linha)
	fmt.Print("Informe o índice da coluna: ")
	fmt.Scanln(&coluna)
	value := array[linha][coluna]
	fmt.Printf("O valor na posição [%d][%d] é: %d\n", linha, coluna, value)
}
