package main

import (
	"fmt"
	"strings"
)

// retormo de dos multiples valores

func main() {
	cadena := "JuAN paBLo CarDEnas"
	minus, mayusc := convert(cadena)
	fmt.Println(minus, mayusc)
}

func convert(texto string) (string, string) {
	minuscula := strings.ToLower(texto)
	mayuscula := strings.ToUpper(texto)
	return minuscula, mayuscula
}
