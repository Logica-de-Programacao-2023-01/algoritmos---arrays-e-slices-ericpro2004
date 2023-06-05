package main

import (
	"fmt"
	"math"
)

func main() {
	// Valor mínimo e máximo de um slice de 10 elementos.
	slice := make([]int, 10)
	for i := 0; i < len(slice); i++ {
		fmt.Printf("Informe o valor do elemento %d: ", i+1)
		fmt.Scanln(&slice[i])
	}
	min := math.MaxInt64
	max := math.MinInt64
	for _, value := range slice {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	fmt.Println("Valor mínimo:", min)
	fmt.Println("Valor máximo:", max)
}
