package main

import (
	"fmt"
	"os"
)

//
func main() {
	file, err := os.Create("archivo.txt")
	if err != nil {
		fmt.Printf("Ocurrio un error al crear el archivo: %v", err)
		return
	}
	// se crea el difer
	defer file.Close()

	_, err = file.Write([]byte("Este es un archivo con contenido plano"))
	if err != nil {

		fmt.Printf("Ocurrio un error al crear el archivo: %v", err)
		return
	}
}
