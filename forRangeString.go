package main

import "fmt"

// for range con maps

func main() {

	cadena := "Hola Mundo"

	for _, v := range cadena {

		fmt.Println(string(v))
	}

}
