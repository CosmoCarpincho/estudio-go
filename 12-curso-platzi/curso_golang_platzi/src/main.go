package main

import (
	"fmt"
	"math"
)

const pi = 3.141592

func main() {
	fmt.Println(areaCirculo(10))
	fmt.Println(areaTrapecio(10.8, 17.5, 8.35))
}

func areaCirculo(radio float64) float64 {
	areaCirculo := math.Pow(radio, 2) * pi
	return areaCirculo
}

func areaTrapecio(base1 float64, base2 float64, altura float64) float64 {
	areaTrapecio := (base1 + base2) * altura / 2

	return areaTrapecio

}
