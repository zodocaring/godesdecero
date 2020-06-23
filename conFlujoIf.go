package main

import "fmt"

// estructuras, son clases en otros lenguajes

func main() {
	cadena := "Luis"
	if cadena == "Luis" {
		fmt.Println("Esta Ok  ", cadena)
	} else if cadena == "Cardozo" {
		fmt.Println("No Esta Ok  ")
	} else {
		fmt.Printf("La cadena es: %s", cadena)
	}
}
