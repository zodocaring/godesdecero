package main

import "fmt"

//
func main() {
	// estos parentesis llaman a la segunda funcion
	valor := hola("Pedro el escamoso")()
	fmt.Println(valor)
}

func hola(name string) func() string {
	return func() string {
		return "Hola " + name
	}

}
