package main

import "fmt"

func main() {
	//Adição de número em todos os 6 elementos de um array em floats.
	array := [6]float64{}
	var numero float64
	fmt.Print("Informe um número: ")
	fmt.Scanln(&numero)
	for i := 0; i < len(array); i++ {
		array[i] = numero
	}
	fmt.Println("Array resultante:", array)
}
