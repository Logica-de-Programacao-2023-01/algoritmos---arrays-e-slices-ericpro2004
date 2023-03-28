package main

func main() {
	números := [6]float64{5, 6.5, 7.8, 9.77, 12, 3.4}
	var soma float64

	for _, número := range números {
		soma += número
	}

	média := soma / float64(len(números))
}
