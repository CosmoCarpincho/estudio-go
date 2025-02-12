package main

import (
	"fmt"
)

func welcomeMessage(nombre string) (message string) {
	message = fmt.Sprintf("Bienvenido %s, al tutorial go. ", nombre)
	return message
}

func swap(x string, y string) (string, string) {
	return y, x
}

var varPaquete string = "estoy en todos el scope del paquete"
var soloDefinida string

var x, y, z int = 1, 2, 3

func main_2() {
	soloDentroDeFuncion := "estoy dentro solo de funcion"
	fmt.Println(welcomeMessage("Cosmo"))
	fmt.Println(swap("Carpincho", "Cosmo"))
	fmt.Println(soloDentroDeFuncion)
	fmt.Println(varPaquete)
	fmt.Println(x, y, z)

	var hola, que, tal = "Hola", "Que", "Tal"
	fmt.Println(hola, que, tal)

	// Tipos basicos de datos

	var entero int = 1
	var entero8 int8 = 1
	var entero16 int16 = 1
	var entero32 int32 = 1
	var entero64 int64 = 1
	var flotante32 float32 = 1.2
	var flotante64 float64 = 1.2
	var texto string = "string"
	var runa rune = 'a'
	var boleano bool = true
	var enteroPositivo uint = 1
	var enteroQueRepresentaDireccionesDeMemoria uintptr = 1
	var bytebyte byte = 2 // es igual a uint8
	var numeroComplejo complex64 = 2 + 2i
	var numeroComplejo2 complex128 = 2 + 2i

	fmt.Println(entero, entero8, entero16, entero32, entero64, flotante32,
		flotante64, texto, runa, boleano, enteroPositivo, enteroQueRepresentaDireccionesDeMemoria,
		bytebyte, numeroComplejo, numeroComplejo2)

}

func main() {
	// for

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	a := 10
	for a < 20 {
		fmt.Println(a)
		a++
	}

	texto := "que tal amigos"

	for k, t := range texto {
		fmt.Printf("k = %v  t= %v\n", k, string(t))

	}

	for {
		fmt.Println("a")
		break
	}

	// if

	edad := 20

	if fmt.Println("Esto se ejecuta antes de la condición"); edad >= 18 {
		fmt.Println("Puedes pasar al boliche")
	} else {
		fmt.Println("No podes pasar")
	}

	switch edad {
	case 20:
		fmt.Println("Podes comprar alcohol")
	default:
		fmt.Println("No podes comprar alcohol")
	}

	switch {
	case edad >= 18:
		fmt.Println("Podes comprar cerveza")
	default:
		fmt.Println("No podes comprar cerveza")
	}

	//defer
	diferir()

}

func diferir() {
	defer fmt.Println("defer 1")
	defer fmt.Println("defer 2")
	defer fmt.Println("defer 3")
	defer fmt.Println("defer 4")
}
