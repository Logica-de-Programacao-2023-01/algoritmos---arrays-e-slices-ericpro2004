package main

import "fmt"

func main() {
	//Solicitação de matriz com o valor de cada elemento.
	matriz := [3][2]int{}
	for i := 0; i < len(matriz); i++ {
		for j := 0; j < len(matriz[i]); j++ {
			fmt.Printf("Informe o valor do elemento [%d][%d]: ", i, j)
			fmt.Scanln(&matriz[i][j])
		}
	}
	fmt.Println("Matriz:")
	for i := 0; i < len(matriz); i++ {
		for j := 0; j < len(matriz[i]); j++ {
			fmt.Printf("%d ", matriz[i][j])
		}
		fmt.Println()
	}
}
