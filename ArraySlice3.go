package main

import "fmt"

func main() {
	//Array com 4 elementos em float.
	array := [4]float64{1.5, 2.5, 3.5, 4.5}
	produto := 1.0
	for _, valor := range array {
		produto *= valor
	}
	fmt.Println(produto)
}
