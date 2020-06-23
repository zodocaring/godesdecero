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
	// vervo (%+v) muestra los campos de la estructura
	fmt.Printf("%+v\n", db)

	// asignamos valores literalmente
	html := curso{"Html", "Beto", "Bolivia"}
	fmt.Printf("%+v\n", html)
	// asignamos un valor a un solo campo, los tros campos se les asigna cero
	css := curso{Profesor: "Alvaro"}
	fmt.Printf("%+v\n", css)

	// para acceder a un valor del campo
	fmt.Printf("%+v\n", db.Nombre)
	fmt.Printf("%+v\n", html.Pais)
	fmt.Printf("%+v\n", css.Profesor)
}
