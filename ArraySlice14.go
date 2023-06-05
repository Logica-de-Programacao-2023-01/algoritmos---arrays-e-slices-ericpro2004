package main

import (
	"fmt"
)

func main() {
	//Solitação de dois índices para troca de posição em um slice.
	slice := make([]int, 8)
	for i := 0; i < len(slice); i++ {
		fmt.Printf("Informe o valor do elemento %d: ", i)
		fmt.Scan(&slice[i])
	}
	var indice1, indice2 int
	fmt.Print("Informe o índice do primeiro elemento: ")
	fmt.Scan(&indice1)
	fmt.Print("Informe o índice do segundo elemento: ")
	fmt.Scan(&indice2)
	if indice1 >= 0 && indice1 < len(slice) && indice2 >= 0 && indice2 < len(slice) {
		slice[indice1], slice[indice2] = slice[indice2], slice[indice1]
		fmt.Println("Slice resultante:", slice)
	} else {
		fmt.Println("Índices inválidos!")
	}
}
