package main

import (
	"fmt"
)

// aqui se utiliza difer y panico para utilizar recover
func main() {
	division(10, 2)
	division(10, 0)
	division(100, 2)
	division(60, 20)
}

func division(dividendo, divisor int) {
	// se aplica defer
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recuperandome del Panico ", r)
		}
	}()
	validarDivisor(divisor)
	fmt.Println(dividendo / divisor)
}

func validarDivisor(divisor int) {
	if divisor == 0 {
		panic("Se produjo un error, divisor con valor cero!! ")
	}
}
