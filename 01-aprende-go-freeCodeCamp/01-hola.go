package main

import (
	"fmt"

	"rsc.io/quote"
)

// Formas de declarar variables en go

var nombre string = "Cosmo"
var apellido = "Carpincho"

func main_1() {
	fmt.Println("Hello, World!")
	fmt.Println(quote.Go())

	lenguaje := "Go"
	fmt.Println(lenguaje)

	// Numeros
	var entero8 int8 = 1
	var entero32 int32 = 10000

	fmt.Println(entero8)
	fmt.Println(entero32)

	// Arreglos
	var listaFrutas = [4]string{"Pera", "Naranja"}
	fmt.Println(listaFrutas[1]) //Naranja

	// Slices
	var listaPaises = []string{"Peru", "Brasil", "Argentina"}
	fmt.Println(listaPaises)
	listaPaises = append(listaPaises, "Chile")
	fmt.Println(listaPaises)

	// Rangos
	listaPaises2 := listaPaises[1:3] // El final no es incluido
	fmt.Println(listaPaises2)

	listaPaises3 := listaPaises[2:]
	fmt.Println(listaPaises3)

	listaPaises4 := listaPaises[:2]
	fmt.Println(listaPaises4)
}
