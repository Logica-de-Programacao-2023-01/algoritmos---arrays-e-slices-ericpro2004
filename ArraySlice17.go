package main

import "fmt"

func main() {
	//Soma de elementos em posições pares de um Array.
	array := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var soma int
	for i := 0; i < len(array); i += 2 {
		soma += array[i]
	}
	fmt.Println("Soma dos elementos em posições pares:", soma)
}
