package main

import "fmt"

func main() {
	//Slice de elementos pares de um Array.
	array := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var slice []int
	for _, num := range array {
		if num%2 == 0 {
			slice = append(slice, num)
		}
	}
	fmt.Println(slice)
}
