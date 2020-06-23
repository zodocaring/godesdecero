package main

import "fmt"

// estructuras, son clases en otros lenguajes

/* func main() {
	animal := "Gato"
	switch animal {
	case "Gato":
		fmt.Println("Es un Gato")
	case "Perro":
		fmt.Println("Es un perro")
	default:
		fmt.Println("No es nada!!")

	}
} */

// segundo caso

/* func main() {
	cosa := "mesa"
	switch cosa {
	case "Gato", "Perro":
		fmt.Println("Es un animal")
	case "Banana", "Manzana":
		fmt.Println("Es una fruta")
	default:
		fmt.Println("No es animal ni fruta")

	}
} */

// tercer caso, operadores logicos

func main() {
	cosa := "Manzana"
	switch cosa {
	case cosa == "Gato" || cosa == "Perro":
		fmt.Println("Es un animal")
	case cosa == "Banana" || cosa == "Manzana":
		fmt.Println("Es una fruta")
	default:
		fmt.Println("No es animal ni fruta")

	}
}
