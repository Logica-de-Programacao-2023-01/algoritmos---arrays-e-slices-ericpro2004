package main

import "fmt"

func main() {
	// Adição de números a um slice de inteiros com tamanho 5 elementos.
	slice := make([]int, 0, 5)
	var numero int
	fmt.Print("Informe um número inteiro: ")
	fmt.Scanln(&numero)
	estaPresente := false
	for _, valor := range slice {
		if valor == numero {
			estaPresente = true
			break
		}
	}
	if !estaPresente {
		slice = append(slice, numero)
	}

	fmt.Println("Slice resultante:", slice)
}
