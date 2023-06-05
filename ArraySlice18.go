package main

import "fmt"

func main() {
	//Array de números primos de um número solicitado.
	var n int
	fmt.Print("Digite um número inteiro (positivo): ")
	fmt.Scan(&n)
	Nump := make([]int, 0)
	conta := 0
	num := 2
	for conta < n {
		primo := true
		for i := 2; i <= num/2; i++ {
			if num%i == 0 {
				primo = false
				break
			}
		}
		if primo {
			Nump = append(Nump, num)
			conta++
		}
		num++
	}
	fmt.Println("Os primeiros números primos são: ", n)
	fmt.Println(Nump)
}
