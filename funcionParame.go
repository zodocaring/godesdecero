package main

import "fmt"

//funciones, paso por referencia de valores

func main() {
	hola("luis", "cardozo")
}

func hola(nombre string, apellido string) {
	fmt.Printf("Como estas %s %s ", nombre, apellido)
}
