// Convertir numeros romanos en decimales
// I -> 1
// V -> 5
// X -> 10
// L -> 50
// C -> 100
// D -> 500
// M -> 1000
// \u0305I -> Barra arriba multiplica por 1000 "vinculum"

/*
	Algorimtos:

1. Si la letra no tiene a la derecha otra de mayot valor, se suma el valor correspondiente
2. Si las letras I,X, C tienen a la derecha otra de mayot valor, se restan
3. Las letras I, X, C y M de pueden repetir hasta 3 veces.
4. Una raya encima los mltiplica por mill. Vinculum.
*/
package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

var romanValues = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func main() {
	fmt.Println("Numeros romanos")

}
