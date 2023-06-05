package main

import (
	"fmt"
)

func main() {
	//Verificação de múltiplos de 3 em um array.
	array := [5]int{9, 4, 12, 7, 18}
	var multiples []int
	for _, num := range array {
		if num%3 == 0 {
			multiples = append(multiples, num)
		}
	}
	fmt.Println("Novo Slice com múltiplos de 3:", multiples)
}
