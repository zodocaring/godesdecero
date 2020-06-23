package main

import "fmt"

// for range permite trabajar con slice(arreglos), maps y string

// slice
/* func main() {
	numeros := []uint8{2, 4, 6, 8}
	// devuelve  i, indice y v, valor
	for i, v := range numeros {
		fmt.Println("Indice: ", i, "Valor: ", v)
	}

} */

// ahora multiplivamos su valor v *=2

func main() {

	// ocultamos i con _ por que no la utilizamos
	numeros := []uint8{2, 4, 6, 8}
	for _, v := range numeros {
		v *= 2
		fmt.Println("Valor: ", v)
	}

}
