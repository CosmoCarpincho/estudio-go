package main

import (
  "fmt"
  //"math"
)



func raiz_cuadrada(x float64) (float64, float64, float64) {

  // Menor distancia entre x e i^2
  var dist float64 = x;
  var numero float64;
  
  for i := 0. ; i < 2 * x; i++ {
    
    aux := x - i * i

    if aux < 0 {
      aux = -aux
    }

    if aux < dist {
      dist = aux
      numero = i
    }
}

// algoritmo método de newton

z := numero
for {
  aux := float32(z)
  z -= (z*z - x) / (2*x)

  if float32(z) == aux {
    break
  }

}

  
  return numero, dist, z

}

func main_raiz_cuadrada() {

  fmt.Println(raiz_cuadrada(10))
  fmt.Println(raiz_cuadrada(25))
  fmt.Println(raiz_cuadrada(40))
  fmt.Println(raiz_cuadrada(48))
  fmt.Println(raiz_cuadrada(100))
}
