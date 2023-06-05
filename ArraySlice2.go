package main

import "fmt"

func main() {
	//Slice de 5 elementos com a remoção de um (2).
	slice := []int{1, 2, 3, 4, 5}

	slice = append(slice[:2], slice[3:]...)

	fmt.Println(slice)
}
