package main

import "fmt"

func main() {
	//Soma de dois arrays solicitados.
	var n int
	fmt.Print("Digite um tamanho para os arrays: ")
	fmt.Scan(&n)

	array1 := make([]int, n)
	array2 := make([]int, n)

	fmt.Println("Digite os elementos do primeiro array: ")
	for i := 0; i < n; i++ {
		fmt.Printf("Elemento %d: ", i)
		fmt.Scan(&array1[i])
	}
	fmt.Println("Digite os elementos do segundo array: ")
	for i := 0; i < n; i++ {
		fmt.Printf("Elemento %d: ", i)
		fmt.Scan(&array2[i])
	}
	sumarray := make([]int, n)
	for i := 0; i < n; i++ {
		sumarray[i] = array1[i] + array2[i]
	}
	fmt.Println("O terceiro array, resultado da soma dos dois primeiros, é: ", sumarray)
}
