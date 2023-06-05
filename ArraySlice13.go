package main

import (
	"fmt"
)

func main() {
	//Adição de primeiro e último número em um array de 7 elementos.
	array := [7]int{2, 4, 6, 8, 10, 12, 14}
	var numero int
	fmt.Print("Informe um número: ")
	fmt.Scan(&numero)
	array[0] += numero
	array[len(array)-1] += numero
	fmt.Println("Array resultante:", array)
}
