package main

import "fmt"

func main() {
	//Verificação em ordem crescente de um array.
	var n int
	fmt.Print("Digite o tamanho do array: ")
	fmt.Scan(&n)

	array := make([]int, n)

	fmt.Println("Digite os elementos do array:")
	for i := 0; i < n; i++ {
		fmt.Printf("Elemento %d: ", i)
		fmt.Scan(&array[i])
	}
	Ordenadacres := true
	for i := 0; i < n-1; i++ {
		if array[i] > array[i+1] {
			Ordenadacres = false
			break
		}
	}
	if Ordenadacres {
		fmt.Println("O array está ordenado em ordem crescente.")
	} else {
		fmt.Println("O array não está ordenado em ordem crescente.")
	}
}
