package main

import (
	"errors"
	"fmt"
)

// caso uno, devulve error si el archivo no existe

func main() {
	/* content, err := ioutil.ReadFile("algo.txt")
	if err != nil {
		fmt.Printf("Ocurrio un error: %v", err)
		return // termina el if
	}
	// se ejecuta esta linea
	fmt.Println(string(content)) // se hace casting */

	// segundo caso, propia funcion que retorna un error

	result, err := division(10, 2)
	if err != nil {
		fmt.Printf("Ocurrio un error: %v", err)
		return // termina el if
	}
	fmt.Println(result)
}

func division(dividendo, divisor int) (int, error) {
	if divisor == 0 {
		// se puso 0, pero puede ir cualquier valor. este parametro no se ultiliza
		// crea una instancia de la clase errors

		return 0, errors.New("No se puede dividir por cero ")
	}
	// como segundo argumento devuelve nil, por que no hay error
	return dividendo / divisor, nil
}
