package main

import "fmt"

// for range con maps

func main() {

	deportes := map[string]string{"futbol": "balon pie", "basketball": "balon mano"}

	for i, v := range deportes {

		fmt.Println("Llave ", i, "Valor ", v)
	}

}
