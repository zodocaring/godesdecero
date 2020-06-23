package main

import "fmt"

// estructuras, son clases en otros lenguajes

func main() {
	type curso struct {
		Nombre   string
		Profesor string
		Pais     string
	}
	// creamos una istancia
	db := curso{
		Nombre:   "Base de datos",
		Profesor: "Lacc",
		Pais:     "Colombia",
	}

	// variable tipo puntero, utiliza el operador &
	p := &db
	fmt.Printf("%+v\n", db)
	fmt.Printf("%+v\n", p)
	// cambiamos el valor de un campo con el puntero
	(*p).Profesor = "Cardozo"
	fmt.Printf("%+v\n", db)
	fmt.Printf("%+v\n", p)
	// no es necesario referenciar con *
	p.Profesor = "Uribe"
	fmt.Printf("%+v\n", db)
	fmt.Printf("%+v\n", p)
}
