package main

import "fmt"

var(
  vocales = "aeiou"
)
func main() {

  fmt.Println(contadorVocales("Lu"))
  fmt.Println(contadorVocales("aeiou aeiou"))
  fmt.Println(contadorVocales("aaoo"))
  fmt.Println(contadorVocales("que tal"))

}

func contadorVocales(palabra string) int {
  contador := 0

  for _, val := range palabra {
    for _, vocal := range vocales {
      if vocal == val {
        contador++
      }
    }
  }

  return contador
}
