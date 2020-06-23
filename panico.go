package main

import (
	"fmt"
)

//
func main() {
	division(10, 2)
	division(10, 0)
}

func division(dividendo, divisor int) {
	validarDivisor(divisor)
	fmt.Println(dividendo / divisor)
}

func validarDivisor(divisor int) {
	if divisor == 0 {
		panic("Panico ")
	}
}
