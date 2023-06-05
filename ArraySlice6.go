package main

import "fmt"

func main() {
	//Array de números inteiros com 10 elementos, com verificação de valores.
	array := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var valor int
	fmt.Print("Informe um valor: ")
	fmt.Scanln(&valor)
	encontrado := false
	for _, num := range array {
		if num == valor {
			encontrado = true
			break
		}
	}
	if encontrado {
		fmt.Printf("O valor %d está presente no array.\n", valor)
	} else {
		fmt.Printf("O valor %d não está presente no array.\n", valor)
	}
}
